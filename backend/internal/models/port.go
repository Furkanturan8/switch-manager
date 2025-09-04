package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type PortMode string
type DuplexMode string

const (
	ModeAccess PortMode = "access"
	ModeTrunk  PortMode = "trunk"
	ModeRouted PortMode = "routed"

	DuplexFull DuplexMode = "full"
	DuplexHalf DuplexMode = "half"
	DuplexAuto DuplexMode = "auto"
)

type Port struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	SwitchID  uint   `json:"switch_id" gorm:"not null"`
	Name      string `json:"name" gorm:"not null"`      // GigabitEthernet0/1
	Interface string `json:"interface" gorm:"not null"` // Gi0/1, Fa0/1

	// Port Durumu
	Status      string     `json:"status" gorm:"default:'down'"`       // up, down, disabled
	AdminStatus string     `json:"admin_status" gorm:"default:'down'"` // up, down (config)
	OperStatus  string     `json:"oper_status" gorm:"default:'down'"`  // gerçek durum
	Speed       int        `json:"speed" gorm:"default:0"`             // 0=auto, aksi halde 10,100,1000,...
	Mode        PortMode   `json:"mode" gorm:"type:text;default:'access'"`
	Duplex      DuplexMode `json:"duplex" gorm:"type:text;default:'auto'"`
	MTU         int        `json:"mtu" gorm:"default:1500"` // MTU değeri: bir porttan iletilebilecek tek seferdeki en büyük paket boyutudur (byte cinsinden).
	MACAddress  string     `json:"mac_address"`
	Description string     `json:"description"`

	// VLAN Bilgileri
	AccessVLAN int   `json:"access_vlan"`
	TrunkVLANS []int `json:"trunk_vlans" gorm:"serializer:json"`

	// Ekstra Özellikler
	Poe    bool `json:"poe" gorm:"default:false"`
	MaxMAC int  `json:"max_mac" gorm:"default:0"` // port-security kaç cihaz bağlanacak? Max bağlantı sayısı (default olarak 0:sınırsız)

	// Monitoring
	LastChange time.Time `json:"last_change"`
	ErrorCount int       `json:"error_count" gorm:"default:0"`

	// İlişkiler
	Switch Switch `json:"switch" gorm:"foreignKey:SwitchID"`

	// Zaman Damgaları
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for Switch
func (Port) TableName() string {
	return "ports"
}

// BeforeCreate GORM hook
func (p *Port) BeforeCreate(tx *gorm.DB) (err error) {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	// Validate çağrısı
	if err = p.Validate(); err != nil {
		return err
	}
	return
}

// BeforeUpdate GORM hook
func (p *Port) BeforeUpdate(tx *gorm.DB) (err error) {
	p.UpdatedAt = time.Now()
	// Validate çağrısı
	if err = p.Validate(); err != nil {
		return err
	}
	return
}

func (p *Port) Validate() error {
	switch p.Mode {
	case ModeAccess, ModeTrunk, ModeRouted:
		// geçerli
	default:
		return fmt.Errorf("invalid Mode: %s", p.Mode)
	}

	switch p.Duplex {
	case DuplexFull, DuplexHalf, DuplexAuto:
		// geçerli
	default:
		return fmt.Errorf("invalid Duplex: %s", p.Duplex)
	}

	return nil
}
