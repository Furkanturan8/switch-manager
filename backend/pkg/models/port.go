package models

import (
	"time"

	"gorm.io/gorm"
)

type Port struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	SwitchID  uint   `json:"switch_id" gorm:"not null"`
	Name      string `json:"name" gorm:"not null"`      // GigabitEthernet0/1
	Interface string `json:"interface" gorm:"not null"` // Gi0/1, Fa0/1

	// Port Durumu
	AdminStatus string `json:"admin_status" gorm:"default:'down'"` // up, down (config)
	OperStatus  string `json:"oper_status" gorm:"default:'down'"`  // gerçek durum
	Mode        string `json:"mode" gorm:"type:enum('access','trunk','routed');default:'access'"`
	Speed       int    `json:"speed" gorm:"default:0"` // 0=auto, aksi halde 10,100,1000,...
	Duplex      string `json:"duplex" gorm:"type:enum('full','half','auto');default:'auto'"`
	MTU         int    `json:"mtu" gorm:"default:1500"` // MTU değeri: bir porttan iletilebilecek tek seferdeki en büyük paket boyutudur (byte cinsinden).
	MACAddress  string `json:"mac_address"`
	Description string `json:"description"`

	// VLAN Bilgileri
	AccessVLAN int   `json:"access_vlan"`
	TrunkVLANs []int `json:"trunk_vlans" gorm:"serializer:json"`

	// Ekstra Özellikler
	PoE    bool `json:"poe" gorm:"default:false"`
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
	return
}

// BeforeUpdate GORM hook
func (p *Port) BeforeUpdate(tx *gorm.DB) (err error) {
	p.UpdatedAt = time.Now()
	return
}
