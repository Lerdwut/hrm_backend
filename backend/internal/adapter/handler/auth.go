package handler

import (
	"hr_management/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthService port.AuthService
}

func NewAuthHandler(authService port.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (ah *AuthHandler) LoginEndpoint(c *fiber.Ctx) error {
	var req loginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
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
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errMsgs})
	// }

	token, err := ah.AuthService.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	res := c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
	return c.Status(fiber.StatusOK).JSON(res)
}
