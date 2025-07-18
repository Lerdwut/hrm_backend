package handler

import (
	"hr_management/internal/core/domain"
	"hr_management/internal/core/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type LeaveHandler struct {
	s *service.LeaveService
}

func NewLeaveHandler(s *service.LeaveService) *LeaveHandler {
	return &LeaveHandler{s: s}
}

func (h *LeaveHandler) RequestLeave(c *fiber.Ctx) error {
	var req domain.Leave
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	err := h.s.RequestLeave(&req)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(req)
}

func (h *LeaveHandler) ApprovedLeave(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	return h.statusHandler(uint(id), domain.Approved, c)
}

func (h *LeaveHandler) RejectedLeave(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	return h.statusHandler(uint(id), domain.Rejected, c)
}

func (h *LeaveHandler) statusHandler(id uint, status domain.LeaveStatus, c *fiber.Ctx) error {
	var err error
	if status == domain.Approved {
		err = h.s.ApprovedLeave(id)
	} else {
		err = h.s.RejectedLeave(id)
	}
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}

func (h *LeaveHandler) GetAllLeaves(c *fiber.Ctx) error {
	leaves, err := h.s.GetAllLeaves()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(leaves)
}