package repository

import (
	"context"
	"switch-manager/internal/models"
	"switch-manager/pkg/database"

	"gorm.io/gorm"
)

type ISwitchRepository interface {
	Create(ctx context.Context, switch_ *models.Switch) error
	GetByID(ctx context.Context, id uint) (*models.Switch, error)
	GetByIP(ctx context.Context, ip string) (*models.Switch, error)
	GetAll(ctx context.Context) ([]models.Switch, error)
	Update(ctx context.Context, switch_ *models.Switch) error
	Delete(ctx context.Context, id uint) error
	UpdateLastSeen(ctx context.Context, id uint) error
}

type SwitchRepository struct {
	db *database.DB
}

func NewSwitchRepository(db *database.DB) ISwitchRepository {
	return &SwitchRepository{db: db}
}

// Create creates a new switch
func (r *SwitchRepository) Create(ctx context.Context, switch_ *models.Switch) error {
	return r.db.WithContext(ctx).Create(switch_).Error
}

// GetByID retrieves a switch by ID
func (r *SwitchRepository) GetByID(ctx context.Context, id uint) (*models.Switch, error) {
	var switch_ models.Switch
	err := r.db.WithContext(ctx).First(&switch_, id).Error
	if err != nil {
		return nil, err
	}
	return &switch_, nil
}

// GetByIP retrieves a switch by IP address
func (r *SwitchRepository) GetByIP(ctx context.Context, ip string) (*models.Switch, error) {
	var switch_ models.Switch
	err := r.db.WithContext(ctx).Where("ip_address = ?", ip).First(&switch_).Error
	if err != nil {
		return nil, err
	}
	return &switch_, nil
}

// GetAll retrieves all switches
func (r *SwitchRepository) GetAll(ctx context.Context) ([]models.Switch, error) {
	var switches []models.Switch
	err := r.db.WithContext(ctx).Find(&switches).Error
	return switches, err
}

// Update updates a switch
func (r *SwitchRepository) Update(ctx context.Context, switch_ *models.Switch) error {
	return r.db.WithContext(ctx).Save(switch_).Error
}

// Delete deletes a switch by ID
func (r *SwitchRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Switch{}, id).Error
}

// UpdateLastSeen updates switch last seen timestamp
func (r *SwitchRepository) UpdateLastSeen(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(&models.Switch{}).Where("id = ?", id).Update("last_seen", gorm.Expr("NOW()")).Error
}
