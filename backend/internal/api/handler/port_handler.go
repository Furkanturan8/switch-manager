package handler

import (
	"strconv"
	"switch-manager/internal/api/service"
	"switch-manager/pkg/errorx"
	"switch-manager/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type PortHandler struct {
	service *service.PortService
}

func NewPortHandler(service *service.PortService) *PortHandler {
	return &PortHandler{service: service}
}

func (h *PortHandler) CreatePort(c *fiber.Ctx) error {
	var port_ models.Port
	if err := c.BodyParser(&port_); err != nil {
		return errorx.WrapErr(errorx.ErrInvalidRequest, err)
	}

	if err := h.service.CreatePort(&port_); err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	return c.Status(fiber.StatusCreated).JSON(port_)
}

// GetPort handles getting a port by ID
func (h *PortHandler) GetPort(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return errorx.WrapMsg(errorx.ErrInvalidRequest, "Invalid port ID")
	}

	port_, err := h.service.GetPort(uint(id))
	if err != nil {
		return errorx.WrapMsg(errorx.ErrNotFound, "Port not found")
	}

	return c.JSON(port_)
}

// GetAllPortes handles getting all ports
func (h *PortHandler) GetAllPortes(c *fiber.Ctx) error {
	ports, err := h.service.GetAllPorts()
	if err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	return c.JSON(fiber.Map{
		"ports": ports,
		"count": len(ports),
	})
}

// UpdatePort handles port update
func (h *PortHandler) UpdatePort(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return errorx.WrapMsg(errorx.ErrInvalidRequest, "Invalid port ID")
	}

	var port_ models.Port
	if err = c.BodyParser(&port_); err != nil {
		return errorx.WrapErr(errorx.ErrInvalidRequest, err)
	}

	port_.ID = uint(id)

	if err = h.service.UpdatePort(&port_); err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	return c.JSON(port_)
}

// DeletePort handles port deletion
func (h *PortHandler) DeletePort(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return errorx.WrapMsg(errorx.ErrInvalidRequest, "Invalid port ID")
	}

	if err = h.service.DeletePort(uint(id)); err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// UpdatePortStatus handles port status update
// todo: gerek var mı??
func (h *PortHandler) UpdatePortStatus(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return errorx.WrapMsg(errorx.ErrInvalidRequest, "Invalid port ID")
	}

	var req struct {
		Status string `json:"status"`
	}

	if err = c.BodyParser(&req); err != nil {
		return errorx.WrapErr(errorx.ErrInvalidRequest, err)
	}

	if err = h.service.UpdatePortStatus(uint(id), req.Status); err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	return c.JSON(fiber.Map{
		"message": "Port status updated successfully",
		"status":  req.Status,
	})
}
