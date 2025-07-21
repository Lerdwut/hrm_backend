package config

import "os"

type (
	Container struct {
		DB     DB
		Google Google
		Server Server
		HTTP   HTTP
	}

	DB struct {
		Host     string
		Port     string
		Database string
		Collection string
	}

	Google struct {
		ClientID     string
		ClientSecret string
		RedirectURL  string
		Scopes       []string
	}

	Server struct {
		Port string
		Host string
	}

	HTTP struct {
		Env            string
		URL            string
		Port           string
		AllowedOrigins string
		Prefix         string
	}
)

func Load() (*Container, error) {
	return &Container{
		DB: DB{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "27017"),
			Database: getEnv("DB_DATABASE", "hexagonal-v1-mongo"),
			Collection: getEnv("DB_COLLECTION", ""),
		},
		Google: Google{
			ClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
			ClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
			RedirectURL:  getEnv("GOOGLE_REDIRECT_URL", "http://localhost:8080/auth/google/callback"),
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
		},
		Server: Server{
			Port: getEnv("SERVER_PORT", "8080"),
			Host: getEnv("SERVER_HOST", "localhost"),
		},
		HTTP: HTTP{
			Env:            getEnv("APP_ENV", ""),
			URL:            getEnv("HTTP_URL", ""),
			Port:           getEnv("HTTP_PORT", ""),
			AllowedOrigins: getEnv("HTTP_ALLOWED_ORIGINS", "http://localhost:3000"),
			Prefix:         getEnv("HTTP_PREFIX", "/api"),
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
