package handler

import (
	"strconv"
	"switch-manager/internal/api/service"
	"switch-manager/internal/models"
	"switch-manager/pkg/errorx"

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
	var req models.SwitchCreateRequest

	if err := c.BodyParser(&req); err != nil {
		return errorx.WrapErr(errorx.ErrInvalidRequest, err)
	}

	switch_ := req.ToModel()
	if err := h.service.CreateSwitch(switch_); err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	response := models.FromSwitchModel(switch_)
	return c.Status(fiber.StatusCreated).JSON(response)
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

	response := models.FromSwitchModel(switch_)
	return c.JSON(response)
}

// GetAllSwitches handles getting all switches
func (h *SwitchHandler) GetAllSwitches(c *fiber.Ctx) error {
	switches, err := h.service.GetAllSwitches()
	if err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	response := models.FromSwitchModelList(switches)
	return c.JSON(fiber.Map{
		"switches": response,
		"count":    len(response),
	})
}

// UpdateSwitch handles switch update
func (h *SwitchHandler) UpdateSwitch(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return errorx.WrapMsg(errorx.ErrInvalidRequest, "Invalid switch ID")
	}

	// Önce mevcut switch'i getir
	switch_, err := h.service.GetSwitch(uint(id))
	if err != nil {
		return errorx.WrapMsg(errorx.ErrNotFound, "Switch not found")
	}

	var req models.SwitchUpdateRequest
	if err = c.BodyParser(&req); err != nil {
		return errorx.WrapErr(errorx.ErrInvalidRequest, err)
	}

	// Sadece gönderilen alanları güncelle
	req.ToModel(switch_)
	if err = h.service.UpdateSwitch(switch_); err != nil {
		return errorx.WrapErr(errorx.ErrInternal, err)
	}

	response := models.FromSwitchModel(switch_)
	return c.JSON(response)
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
