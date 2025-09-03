package repository

import (
	"switch-manager/pkg/database"
	"switch-manager/pkg/models"
)

type PortRepository struct {
	db *database.DB
}

func NewPortRepository(db *database.DB) *PortRepository {
	return &PortRepository{db: db}
}

// Create creates a new port
func (r *PortRepository) Create(port_ *models.Port) error {
	return r.db.Create(port_).Error
}

// GetByID retrieves a port by ID
func (r *PortRepository) GetByID(id uint) (*models.Port, error) {
	var port_ models.Port
	err := r.db.First(&port_, id).Error
	if err != nil {
		return nil, err
	}
	return &port_, nil
}

// GetByName retrieves a port by name
func (r *PortRepository) GetByName(name string) (*models.Port, error) {
	var port_ models.Port
	err := r.db.Where("name = ?", name).First(&port_).Error
	if err != nil {
		return nil, err
	}
	return &port_, nil
}

// GetAll retrieves all ports
func (r *PortRepository) GetAll() ([]models.Port, error) {
	var ports []models.Port
	err := r.db.Find(&ports).Error
	return ports, err
}

// Update updates a port
func (r *PortRepository) Update(port_ *models.Port) error {
	return r.db.Save(port_).Error
}

// Delete deletes a port by ID
func (r *PortRepository) Delete(id uint) error {
	return r.db.Delete(&models.Port{}, id).Error
}

// UpdateAdminStatus updates port admin status
func (r *PortRepository) UpdateAdminStatus(id uint, status string) error {
	return r.db.Model(&models.Port{}).Where("id = ?", id).Update("admin_status", status).Error
}

// UpdateOperStatus updates port operational status
func (r *PortRepository) UpdateOperStatus(id uint, status string) error {
	return r.db.Model(&models.Port{}).Where("id = ?", id).Update("oper_status", status).Error
}

// UpdateStatus updates port status
func (r *PortRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&models.Port{}).Where("id = ?", id).Update("status", status).Error
}
