package repository

import (
	"switch-manager/pkg/database"
	"switch-manager/pkg/models"

	"gorm.io/gorm"
)

type SwitchRepository struct {
	db *database.DB
}

func NewSwitchRepository(db *database.DB) *SwitchRepository {
	return &SwitchRepository{db: db}
}

// Create creates a new switch
func (r *SwitchRepository) Create(switch_ *models.Switch) error {
	return r.db.Create(switch_).Error
}

// GetByID retrieves a switch by ID
func (r *SwitchRepository) GetByID(id uint) (*models.Switch, error) {
	var switch_ models.Switch
	err := r.db.First(&switch_, id).Error
	if err != nil {
		return nil, err
	}
	return &switch_, nil
}

// GetByIP retrieves a switch by IP address
func (r *SwitchRepository) GetByIP(ip string) (*models.Switch, error) {
	var switch_ models.Switch
	err := r.db.Where("ip_address = ?", ip).First(&switch_).Error
	if err != nil {
		return nil, err
	}
	return &switch_, nil
}

// GetAll retrieves all switches
func (r *SwitchRepository) GetAll() ([]models.Switch, error) {
	var switches []models.Switch
	err := r.db.Find(&switches).Error
	return switches, err
}

// Update updates a switch
func (r *SwitchRepository) Update(switch_ *models.Switch) error {
	return r.db.Save(switch_).Error
}

// Delete deletes a switch by ID
func (r *SwitchRepository) Delete(id uint) error {
	return r.db.Delete(&models.Switch{}, id).Error
}

// UpdateStatus updates switch status
func (r *SwitchRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&models.Switch{}).Where("id = ?", id).Update("status", status).Error
}

// UpdateLastSeen updates switch last seen timestamp
func (r *SwitchRepository) UpdateLastSeen(id uint) error {
	return r.db.Model(&models.Switch{}).Where("id = ?", id).Update("last_seen", gorm.Expr("NOW()")).Error
}
