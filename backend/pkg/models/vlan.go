package models

import (
	"time"

	"gorm.io/gorm"
)

type VLAN struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	SwitchID    uint   `json:"switch_id" gorm:"not null;index:idx_switch_vlan,unique"`
	VLANID      int    `json:"vlan_id" gorm:"not null;check:vlan_id > 0 AND vlan_id <= 4094;index:idx_switch_vlan,unique"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`

	// Durum
	AdminStatus string `json:"admin_status" gorm:"default:'enabled'"` // enabled, disabled
	OperStatus  string `json:"oper_status" gorm:"default:'down'"`     // up, down
	Status      string `json:"status" gorm:"default:'active'"`        // active, suspended

	// VLAN özellikleri
	MTU        int  `json:"mtu" gorm:"default:1500"`
	STPEnabled bool `json:"stp_enabled" gorm:"default:true"`
	Priority   int  `json:"priority" gorm:"default:32768"`

	// İlişkiler
	Switch Switch `json:"switch" gorm:"foreignKey:SwitchID"`
	Ports  []Port `json:"ports" gorm:"many2many:port_vlans;"`

	// Zaman Damgaları
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for VLAN
func (VLAN) TableName() string { return "vlans" }

// BeforeCreate GORM hook
func (v *VLAN) BeforeCreate(tx *gorm.DB) (err error) {
	v.CreatedAt = time.Now()
	v.UpdatedAt = time.Now()
	return
}

// BeforeUpdate GORM hook
func (v *VLAN) BeforeUpdate(tx *gorm.DB) (err error) {
	v.UpdatedAt = time.Now()
	return
}
