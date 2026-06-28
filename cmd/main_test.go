package main

import (
	"context"
	"io"
	"log/slog"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/mhpenta/starterA/internal/app"
	"github.com/mhpenta/starterA/internal/config"
	httphandlers "github.com/mhpenta/starterA/internal/handlers/http"
)

func TestRunServerReturnsListenError(t *testing.T) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("listen on test port: %v", err)
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	err = runServer(
		ctx,
		config.Server{
			Port:                 strconv.Itoa(port),
			TaskTimeOutInSeconds: 1,
		},
		&app.Application{Logger: logger},
		httphandlers.New(nil, logger),
	)
	if err == nil {
		t.Fatal("expected listen error, got nil")
	}
	if !strings.Contains(err.Error(), "could not start HTTP server") {
		t.Fatalf("expected HTTP startup error, got %v", err)
	}
}
