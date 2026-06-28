package config

import "github.com/ilyakaznacheev/cleanenv"

const (
	DevelopmentEnvironment = "dev"
	ProductionEnvironment  = "prod"
)

// Config holds all application configuration
type Config struct {
	Database Database `toml:"Database"`
	Server   Server   `toml:"Server"`
	App      App      `toml:"App"`
}

// App contains application-wide settings
type App struct {
	Environment string `toml:"Environment" env:"ENVIRONMENT" env-default:"dev"`
}

// Server contains HTTP/HTTPS server configuration
type Server struct {
	Port                 string   `toml:"Port" env:"SERVER_PORT" env-default:"8080"`
	EnableHTTPS          bool     `toml:"EnableHTTPS" env:"ENABLE_HTTPS" env-default:"false"`
	HTTPSPort            string   `toml:"HTTPSPort" env:"HTTPS_SERVER_PORT" env-default:"443"`
	AllowedCorsURLs      []string `toml:"AllowedCorsURLs" env:"ALLOWED_CORS_URLS"`
	TaskTimeOutInSeconds int      `toml:"TaskTimeOutInSeconds" env:"TASK_TIMEOUT_IN_SECONDS" env-default:"3600"`
	ServerDomain         string   `toml:"ServerDomain" env:"SERVER_DOMAIN"`
}

// Database contains database connection settings
type Database struct {
	TursoConnectionString string `toml:"TursoConnectionString" env:"TURSO_CONNECTION_STRING"`
}

// Load reads configuration from the specified file path
// and returns a populated Config struct or an error
func Load(filename string) (*Config, error) {
	var config Config
	err := cleanenv.ReadConfig(filename, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
