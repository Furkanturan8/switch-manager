package models

import (
	"time"
)

// PortCreateRequest - Port oluşturma isteği
type PortCreateRequest struct {
	SwitchID    uint       `json:"switch_id" validate:"required"`
	Name        string     `json:"name" validate:"required,min=1,max=50"`
	Interface   string     `json:"interface" validate:"required,min=1,max=50"`
	Description string     `json:"description" validate:"max=255"`
	Speed       int        `json:"speed" validate:"min=0"`
	Mode        PortMode   `json:"mode" validate:"oneof=access trunk routed"`
	Duplex      DuplexMode `json:"duplex" validate:"oneof=full half auto"`
	MTU         int        `json:"mtu" validate:"min=64,max=9216"`
	AccessVLAN  int        `json:"access_vlan" validate:"min=1,max=4094"`
	TrunkVLANS  []int      `json:"trunk_vlans"`
	Poe         bool       `json:"poe"`
	MaxMAC      int        `json:"max_mac" validate:"min=0"`
}

// PortUpdateRequest - Port güncelleme isteği
type PortUpdateRequest struct {
	Name        string     `json:"name" validate:"omitempty,min=1,max=50"`
	Interface   string     `json:"interface" validate:"omitempty,min=1,max=50"`
	Description string     `json:"description" validate:"omitempty,max=255"`
	Speed       int        `json:"speed" validate:"omitempty,min=0"`
	Mode        PortMode   `json:"mode" validate:"omitempty,oneof=access trunk routed"`
	Duplex      DuplexMode `json:"duplex" validate:"omitempty,oneof=full half auto"`
	MTU         int        `json:"mtu" validate:"omitempty,min=64,max=9216"`
	AccessVLAN  int        `json:"access_vlan" validate:"omitempty,min=1,max=4094"`
	TrunkVLANS  []int      `json:"trunk_vlans"`
	Poe         *bool      `json:"poe,omitempty"`
	MaxMAC      *int       `json:"max_mac,omitempty" validate:"omitempty,min=0"`
	Status      string     `json:"status" validate:"required,oneof=up down disabled"`
	AdminStatus string     `json:"admin_status" validate:"required,oneof=up down"`
}

