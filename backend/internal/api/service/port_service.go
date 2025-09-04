package service

import (
	"switch-manager/internal/api/repository"
	"switch-manager/internal/models"
	"switch-manager/pkg/errorx"
	"switch-manager/pkg/logger"
)

type PortService struct {
	repo   *repository.PortRepository
	logger logger.Logger
}

func NewPortService(repo *repository.PortRepository) *PortService {
	return &PortService{
		repo:   repo,
		logger: logger.New(),
	}
}

// CreatePort creates a new port
func (s *PortService) CreatePort(port_ *models.Port) error {
	s.logger.Info("Creating port:", port_.Name)

	// Check if port with same name already exists
	existing, err := s.repo.GetByName(port_.Name)
	if err == nil && existing != nil {
		return errorx.New(errorx.ErrDuplicate.Code, "Port with this name already exists")
	}

	err = s.repo.Create(port_)
	if err != nil {
		s.logger.Error("Failed to create port:", err)
		return err
	}

	s.logger.Info("Port created successfully:", port_.Name)
	return nil
}

// GetPort retrieves a port by ID
func (s *PortService) GetPort(id uint) (*models.Port, error) {
	s.logger.Info("Getting port by ID:", id)

	port_, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Error("Failed to get port:", err)
		return nil, err
	}

	return port_, nil
}

// GetAllPorts retrieves all ports
func (s *PortService) GetAllPorts() ([]models.Port, error) {
	s.logger.Info("Getting all ports")

	ports, err := s.repo.GetAll()
	if err != nil {
		s.logger.Error("Failed to get port:", err)
		return nil, err
	}

	s.logger.Info("Retrieved ", len(ports), " port(s)")
	return ports, nil
}

// UpdatePort updates a port
func (s *PortService) UpdatePort(port_ *models.Port) error {
	s.logger.Info("Updating port:", port_.Name)

	err := s.repo.Update(port_)
	if err != nil {
		s.logger.Error("Failed to update port:", err)
		return err
	}

	s.logger.Info("Port updated successfully:", port_.Name)
	return nil
}

// DeletePort deletes a port
func (s *PortService) DeletePort(id uint) error {
	s.logger.Info("Deleting port with ID:", id)

	err := s.repo.Delete(id)
	if err != nil {
		s.logger.Error("Failed to delete port:", err)
		return err
	}

	s.logger.Info("Port deleted successfully:", id)
	return nil
}
