package handler

import (
	"hr_management/internal/core/domain"
	"hr_management/internal/core/port"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	return &UserHandler{userService}
}

// RegisterRequest represents the request body for user registration
type registerRequest struct {
	Username             string `json:"username" validate:"required" example:"john_doe"`
	Email                string `json:"email,omitempty" validate:"required,email" example:"john@example.com"`
	Password             string `json:"password,omitempty" validate:"required,min=8" example:"password123"`
	PasswordConfirmation string `json:"password_confirmation,omitempty" validate:"required,min=8" example:"password123"`
}

// RegisterResponse represents the response for user registration
type RegisterResponse struct {
	Message string `json:"message" example:"User registered successfully"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid request body"`
}

// RegisterEndpoint godoc
// @Summary Register a new user
// @Description Register a new user account
// @Tags users
// @Accept json
// @Produce json
// @Param user body registerRequest true "User registration data"
// @Success 201 {object} RegisterResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/register [post]
func (h *UserHandler) RegisterEndpoint(c *fiber.Ctx) error {
	var req registerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// myValidator := &Xvalidator{validator: validator.New()}

	// if errs := myValidator.ValidateRequest(req); len(errs) > 0 {
	// 	errMsgs := make([]string, 0)

	// 	for _, err := range errs {
	// 		errMsgs = append(errMsgs, fmt.Sprintf(
	// 			"[%s]: '%v' | Needs to implement '%s'",
	// 			err.FailedField,
	// 			err.Value,
	// 			err.Tag,
	// 		))
	// 	}
	// 	errors := newErrorResponse(errMsgs)
	// 	return c.Status(fiber.StatusBadRequest).JSON(errors)
	// }

	if req.Password != req.PasswordConfirmation {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Password confirmation does not match"})
	}

	// uuid := uuid.New().String()

	user := domain.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := h.userService.Register(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully", "user": user})
}

func (h *UserHandler) GetUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	if username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username is required"})
	}

	user, err := h.userService.FindByUsername(username)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve user"})
	}

	return c.JSON(user)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.userService.ListUsers()
	if err != nil {
		if err == domain.ErrDataNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No users found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve users"})
	}

	return c.JSON(users)
}
