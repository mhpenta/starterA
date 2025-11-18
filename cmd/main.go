package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mhpenta/starterA/internal/app"
	"github.com/mhpenta/starterA/internal/config"
	httphandlers "github.com/mhpenta/starterA/internal/handlers/http"
	"github.com/mhpenta/starterA/internal/routes"
	"github.com/mhpenta/starterA/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jessevdk/go-flags"
	"github.com/rs/cors"
	"golang.org/x/crypto/acme/autocert"
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

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGTERM)
	defer cancel()

	if err := run(ctx, cfg); err != nil {
		slog.Error("Error running application", "error", err)
	}
}

func run(ctx context.Context, cfg *config.Config) error {
	logger := slog.Default()

	a, err := app.New(ctx, logger, cfg)
	if err != nil {
		return fmt.Errorf("app initialization error: %w", err)
	}
	defer func(a *app.Application) {
		err := a.Close()
		if err != nil {
			slog.Error("Error closing application", "error", err)
		}
	}(a)

	svc := service.New(ctx, a, a.Logger)

	httpHandlers := httphandlers.New(svc, a.Logger)

	return runServer(ctx, cfg.Server, a, httpHandlers)
}

// runServer starts the server using the given configuration and initializes routes
func runServer(
	ctx context.Context,
	serverCfg config.Server,
	a *app.Application,
	httpHandlers *httphandlers.HTTPHandlers) error {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Duration(serverCfg.TaskTimeOutInSeconds) * time.Second)) // Use the longest timeout for all routes by default

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: serverCfg.AllowedCorsURLs,
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})
	wrappedHandler := corsHandler.Handler(r)

	routes.RegisterRoutes(r, httpHandlers)

	server := &http.Server{
		Addr:              ":" + fmt.Sprint(serverCfg.Port),
		Handler:           wrappedHandler,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go serve(serverCfg, server, a.Logger)
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Forced shutdown", "err", err)
	}

	return nil
}

func serve(serverConfig config.Server, server *http.Server, logger *slog.Logger) {
	if serverConfig.EnableHTTPS {
		logger.Info("Starting HTTPS server")

		certManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Cache:      autocert.DirCache("certs"),
			HostPolicy: autocert.HostWhitelist(serverConfig.ServerDomain),
		}

		server.Addr = ":" + fmt.Sprint(serverConfig.HTTPSPort)
		server.TLSConfig = &tls.Config{
			GetCertificate: certManager.GetCertificate,
		}

		// Serve HTTP redirect to HTTPS
		go func() {
			err := http.ListenAndServe(":"+fmt.Sprint(serverConfig.Port), certManager.HTTPHandler(nil))
			if err != nil {
				logger.Error("Could not start HTTP server", "err", err)
			}
		}()

		// Start HTTPS server
		err := server.ListenAndServeTLS("", "")
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Could not start HTTPS server", "err", err)
		}
		logger.Info("Listening on port " + fmt.Sprint(serverConfig.HTTPSPort) + " with HTTPS enabled")
	} else {
		logger.Info("Starting HTTP server")
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Could not start HTTP server", "err", err)
		}
	}
}
