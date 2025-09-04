package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
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
	// .env dosyasını yükle
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	// Konfigürasyonu oluştur
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

	return config, nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsBool(key string, defaultVal bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultVal
}
