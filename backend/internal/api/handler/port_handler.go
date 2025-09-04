package handler

import (
	"strconv"
	"switch-manager/internal/api/service"
	"switch-manager/internal/models"
	"switch-manager/pkg/errorx"

	"github.com/gofiber/fiber/v2"
)

type PortHandler struct {
	service *service.PortService
}

func NewPortHandler(service *service.PortService) *PortHandler {
	return &PortHandler{service: service}
}

func (h *PortHandler) CreatePort(c *fiber.Ctx) error {
	var req models.PortCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return errorx.WrapErr(errorx.ErrInvalidRequest, err)
	}

	port_ := req.ToModel()
	if err := h.service.CreatePort(c.Context(), port_); err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	response := models.FromPortModel(port_)
	return c.Status(fiber.StatusCreated).JSON(response)
}

// GetPort handles getting a port by ID
func (h *PortHandler) GetPort(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return errorx.WrapMsg(errorx.ErrInvalidRequest, "Invalid port ID")
	}

	port_, err := h.service.GetPort(c.Context(), uint(id))
	if err != nil {
		return errorx.WrapMsg(errorx.ErrNotFound, "Port not found")
	}

	response := models.FromPortModel(port_)
	return c.JSON(response)
}

// GetAllPortes handles getting all ports
func (h *PortHandler) GetAllPortes(c *fiber.Ctx) error {
	ports, err := h.service.GetAllPorts(c.Context())
	if err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	response := models.FromPortModelList(ports)
	return c.JSON(fiber.Map{
		"ports": response,
		"count": len(response),
	})
}

// UpdatePort handles port update
func (h *PortHandler) UpdatePort(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return errorx.WrapMsg(errorx.ErrInvalidRequest, "Invalid port ID")
	}

	// Önce mevcut port'u getir
	port_, err := h.service.GetPort(c.Context(), uint(id))
	if err != nil {
		return errorx.WrapMsg(errorx.ErrNotFound, "Port not found")
	}

	var req models.PortUpdateRequest
	if err = c.BodyParser(&req); err != nil {
		return errorx.WrapErr(errorx.ErrInvalidRequest, err)
	}

	// Sadece gönderilen alanları güncelle
	req.ToModel(port_)
	if err = h.service.UpdatePort(c.Context(), port_); err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	response := models.FromPortModel(port_)
	return c.JSON(response)
}

// DeletePort handles port deletion
func (h *PortHandler) DeletePort(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return errorx.WrapMsg(errorx.ErrInvalidRequest, "Invalid port ID")
	}

	if err = h.service.DeletePort(c.Context(), uint(id)); err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
