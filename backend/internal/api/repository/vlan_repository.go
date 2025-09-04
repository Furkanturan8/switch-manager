package repository

import (
	"context"
	"switch-manager/internal/models"
	"switch-manager/pkg/database"
)

type IVLANRepository interface {
	Create(ctx context.Context, vlan *models.VLAN) error
	GetByID(ctx context.Context, id uint) (*models.VLAN, error)
	GetByVLANID(ctx context.Context, switchID uint, vlanID int) (*models.VLAN, error)
	GetAll(ctx context.Context) ([]models.VLAN, error)
	GetBySwitchID(ctx context.Context, switchID uint) ([]models.VLAN, error)
	Update(ctx context.Context, vlan *models.VLAN) error
	Delete(ctx context.Context, id uint) error
}

type VLANRepository struct {
	db *database.DB
}

func NewVLANRepository(db *database.DB) IVLANRepository {
	return &VLANRepository{db: db}
}

// Create creates a new VLAN
func (r *VLANRepository) Create(ctx context.Context, vlan *models.VLAN) error {
	return r.db.WithContext(ctx).Create(vlan).Error
}

// GetByID retrieves a VLAN by ID
func (r *VLANRepository) GetByID(ctx context.Context, id uint) (*models.VLAN, error) {
	var vlan models.VLAN
	err := r.db.WithContext(ctx).Preload("Switch").First(&vlan, id).Error
	if err != nil {
		return nil, err
	}
	return &vlan, nil
}

// GetByVLANID retrieves a VLAN by switch ID and VLAN ID
func (r *VLANRepository) GetByVLANID(ctx context.Context, switchID uint, vlanID int) (*models.VLAN, error) {
	var vlan models.VLAN
	err := r.db.WithContext(ctx).Where("switch_id = ? AND vlan_id = ?", switchID, vlanID).First(&vlan).Error
	if err != nil {
		return nil, err
	}
	return &vlan, nil
}

// GetAll retrieves all VLANs
func (r *VLANRepository) GetAll(ctx context.Context) ([]models.VLAN, error) {
	var vlans []models.VLAN
	err := r.db.WithContext(ctx).Preload("Switch").Find(&vlans).Error
	return vlans, err
}

// GetBySwitchID retrieves all VLANs for a specific switch
func (r *VLANRepository) GetBySwitchID(ctx context.Context, switchID uint) ([]models.VLAN, error) {
	var vlans []models.VLAN
	err := r.db.WithContext(ctx).Where("switch_id = ?", switchID).Find(&vlans).Error
	return vlans, err
}

// Update updates a VLAN
func (r *VLANRepository) Update(ctx context.Context, vlan *models.VLAN) error {
	return r.db.WithContext(ctx).Save(vlan).Error
}

// Delete deletes a VLAN by ID
func (r *VLANRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.VLAN{}, id).Error
}
