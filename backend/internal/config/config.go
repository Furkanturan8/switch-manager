package config

import (
	"encoding/json"
	"os"
	"strconv"
)

type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
	SSH      SSHConfig      `json:"ssh"`
	Log      LogConfig      `json:"log"`
}

type ServerConfig struct {
	Port string `json:"port" env:"SERVER_PORT" envDefault:"8080"`
	Host string `json:"host" env:"SERVER_HOST" envDefault:"localhost"`
}

type DatabaseConfig struct {
	Host     string `json:"host" env:"DB_HOST" envDefault:"localhost"`
	Port     string `json:"port" env:"DB_PORT" envDefault:"5432"`
	Username string `json:"username" env:"DB_USERNAME" envDefault:"postgres"`
	Password string `json:"password" env:"DB_PASSWORD" envDefault:"postgres"`
	Database string `json:"database" env:"DB_NAME" envDefault:"switch_manager"`
	SSLMode  string `json:"ssl_mode" env:"DB_SSL_MODE" envDefault:"disable"`
}

type SSHConfig struct {
	Timeout     int    `json:"timeout" env:"SSH_TIMEOUT" envDefault:"30"`
	KeyPath     string `json:"key_path" env:"SSH_KEY_PATH" envDefault:"~/.ssh/id_rsa"`
	DefaultUser string `json:"default_user" env:"SSH_DEFAULT_USER" envDefault:"admin"`
}

type LogConfig struct {
	Level  string `json:"level" env:"LOG_LEVEL" envDefault:"info"`
	Format string `json:"format" env:"LOG_FORMAT" envDefault:"text"`
}

func Load() (*Config, error) {
	// Varsayılan konfigürasyon
	config := &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Host: getEnv("SERVER_HOST", "localhost"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			Username: getEnv("DB_USERNAME", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			Database: getEnv("DB_NAME", "switch_manager"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		SSH: SSHConfig{
			Timeout:     getEnvAsInt("SSH_TIMEOUT", 30),
			KeyPath:     getEnv("SSH_KEY_PATH", "~/.ssh/id_rsa"),
			DefaultUser: getEnv("SSH_DEFAULT_USER", "admin"),
		},
		Log: LogConfig{
			Level:  getEnv("LOG_LEVEL", "info"),
			Format: getEnv("LOG_FORMAT", "text"),
		},
	}

	// Konfigürasyon dosyası varsa yükle
	if _, err := os.Stat("config.json"); err == nil {
		file, err := os.Open("config.json")
		if err != nil {
			return nil, err
		}
		defer file.Close()

		if err := json.NewDecoder(file).Decode(config); err != nil {
			return nil, err
		}
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
