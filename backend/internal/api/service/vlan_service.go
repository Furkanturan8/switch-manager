package service

import (
	"context"
	"switch-manager/internal/api/repository"
	"switch-manager/internal/models"
	"switch-manager/pkg/errorx"
	"switch-manager/pkg/logger"
)

type VLANService struct {
	repo   repository.IVLANRepository
	logger logger.Logger
}

func NewVLANService(repo repository.IVLANRepository) *VLANService {
	return &VLANService{
		repo:   repo,
		logger: logger.New(),
	}
}

// CreateVLAN creates a new VLAN
func (s *VLANService) CreateVLAN(ctx context.Context, vlan *models.VLAN) error {
	s.logger.Info("Creating VLAN:", vlan.Name, "VLAN ID:", vlan.VLANID)

	// Check if VLAN with same ID already exists on the switch
	existing, err := s.repo.GetByVLANID(ctx, vlan.SwitchID, vlan.VLANID)
	if err == nil && existing != nil {
		return errorx.New(errorx.ErrDuplicate.Code, "VLAN with this ID already exists on this switch")
	}

	err = s.repo.Create(ctx, vlan)
	if err != nil {
		s.logger.Error("Failed to create VLAN:", err)
		return err
	}

	s.logger.Info("VLAN created successfully:", vlan.Name)
	return nil
}

// GetVLAN retrieves a VLAN by ID
func (s *VLANService) GetVLAN(ctx context.Context, id uint) (*models.VLAN, error) {
	s.logger.Info("Getting VLAN by ID:", id)

	vlan, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to get VLAN:", err)
		return nil, err
	}

	return vlan, nil
}

// GetAllVLANs retrieves all VLANs
func (s *VLANService) GetAllVLANs(ctx context.Context) ([]models.VLAN, error) {
	s.logger.Info("Getting all VLANs")

	vlans, err := s.repo.GetAll(ctx)
	if err != nil {
		s.logger.Error("Failed to get VLANs:", err)
		return nil, err
	}

	s.logger.Info("Retrieved ", len(vlans), " VLAN(s)")
	return vlans, nil
}

// GetVLANsBySwitchID retrieves all VLANs for a specific switch
func (s *VLANService) GetVLANsBySwitchID(ctx context.Context, switchID uint) ([]models.VLAN, error) {
	s.logger.Info("Getting VLANs for switch ID:", switchID)

	vlans, err := s.repo.GetBySwitchID(ctx, switchID)
	if err != nil {
		s.logger.Error("Failed to get VLANs for switch:", err)
		return nil, err
	}

	s.logger.Info("Retrieved ", len(vlans), " VLAN(s) for switch:", switchID)
	return vlans, nil
}

// UpdateVLAN updates a VLAN
func (s *VLANService) UpdateVLAN(ctx context.Context, vlan *models.VLAN) error {
	s.logger.Info("Updating VLAN:", vlan.Name)

	err := s.repo.Update(ctx, vlan)
	if err != nil {
		s.logger.Error("Failed to update VLAN:", err)
		return err
	}

	s.logger.Info("VLAN updated successfully:", vlan.Name)
	return nil
}

// DeleteVLAN deletes a VLAN
func (s *VLANService) DeleteVLAN(ctx context.Context, id uint) error {
	s.logger.Info("Deleting VLAN with ID:", id)

	err := s.repo.Delete(ctx, id)
	if err != nil {
		s.logger.Error("Failed to delete VLAN:", err)
		return err
	}

	s.logger.Info("VLAN deleted successfully:", id)
	return nil
}
