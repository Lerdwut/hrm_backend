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

// RequestLeave godoc
// @Summary Request leave
// @Description Submit a new leave request
// @Tags leaves
// @Accept json
// @Produce json
// @Param leave body domain.Leave true "Leave request data"
// @Success 200 {object} domain.Leave
// @Failure 400 {object} map[string]interface{}
// @Failure 500
// @Router /leaves/request [post]
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

// ApprovedLeave godoc
// @Summary Approve leave
// @Description Approve a leave request by ID
// @Tags leaves
// @Produce json
// @Param id path int true "Leave ID"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 500
// @Router /leaves/{id}/approve [put]
func (h *LeaveHandler) ApprovedLeave(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	return h.statusHandler(uint(id), domain.Approved, c)
}

// RejectedLeave godoc
// @Summary Reject leave
// @Description Reject a leave request by ID
// @Tags leaves
// @Produce json
// @Param id path int true "Leave ID"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 500
// @Router /leaves/{id}/reject [put]
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

// GetAllLeaves godoc
// @Summary Get all leaves
// @Description Get all leave requests
// @Tags leaves
// @Produce json
// @Success 200 {array} domain.Leave
// @Failure 500
// @Router /leaves/all [get]
func (h *LeaveHandler) GetAllLeaves(c *fiber.Ctx) error {
	leaves, err := h.s.GetAllLeaves()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(leaves)
}
