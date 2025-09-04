package repository

import (
	"context"
	"switch-manager/internal/models"
	"switch-manager/pkg/database"
)

type IPortRepository interface {
	Create(ctx context.Context, port_ *models.Port) error
	GetByID(ctx context.Context, id uint) (*models.Port, error)
	GetByName(ctx context.Context, name string) (*models.Port, error)
	GetAll(ctx context.Context) ([]models.Port, error)
	Update(ctx context.Context, port_ *models.Port) error
	Delete(ctx context.Context, id uint) error
}

type PortRepository struct {
	db *database.DB
}

func NewPortRepository(db *database.DB) IPortRepository {
	return &PortRepository{db: db}
}

// Create creates a new port
func (r *PortRepository) Create(ctx context.Context, port_ *models.Port) error {
	return r.db.WithContext(ctx).Create(port_).Error
}

// GetByID retrieves a port by ID
func (r *PortRepository) GetByID(ctx context.Context, id uint) (*models.Port, error) {
	var port_ models.Port
	err := r.db.WithContext(ctx).Preload("Switch").First(&port_, id).Error
	if err != nil {
		return nil, err
	}
	return &port_, nil
}

// GetByName retrieves a port by name
func (r *PortRepository) GetByName(ctx context.Context, name string) (*models.Port, error) {
	var port_ models.Port
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&port_).Error
	if err != nil {
		return nil, err
	}
	return &port_, nil
}

// GetAll retrieves all ports
func (r *PortRepository) GetAll(ctx context.Context) ([]models.Port, error) {
	var ports []models.Port
	err := r.db.WithContext(ctx).Preload("Switch").Find(&ports).Error
	return ports, err
}

// Update updates a port
func (r *PortRepository) Update(ctx context.Context, port_ *models.Port) error {
	return r.db.WithContext(ctx).Save(port_).Error
}

// Delete deletes a port by ID
func (r *PortRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Port{}, id).Error
}
