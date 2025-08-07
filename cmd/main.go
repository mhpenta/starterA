package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jessevdk/go-flags"
	"github.com/mhpenta/app"
	"github.com/rs/cors"
	"golang.org/x/crypto/acme/autocert"
	"log/slog"
	"net/http"
	"os"
	"starterA/internal/application"
	"starterA/internal/config"
	"starterA/internal/database"
	"starterA/internal/routes"
	"starterA/internal/service"
	"time"
)

// Options defines command-line options for the application
type Options struct {
	ConfigPath string `short:"c" long:"config" description:"Path to configuration file" default:"config.toml"`
	Verbose    bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
}

func main() {
	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	cfg, err := config.Load(opts.ConfigPath)
	if err != nil {
		slog.Error("Error loading config", "error", err)
		os.Exit(1)
	}

	ctx, cancel := app.MainContext()
	defer cancel()

	if err := run(ctx, cfg); err != nil {
		slog.Error("Error running", "error", err)
	}
}

func run(ctx context.Context, cfg *config.Config) error {
	db, err := database.GetDatabase(cfg.Database)
	if err != nil {
		return fmt.Errorf("error getting DB connection: %w", err)
	}

	// Initialize logger
	logger := slog.Default()

	// Initialize service layer
	svc := service.New(db, logger, cfg)

	// Initialize application layer
	a := application.New(svc, logger, cfg)

	return runServer(ctx, cfg, a)
}

// runServer starts the server using the given configuration and initializes routes
func runServer(
	ctx context.Context,
	cfg *config.Config,
	a *application.Application) error {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Duration(cfg.Server.TaskTimeOutInSeconds) * time.Second)) // Use the longest timeout for all routes by default

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: cfg.Server.AllowedCorsURLs,
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})
	wrappedHandler := corsHandler.Handler(r)

	routes.RegisterRoutes(r, a)

	server := &http.Server{
		Addr:              ":" + fmt.Sprint(cfg.Server.Port),
		Handler:           wrappedHandler,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go serve(cfg, server, a.Logger)
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Forced shutdown", "err", err)
	}

	return nil
}

func serve(cfg *config.Config, server *http.Server, logger *slog.Logger) {
	if cfg.Server.EnableHTTPS {

		// Note, this is to be used when running on a server, as opposed to a serverless platform with automatic HTTPS support
		logger.Info("Starting HTTPS server")

		certManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Cache:      autocert.DirCache("certs"),
			HostPolicy: autocert.HostWhitelist(cfg.Server.ServerDomain),
		}

		server.Addr = ":" + fmt.Sprint(cfg.Server.HTTPSPort)
		server.TLSConfig = &tls.Config{
			GetCertificate: certManager.GetCertificate,
		}

		// Serve HTTP redirect to HTTPS
		go func() {
			err := http.ListenAndServe(":"+fmt.Sprint(cfg.Server.Port), certManager.HTTPHandler(nil))
			if err != nil {
				logger.Error("Could not start HTTP server", "err", err)
			}
		}()

		// Start HTTPS server
		err := server.ListenAndServeTLS("", "")
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Could not start HTTPS server", "err", err)
		}
		logger.Info("Listening on port " + fmt.Sprint(cfg.Server.HTTPSPort) + " with HTTPS enabled")
	} else {
		logger.Info("Starting HTTP server")
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Could not start HTTP server", "err", err)
		}
	}
}
