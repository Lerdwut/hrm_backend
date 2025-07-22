package handler

import (
	"hr_management/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type GoogleHandler struct {
	googleService port.GoogleService
}

func NewGoogleHandler(googleService port.GoogleService) *GoogleHandler {
	return &GoogleHandler{
		googleService: googleService,
	}
}

// @Summary Google OAuth Login
// @Description Redirect to Google OAuth consent screen
// @Tags auth
// @Accept json
// @Produce json
// @Success 302 {string} string "Redirect to Google"
// @Router /auth/google/login [get]
func (gh *GoogleHandler) GoogleLogin(c *fiber.Ctx) error {
	state := gh.googleService.GenerateState()
	authURL := gh.googleService.GetAuthURL(state)

	return c.Redirect(authURL, fiber.StatusTemporaryRedirect)
}

// @Summary Google OAuth Callback
// @Description Handle callback from Google OAuth
// @Tags auth
// @Accept json
// @Produce json
// @Param code query string true "Authorization code from Google"
// @Param state query string true "State parameter for security"
// @Success 200 {object} map[string]interface{} "User info and success message"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/google/callback [get]
func (gh *GoogleHandler) GoogleCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	state := c.Query("state")

	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "authorization code is required",
		})
	}

	if state == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "state parameter is required",
		})
	}

	user, err := gh.googleService.HandleCallback(code, state)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"avatar":   user.Avatar,
			"provider": user.Provider,
		},
	})
}
