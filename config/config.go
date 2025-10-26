package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL      string
	DatabaseNoSSLURL string
	AppEnv           string
	AppURI           string
}

var (
	cfg  *Config
	once sync.Once
)

const (
	/** Limit for tables/lists from db */
	DefaultListLimit = 50
)

func LoadConfig() *Config {
	once.Do(func() {
		root, err := findProjectRoot()
		if err != nil {
			log.Fatalf("failed to find project root: %v", err)
		}
		var envFileName string
		if IsTestEnv() {
			envFileName = ".env.test"
		} else {
			envFileName = ".env"
		}

		envPath := filepath.Join(root, envFileName)

		// Detect test mode
		envErr := godotenv.Load(envPath)
		if envErr != nil {
			log.Fatal(".env not found")
		}

		dbURL := os.Getenv("DATABASE_URL")
		if dbURL == "" {
			log.Fatal("DATABASE_URL is required")
		}

		dbNoSSLURL := os.Getenv("DATABASE_NO_SSL_URL")
		if dbURL == "" {
			log.Fatal("DATABASE_NO_SSL_URL is required")
		}

		appUri := os.Getenv("APP_URI")
		if appUri == "" {
			log.Fatal("APP_URI is required")
		}

		cfg = &Config{
			DatabaseURL:      dbURL,
			DatabaseNoSSLURL: dbNoSSLURL,
			AppURI:           appUri,
		}
	})
	return cfg
}

// Traverse upward to find project root (where go.mod is)
func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break // reached root
		}
		dir = parent
	}

	return "", os.ErrNotExist
}

func IsTestEnv() bool {
	isTestEnv := os.Getenv("APP_ENV") == "test"

	return isTestEnv
}
