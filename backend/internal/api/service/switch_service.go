package service

import (
	"context"
	"switch-manager/internal/api/repository"
	"switch-manager/internal/models"
	"switch-manager/pkg/errorx"
	"switch-manager/pkg/logger"
)

type SwitchService struct {
	repo   repository.ISwitchRepository
	logger logger.Logger
}

func NewSwitchService(repo repository.ISwitchRepository) *SwitchService {
	return &SwitchService{
		repo:   repo,
		logger: logger.New(),
	}
}

// CreateSwitch creates a new switch
func (s *SwitchService) CreateSwitch(ctx context.Context, switch_ *models.Switch) error {
	s.logger.Info("Creating switch:", switch_.Name, "IP:", switch_.IPAddress)

	// Check if switch with same IP already exists
	existing, err := s.repo.GetByIP(ctx, switch_.IPAddress)
	if err == nil && existing != nil {
		return errorx.New(errorx.ErrDuplicate.Code, "Switch with this IP address already exists")
	}

	err = s.repo.Create(ctx, switch_)
	if err != nil {
		s.logger.Error("Failed to create switch:", err)
		return err
	}

	s.logger.Info("Switch created successfully:", switch_.Name)
	return nil
}

// GetSwitch retrieves a switch by ID
func (s *SwitchService) GetSwitch(ctx context.Context, id uint) (*models.Switch, error) {
	s.logger.Info("Getting switch by ID:", id)

	switch_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to get switch:", err)
		return nil, err
	}

	return switch_, nil
}

// GetAllSwitches retrieves all switches
func (s *SwitchService) GetAllSwitches(ctx context.Context) ([]models.Switch, error) {
	s.logger.Info("Getting all switches")

	switches, err := s.repo.GetAll(ctx)
	if err != nil {
		s.logger.Error("Failed to get switches:", err)
		return nil, err
	}

	s.logger.Info("Retrieved ", len(switches), " switches")
	return switches, nil
}

// UpdateSwitch updates a switch
func (s *SwitchService) UpdateSwitch(ctx context.Context, switch_ *models.Switch) error {
	s.logger.Info("Updating switch:", switch_.Name)

	err := s.repo.Update(ctx, switch_)
	if err != nil {
		s.logger.Error("Failed to update switch:", err)
		return err
	}

	s.logger.Info("Switch updated successfully:", switch_.Name)
	return nil
}

// DeleteSwitch deletes a switch
func (s *SwitchService) DeleteSwitch(ctx context.Context, id uint) error {
	s.logger.Info("Deleting switch with ID:", id)

	err := s.repo.Delete(ctx, id)
	if err != nil {
		s.logger.Error("Failed to delete switch:", err)
		return err
	}

	s.logger.Info("Switch deleted successfully:", id)
	return nil
}
