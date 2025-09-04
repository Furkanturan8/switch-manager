package handler

import (
	"strconv"
	"switch-manager/internal/api/service"
	"switch-manager/internal/models"
	"switch-manager/pkg/errorx"

	"github.com/gofiber/fiber/v2"
)

type VLANHandler struct {
	service *service.VLANService
}

func NewVLANHandler(service *service.VLANService) *VLANHandler {
	return &VLANHandler{service: service}
}

// CreateVLAN handles VLAN creation
func (h *VLANHandler) CreateVLAN(c *fiber.Ctx) error {
	var req models.VLANCreateRequest

	if err := c.BodyParser(&req); err != nil {
		return errorx.WrapErr(errorx.ErrInvalidRequest, err)
	}

	vlan := req.ToModel()
	if err := h.service.CreateVLAN(c.Context(), vlan); err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	response := models.FromVLANModel(vlan)
	return c.Status(fiber.StatusCreated).JSON(response)
}

// GetVLAN handles getting a VLAN by ID
func (h *VLANHandler) GetVLAN(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return errorx.WrapMsg(errorx.ErrInvalidRequest, "Invalid VLAN ID")
	}

	vlan, err := h.service.GetVLAN(c.Context(), uint(id))
	if err != nil {
		return errorx.WrapMsg(errorx.ErrNotFound, "VLAN not found")
	}

	response := models.FromVLANModel(vlan)
	return c.JSON(response)
}

// GetAllVLANs handles getting all VLANs
func (h *VLANHandler) GetAllVLANs(c *fiber.Ctx) error {
	vlans, err := h.service.GetAllVLANs(c.Context())
	if err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	response := models.FromVLANModelList(vlans)
	return c.JSON(fiber.Map{
		"vlans": response,
		"count": len(response),
	})
}

// GetVLANsBySwitchID handles getting VLANs by switch ID
func (h *VLANHandler) GetVLANsBySwitchID(c *fiber.Ctx) error {
	switchID, err := strconv.ParseUint(c.Params("switch_id"), 10, 32)
	if err != nil {
		return errorx.WrapMsg(errorx.ErrInvalidRequest, "Invalid switch ID")
	}

	vlans, err := h.service.GetVLANsBySwitchID(c.Context(), uint(switchID))
	if err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	response := models.FromVLANModelList(vlans)
	return c.JSON(fiber.Map{
		"vlans": response,
		"count": len(response),
	})
}

// UpdateVLAN handles VLAN update
func (h *VLANHandler) UpdateVLAN(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return errorx.WrapMsg(errorx.ErrInvalidRequest, "Invalid VLAN ID")
	}

	// Önce mevcut VLAN'ı getir
	vlan, err := h.service.GetVLAN(c.Context(), uint(id))
	if err != nil {
		return errorx.WrapMsg(errorx.ErrNotFound, "VLAN not found")
	}

	var req models.VLANUpdateRequest
	if err = c.BodyParser(&req); err != nil {
		return errorx.WrapErr(errorx.ErrInvalidRequest, err)
	}

	// Sadece gönderilen alanları güncelle
	req.ToModel(vlan)
	if err = h.service.UpdateVLAN(c.Context(), vlan); err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	response := models.FromVLANModel(vlan)
	return c.JSON(response)
}

// DeleteVLAN handles VLAN deletion
func (h *VLANHandler) DeleteVLAN(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return errorx.WrapMsg(errorx.ErrInvalidRequest, "Invalid VLAN ID")
	}

	if err = h.service.DeleteVLAN(c.Context(), uint(id)); err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
