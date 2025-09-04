package models

import (
	"time"
)

// VLANCreateRequest - VLAN oluşturma isteği
type VLANCreateRequest struct {
	SwitchID    uint   `json:"switch_id" validate:"required"`
	VLANID      int    `json:"vlan_id" validate:"required,min=1,max=4094"`
	Name        string `json:"name" validate:"required,min=1,max=100"`
	Description string `json:"description" validate:"max=255"`
	MTU         int    `json:"mtu" validate:"min=64,max=9216"`
	STPEnabled  bool   `json:"stp_enabled"`
	Priority    int    `json:"priority" validate:"min=0,max=65535"`
}

// VLANUpdateRequest - VLAN güncelleme isteği
type VLANUpdateRequest struct {
	Name        string `json:"name" validate:"omitempty,min=1,max=100"`
	Description string `json:"description" validate:"omitempty,max=255"`
	MTU         int    `json:"mtu" validate:"omitempty,min=64,max=9216"`
	STPEnabled  *bool  `json:"stp_enabled,omitempty"`
	Priority    int    `json:"priority" validate:"omitempty,min=0,max=65535"`
}

// VLANStatusUpdateRequest - VLAN durum güncelleme isteği
type VLANStatusUpdateRequest struct {
	AdminStatus string `json:"admin_status" validate:"required,oneof=enabled disabled"`
	Status      string `json:"status" validate:"required,oneof=active suspended"`
}

// VLANResponse - VLAN yanıt modeli (detaylı)
type VLANResponse struct {
	ID          uint   `json:"id"`
	SwitchID    uint   `json:"switch_id"`
	VLANID      int    `json:"vlan_id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	// Durum
	AdminStatus string `json:"admin_status"`
	OperStatus  string `json:"oper_status"`
	Status      string `json:"status"`

	// VLAN özellikleri
	MTU        int  `json:"mtu"`
	STPEnabled bool `json:"stp_enabled"`
	Priority   int  `json:"priority"`

	// İstatistikler
	PortCount int `json:"port_count"`

	// Switch Bilgileri
	SwitchName string `json:"switch_name"`
	SwitchIP   string `json:"switch_ip"`

	// Zaman Damgaları
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// VLANListResponse - VLAN listesi yanıt modeli (kısaltılmış)
type VLANListResponse struct {
	ID          uint      `json:"id"`
	SwitchID    uint      `json:"switch_id"`
	VLANID      int       `json:"vlan_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	AdminStatus string    `json:"admin_status"`
	OperStatus  string    `json:"oper_status"`
	Status      string    `json:"status"`
	PortCount   int       `json:"port_count"`
	SwitchName  string    `json:"switch_name"`
	SwitchIP    string    `json:"switch_ip"`
	CreatedAt   time.Time `json:"created_at"`
}

// VLANListWithPagination - Sayfalama ile VLAN listesi
type VLANListWithPagination struct {
	VLANS    []VLANListResponse `json:"vlans"`
	Total    int                `json:"total"`
	Page     int                `json:"page"`
	PageSize int                `json:"page_size"`
	Pages    int                `json:"pages"`
}

// VLANStatusResponse - VLAN durum güncelleme yanıtı
type VLANStatusResponse struct {
	Message     string `json:"message"`
	AdminStatus string `json:"admin_status"`
	Status      string `json:"status"`
}

// ToModel - Request'i model'e dönüştür
func (req *VLANCreateRequest) ToModel() *VLAN {
	return &VLAN{
		SwitchID:    req.SwitchID,
		VLANID:      req.VLANID,
		Name:        req.Name,
		Description: req.Description,
		MTU:         req.MTU,
		STPEnabled:  req.STPEnabled,
		Priority:    req.Priority,
		AdminStatus: "enabled", // Default admin status
		OperStatus:  "down",    // Default oper status
		Status:      "active",  // Default status
	}
}

// ToModel - Update request'i mevcut model'e uygula
func (req *VLANUpdateRequest) ToModel(existing *VLAN) {
	if req.Name != "" {
		existing.Name = req.Name
	}
	if req.Description != "" {
		existing.Description = req.Description
	}
	if req.MTU != 0 {
		existing.MTU = req.MTU
	}
	if req.STPEnabled != nil {
		existing.STPEnabled = *req.STPEnabled
	}
	if req.Priority != 0 {
		existing.Priority = req.Priority
	}
}

// FromModel - Model'den response'a dönüştür
func FromVLANModel(vlan *VLAN) *VLANResponse {
	return &VLANResponse{
		ID:          vlan.ID,
		SwitchID:    vlan.SwitchID,
		VLANID:      vlan.VLANID,
		Name:        vlan.Name,
		Description: vlan.Description,
		AdminStatus: vlan.AdminStatus,
		OperStatus:  vlan.OperStatus,
		Status:      vlan.Status,
		MTU:         vlan.MTU,
		STPEnabled:  vlan.STPEnabled,
		Priority:    vlan.Priority,
		PortCount:   len(vlan.Ports),
		SwitchName:  vlan.Switch.Name,
		SwitchIP:    vlan.Switch.IPAddress,
		CreatedAt:   vlan.CreatedAt,
		UpdatedAt:   vlan.UpdatedAt,
	}
}

// FromVLANModelList - Model listesinden response listesine dönüştür
func FromVLANModelList(vlans []VLAN) []VLANListResponse {
	responses := make([]VLANListResponse, len(vlans))
	for i, vlan := range vlans {
		responses[i] = VLANListResponse{
			ID:          vlan.ID,
			SwitchID:    vlan.SwitchID,
			VLANID:      vlan.VLANID,
			Name:        vlan.Name,
			Description: vlan.Description,
			AdminStatus: vlan.AdminStatus,
			OperStatus:  vlan.OperStatus,
			Status:      vlan.Status,
			PortCount:   len(vlan.Ports),
			SwitchName:  vlan.Switch.Name,
			SwitchIP:    vlan.Switch.IPAddress,
			CreatedAt:   vlan.CreatedAt,
		}
	}
	return responses
}
