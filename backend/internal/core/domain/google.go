package domain

import (
	"time"
)

type OAuthProvider string

const (
	ProviderGoogle OAuthProvider = "google"
	// ProviderLocal  OAuthProvider = "local"
)

type GoogleUser struct {
	ID       string        `json:"id"`
	Email    string        `json:"email"`
	Name     string        `json:"name"`
	Picture  string        `json:"picture"`
	Provider OAuthProvider `json:"provider"`
	Verified bool          `json:"verified"`
}

type GoogleState struct {
	State     string    `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
