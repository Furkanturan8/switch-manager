package service

import (
	"switch-manager/internal/api/repository"
	"switch-manager/internal/models"
	"switch-manager/pkg/errorx"
	"switch-manager/pkg/logger"
)

type SwitchService struct {
	repo   *repository.SwitchRepository
	logger logger.Logger
}

func NewSwitchService(repo *repository.SwitchRepository) *SwitchService {
	return &SwitchService{
		repo:   repo,
		logger: logger.New(),
	}
}

// CreateSwitch creates a new switch
func (s *SwitchService) CreateSwitch(switch_ *models.Switch) error {
	s.logger.Info("Creating switch:", switch_.Name, "IP:", switch_.IPAddress)

	// Check if switch with same IP already exists
	existing, err := s.repo.GetByIP(switch_.IPAddress)
	if err == nil && existing != nil {
		return errorx.New(errorx.ErrDuplicate.Code, "Switch with this IP address already exists")
	}

	err = s.repo.Create(switch_)
	if err != nil {
		s.logger.Error("Failed to create switch:", err)
		return err
	}

	s.logger.Info("Switch created successfully:", switch_.Name)
	return nil
}

// GetSwitch retrieves a switch by ID
func (s *SwitchService) GetSwitch(id uint) (*models.Switch, error) {
	s.logger.Info("Getting switch by ID:", id)

	switch_, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Error("Failed to get switch:", err)
		return nil, err
	}

	return switch_, nil
}

// GetAllSwitches retrieves all switches
func (s *SwitchService) GetAllSwitches() ([]models.Switch, error) {
	s.logger.Info("Getting all switches")

	switches, err := s.repo.GetAll()
	if err != nil {
		s.logger.Error("Failed to get switches:", err)
		return nil, err
	}

	s.logger.Info("Retrieved ", len(switches), " switches")
	return switches, nil
}

// UpdateSwitch updates a switch
func (s *SwitchService) UpdateSwitch(switch_ *models.Switch) error {
	s.logger.Info("Updating switch:", switch_.Name)

	err := s.repo.Update(switch_)
	if err != nil {
		s.logger.Error("Failed to update switch:", err)
		return err
	}

	s.logger.Info("Switch updated successfully:", switch_.Name)
	return nil
}

// DeleteSwitch deletes a switch
func (s *SwitchService) DeleteSwitch(id uint) error {
	s.logger.Info("Deleting switch with ID:", id)

	err := s.repo.Delete(id)
	if err != nil {
		s.logger.Error("Failed to delete switch:", err)
		return err
	}

	s.logger.Info("Switch deleted successfully:", id)
	return nil
}
