package handler

import (
	"strconv"
	"switch-manager/internal/api/service"
	"switch-manager/pkg/errorx"
	"switch-manager/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type SwitchHandler struct {
	service *service.SwitchService
}

func NewSwitchHandler(service *service.SwitchService) *SwitchHandler {
	return &SwitchHandler{service: service}
}

// CreateSwitch handles switch creation
func (h *SwitchHandler) CreateSwitch(c *fiber.Ctx) error {
	var switch_ models.Switch

	if err := c.BodyParser(&switch_); err != nil {
		return errorx.WrapErr(errorx.ErrInvalidRequest, err)
	}

	if err := h.service.CreateSwitch(&switch_); err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	return c.Status(fiber.StatusCreated).JSON(switch_)
}

// GetSwitch handles getting a switch by ID
func (h *SwitchHandler) GetSwitch(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return errorx.WrapMsg(errorx.ErrInvalidRequest, "Invalid switch ID")
	}

	switch_, err := h.service.GetSwitch(uint(id))
	if err != nil {
		return errorx.WrapMsg(errorx.ErrNotFound, "Switch not found")
	}

	return c.JSON(switch_)
}

// GetAllSwitches handles getting all switches
func (h *SwitchHandler) GetAllSwitches(c *fiber.Ctx) error {
	switches, err := h.service.GetAllSwitches()
	if err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	return c.JSON(fiber.Map{
		"switches": switches,
		"count":    len(switches),
	})
}

// UpdateSwitch handles switch update
func (h *SwitchHandler) UpdateSwitch(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return errorx.WrapMsg(errorx.ErrInvalidRequest, "Invalid switch ID")
	}

	var switch_ models.Switch
	if err = c.BodyParser(&switch_); err != nil {
		return errorx.WrapErr(errorx.ErrInvalidRequest, err)
	}

	switch_.ID = uint(id)

	if err = h.service.UpdateSwitch(&switch_); err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	return c.JSON(switch_)
}

// DeleteSwitch handles switch deletion
func (h *SwitchHandler) DeleteSwitch(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return errorx.WrapMsg(errorx.ErrInvalidRequest, "Invalid switch ID")
	}

	if err = h.service.DeleteSwitch(uint(id)); err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// UpdateSwitchStatus handles switch status update
func (h *SwitchHandler) UpdateSwitchStatus(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return errorx.WrapMsg(errorx.ErrInvalidRequest, "Invalid switch ID")
	}

	var req struct {
		Status string `json:"status"`
	}

	if err = c.BodyParser(&req); err != nil {
		return errorx.WrapErr(errorx.ErrInvalidRequest, err)
	}

	if err = h.service.UpdateSwitchStatus(uint(id), req.Status); err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	return c.JSON(fiber.Map{
		"message": "Switch status updated successfully",
		"status":  req.Status,
	})
}
