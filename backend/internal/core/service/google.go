package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"hr_management/internal/core/domain"
	"hr_management/internal/core/port"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type googleService struct {
	config       *oauth2.Config
	userService  port.UserService
	stateStorage map[string]*domain.GoogleState
}

func NewGoogleService(userService port.UserService) port.GoogleService {
	return &googleService{
		config: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
			Scopes:       []string{"openid", "profile", "email"},
			Endpoint:     google.Endpoint,
		},
		userService:  userService,
		stateStorage: make(map[string]*domain.GoogleState),
	}
}

func (gs *googleService) GetAuthURL(state string) string {
	gs.config.Scopes = []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	}

	return gs.config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (gs *googleService) HandleCallback(code, state string) (*domain.User, error) {
	// Validate state
	if err := gs.ValidateState(state); err != nil {
		return nil, err
	}

	// Exchange code for token
	token, err := gs.config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code for token: %v", err)
	}

	// Get user info from Google
	oauthUser, err := gs.getUserInfo(token)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %v", err)
	}

	// Find or create user
	user, err := gs.userService.FindByEmail(oauthUser.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}

	if user == nil {
		// Create new user
		user, err = gs.userService.CreateGoogleUser(oauthUser)
		if err != nil {
			return nil, fmt.Errorf("failed to create user: %v", err)
		}
	} else {
		// Update existing user
		user, err = gs.userService.UpdateGoogleUser(user, oauthUser)
		if err != nil {
			return nil, fmt.Errorf("failed to update user: %v", err)
		}
	}

	return user, nil
}

func (gs *googleService) ValidateState(state string) error {
	storedState, exists := gs.stateStorage[state]
	if !exists {
		return fmt.Errorf("invalid state parameter")
	}

	if time.Now().After(storedState.ExpiresAt) {
		delete(gs.stateStorage, state)
		return fmt.Errorf("state parameter expired")
	}

	// Remove used state
	delete(gs.stateStorage, state)
	return nil
}

func (gs *googleService) GenerateState() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	state := hex.EncodeToString(b)

	// Store the state with a timeout
	gs.stateStorage[state] = &domain.GoogleState{
		State:     state,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(10 * time.Minute), // Add expiration time
	}

	return state
}

func (gs *googleService) getUserInfo(token *oauth2.Token) (*domain.GoogleUser, error) {
	client := gs.config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get user info: status code %d", resp.StatusCode)
	}

	var googleUser domain.GoogleUser
	if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %v", err)
	}

	googleUser.Provider = domain.ProviderGoogle
	return &googleUser, nil
}
