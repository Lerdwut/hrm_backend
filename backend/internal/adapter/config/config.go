package config

import "os"

type (
	Container struct {
		DB     DB
		Google Google
		Server Server
	}

	DB struct {
		Host     string
		Port     string
		Username string
		Password string
		DBName   string
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
)

func Load() *Container {
	return &Container{
		DB: DB{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			Username: getEnv("DB_USERNAME", "root"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "hrm_db"),
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
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
