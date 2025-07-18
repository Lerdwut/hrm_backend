package handler

import (
	"hr_management/internal/core/domain"
	"hr_management/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	return &UserHandler{userService}
}

type registerRequest struct {
	Username             string `json:"username" validate:"required"`
	Email                string `json:"email,omitempty" validate:"required,email"`
	Password             string `json:"password,omitempty" validate:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation,omitempty" validate:"required,min=8"`
}

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

	user := domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	_, err := h.userService.Register(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	res := c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
	return c.Status(fiber.StatusOK).JSON(res)
}