// PortResponse - Port yanıt modeli (detaylı)
type PortResponse struct {
	ID          uint   `json:"id"`
	SwitchID    uint   `json:"switch_id"`
	Name        string `json:"name"`
	Interface   string `json:"interface"`
	Description string `json:"description"`

	// Port Durumu
	Status      string     `json:"status"`
	AdminStatus string     `json:"admin_status"`
	OperStatus  string     `json:"oper_status"`
	Speed       int        `json:"speed"`
	Mode        PortMode   `json:"mode"`
	Duplex      DuplexMode `json:"duplex"`
	MTU         int        `json:"mtu"`
	MACAddress  string     `json:"mac_address"`

	// VLAN Bilgileri
	AccessVLAN int   `json:"access_vlan"`
	TrunkVLANS []int `json:"trunk_vlans"`

	// Ekstra Özellikler
	Poe    bool `json:"poe"`
	MaxMAC int  `json:"max_mac"`

	// Monitoring
	LastChange time.Time `json:"last_change"`
	ErrorCount int       `json:"error_count"`

	// Switch Bilgileri
	SwitchName string `json:"switch_name"`
	SwitchIP   string `json:"switch_ip"`

	// Zaman Damgaları
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PortListResponse - Port listesi yanıt modeli (kısaltılmış)
type PortListResponse struct {
	ID          uint      `json:"id"`
	SwitchID    uint      `json:"switch_id"`
	Name        string    `json:"name"`
	Interface   string    `json:"interface"`
	Status      string    `json:"status"`
	AdminStatus string    `json:"admin_status"`
	OperStatus  string    `json:"oper_status"`
	Speed       int       `json:"speed"`
	Mode        PortMode  `json:"mode"`
	AccessVLAN  int       `json:"access_vlan"`
	Poe         bool      `json:"poe"`
	SwitchName  string    `json:"switch_name"`
	SwitchIP    string    `json:"switch_ip"`
	LastChange  time.Time `json:"last_change"`
	CreatedAt   time.Time `json:"created_at"`
}

// PortListWithPagination - Sayfalama ile port listesi
type PortListWithPagination struct {
	Ports    []PortListResponse `json:"ports"`
	Total    int                `json:"total"`
	Page     int                `json:"page"`
	PageSize int                `json:"page_size"`
	Pages    int                `json:"pages"`
}

// ToModel - Request'i model'e dönüştür
func (req *PortCreateRequest) ToModel() *Port {
	return &Port{
		SwitchID:    req.SwitchID,
		Name:        req.Name,
		Interface:   req.Interface,
		Description: req.Description,
		Speed:       req.Speed,
		Mode:        req.Mode,
		Duplex:      req.Duplex,
		MTU:         req.MTU,
		AccessVLAN:  req.AccessVLAN,
		TrunkVLANS:  req.TrunkVLANS,
		Poe:         req.Poe,
		MaxMAC:      req.MaxMAC,
		Status:      "down", // Default status
		AdminStatus: "down", // Default admin status
		OperStatus:  "down", // Default oper status
	}
}

// ToModel - Update request'i mevcut model'e uygula
func (req *PortUpdateRequest) ToModel(existing *Port) {
	existing.Name = req.Name
	existing.Interface = req.Interface
	existing.Description = req.Description
	existing.Speed = req.Speed
	existing.Mode = req.Mode
	existing.Duplex = req.Duplex
	existing.MTU = req.MTU
	existing.AccessVLAN = req.AccessVLAN
	existing.TrunkVLANS = req.TrunkVLANS
	existing.Poe = *req.Poe
	existing.MaxMAC = *req.MaxMAC
	existing.Status = req.Status
	existing.AdminStatus = req.AdminStatus
}

// FromModel - Model'den response'a dönüştür
func FromPortModel(port *Port) *PortResponse {
	return &PortResponse{
		ID:          port.ID,
		SwitchID:    port.SwitchID,
		Name:        port.Name,
		Interface:   port.Interface,
		Description: port.Description,
		Status:      port.Status,
		AdminStatus: port.AdminStatus,
		OperStatus:  port.OperStatus,
		Speed:       port.Speed,
		Mode:        port.Mode,
		Duplex:      port.Duplex,
		MTU:         port.MTU,
		MACAddress:  port.MACAddress,
		AccessVLAN:  port.AccessVLAN,
		TrunkVLANS:  port.TrunkVLANS,
		Poe:         port.Poe,
		MaxMAC:      port.MaxMAC,
		LastChange:  port.LastChange,
		ErrorCount:  port.ErrorCount,
		SwitchName:  port.Switch.Name,
		SwitchIP:    port.Switch.IPAddress,
		CreatedAt:   port.CreatedAt,
		UpdatedAt:   port.UpdatedAt,
	}
}

// FromPortModelList - Model listesinden response listesine dönüştür
func FromPortModelList(ports []Port) []PortListResponse {
	responses := make([]PortListResponse, len(ports))
	for i, port := range ports {
		responses[i] = PortListResponse{
			ID:          port.ID,
			SwitchID:    port.SwitchID,
			Name:        port.Name,
			Interface:   port.Interface,
			Status:      port.Status,
			AdminStatus: port.AdminStatus,
			OperStatus:  port.OperStatus,
			Speed:       port.Speed,
			Mode:        port.Mode,
			AccessVLAN:  port.AccessVLAN,
			Poe:         port.Poe,
			SwitchName:  port.Switch.Name,
			SwitchIP:    port.Switch.IPAddress,
			LastChange:  port.LastChange,
			CreatedAt:   port.CreatedAt,
		}
	}
	return responses
}
