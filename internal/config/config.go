package config

import "github.com/ilyakaznacheev/cleanenv"

const (
	DevelopmentEnvironment = "dev"
	ProductionEnvironment  = "prod"
)

// Config holds all application configuration
type Config struct {
	Database Database
	Server   Server
	App      App
}

// App contains application-wide settings
type App struct {
	Environment string `yaml:"enviroment" env:"ENVIROMENT" env-default:"development"`
}

// Server contains HTTP/HTTPS server configuration
type Server struct {
	Port                 string   `yaml:"server_port" env:"SERVER_PORT" env-default:"8080"`
	EnableHTTPS          bool     `yaml:"enable_https" env:"ENABLE_HTTPS" env-default:"false"`
	HTTPSPort            string   `yaml:"https_server_port" env:"HTTPS_SERVER_PORT" env-default:"443"`
	TimeOutInSeconds     int      `yaml:"time_out_in_seconds" env:"TIME_OUT_IN_SECONDS" env-default:"60"`
	AllowedCorsURLs      []string `yaml:"allowed_cors_urls" env:"REACT_CLIENT_URLS"`
	TaskTimeOutInSeconds int      `yaml:"task_time_out_in_seconds" env:"TASK_TIME_OUT_IN_SECONDS" env-default:"3600"`
	ServerDomain         string   `yaml:"server_domain" env:"SERVER_DOMAIN"`
}

// Database contains database connection settings
type Database struct {
	TursoConnectionString string
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
