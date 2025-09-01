package models

import (
	"time"
)

// Switch represents a network switch device
type Switch struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"uniqueIndex;not null"`
	IPAddress   string `json:"ip_address" gorm:"uniqueIndex;not null"`
	Hostname    string `json:"hostname"`
	Model       string `json:"model"`
	Location    string `json:"location"`
	Description string `json:"description"`

	// Bağlantı Bilgileri
	SSHUsername string `json:"ssh_username"`
	SSHPassword string `json:"ssh_password" gorm:"-"`
	SSHKeyPath  string `json:"ssh_key_path"`
	SSHPort     int    `json:"ssh_port" gorm:"default:22"`

	// Durum Bilgileri
	Status     string    `json:"status" gorm:"default:'offline'"` // online, offline, error
	LastSeen   time.Time `json:"last_seen"`
	IOSVersion string    `json:"ios_version"`

	// todo: ilişkiler eklenecek (örneğin, VLAN'lar, portlar, backups, alerts)

	// Zaman Damgaları
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for Switch
func (Switch) TableName() string {
	return "switches"
}

// BeforeCreate GORM hook
func (s *Switch) BeforeCreate() error {
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate GORM hook
func (s *Switch) BeforeUpdate() error {
	s.UpdatedAt = time.Now()
	return nil
}
