package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadReadsTomlConfig(t *testing.T) {
	configPath := writeConfig(t, `
[Server]
Port = "9090"
EnableHTTPS = true
HTTPSPort = "9443"
AllowedCorsURLs = ["http://localhost:3000", "https://example.com"]
TaskTimeOutInSeconds = 42
ServerDomain = "example.com"

[App]
Environment = "prod"

[Database]
TursoConnectionString = "libsql://example.turso.io?authToken=test-token"
`)

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	if cfg.App.Environment != ProductionEnvironment {
		t.Fatalf("Environment = %q, want %q", cfg.App.Environment, ProductionEnvironment)
	}
	if cfg.Server.Port != "9090" {
		t.Fatalf("Port = %q, want 9090", cfg.Server.Port)
	}
	if !cfg.Server.EnableHTTPS {
		t.Fatal("EnableHTTPS = false, want true")
	}
	if cfg.Server.HTTPSPort != "9443" {
		t.Fatalf("HTTPSPort = %q, want 9443", cfg.Server.HTTPSPort)
	}
	if len(cfg.Server.AllowedCorsURLs) != 2 {
		t.Fatalf("AllowedCorsURLs length = %d, want 2", len(cfg.Server.AllowedCorsURLs))
	}
	if cfg.Server.TaskTimeOutInSeconds != 42 {
		t.Fatalf("TaskTimeOutInSeconds = %d, want 42", cfg.Server.TaskTimeOutInSeconds)
	}
	if cfg.Server.ServerDomain != "example.com" {
		t.Fatalf("ServerDomain = %q, want example.com", cfg.Server.ServerDomain)
	}
	if cfg.Database.TursoConnectionString == "" {
		t.Fatal("TursoConnectionString is empty")
	}
}

func TestLoadUsesEnvironmentOverride(t *testing.T) {
	t.Setenv("ENVIRONMENT", ProductionEnvironment)

	configPath := writeConfig(t, `
[App]
Environment = "dev"
`)

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	if cfg.App.Environment != ProductionEnvironment {
		t.Fatalf("Environment = %q, want %q", cfg.App.Environment, ProductionEnvironment)
	}
}

func writeConfig(t *testing.T, contents string) string {
	t.Helper()

	configPath := filepath.Join(t.TempDir(), "config.toml")
	if err := os.WriteFile(configPath, []byte(contents), 0o600); err != nil {
		t.Fatalf("write config: %v", err)
	}

	return configPath
}
