package models

import (
	"time"
)

// SwitchCreateRequest - Switch oluşturma isteği
type SwitchCreateRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=100"`
	IPAddress   string `json:"ip_address" validate:"required,ip"`
	Hostname    string `json:"hostname" validate:"max=255"`
	Model       string `json:"model" validate:"max=100"`
	Location    string `json:"location" validate:"max=255"`
	Description string `json:"description" validate:"max=500"`

	// Bağlantı Bilgileri
	SSHUsername string `json:"ssh_username" validate:"required,max=50"`
	SSHPassword string `json:"ssh_password" validate:"max=100"`
	SSHKeyPath  string `json:"ssh_key_path" validate:"max=255"`
	SSHPort     int    `json:"ssh_port" validate:"min=1,max=65535"`
}

// SwitchUpdateRequest - Switch güncelleme isteği
type SwitchUpdateRequest struct {
	Name        string `json:"name" validate:"omitempty,min=1,max=100"`
	IPAddress   string `json:"ip_address" validate:"omitempty,ip"`
	Hostname    string `json:"hostname" validate:"omitempty,max=255"`
	Model       string `json:"model" validate:"omitempty,max=100"`
	Location    string `json:"location" validate:"omitempty,max=255"`
	Description string `json:"description" validate:"omitempty,max=500"`
	Status      string `json:"status" validate:"required,oneof=online offline error"`

	// Bağlantı Bilgileri
	SSHUsername string `json:"ssh_username" validate:"omitempty,max=50"`
	SSHPassword string `json:"ssh_password" validate:"omitempty,max=100"`
	SSHKeyPath  string `json:"ssh_key_path" validate:"omitempty,max=255"`
	SSHPort     int    `json:"ssh_port" validate:"omitempty,min=1,max=65535"`
}

// SwitchResponse - Switch yanıt modeli (detaylı)
type SwitchResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	IPAddress   string `json:"ip_address"`
	Hostname    string `json:"hostname"`
	Model       string `json:"model"`
	Location    string `json:"location"`
	Description string `json:"description"`

	// Bağlantı Bilgileri (şifre hariç)
	SSHUsername string `json:"ssh_username"`
	SSHKeyPath  string `json:"ssh_key_path"`
	SSHPort     int    `json:"ssh_port"`

	// Durum Bilgileri
	Status     string    `json:"status"`
	LastSeen   time.Time `json:"last_seen"`
	IOSVersion string    `json:"ios_version"`

	// İstatistikler
	PortCount int `json:"port_count"`
	VLANCount int `json:"vlan_count"`

	// Zaman Damgaları
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SwitchListResponse - Switch listesi yanıt modeli (kısaltılmış)
type SwitchListResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	IPAddress string    `json:"ip_address"`
	Hostname  string    `json:"hostname"`
	Model     string    `json:"model"`
	Location  string    `json:"location"`
	Status    string    `json:"status"`
	LastSeen  time.Time `json:"last_seen"`
	PortCount int       `json:"port_count"`
	VLANCount int       `json:"vlan_count"`
	CreatedAt time.Time `json:"created_at"`
}

// SwitchListWithPagination - Sayfalama ile switch listesi
type SwitchListWithPagination struct {
	Switches []SwitchListResponse `json:"switches"`
	Total    int                  `json:"total"`
	Page     int                  `json:"page"`
	PageSize int                  `json:"page_size"`
	Pages    int                  `json:"pages"`
}

// ToModel - Request'i model'e dönüştür
func (req *SwitchCreateRequest) ToModel() *Switch {
	return &Switch{
		Name:        req.Name,
		IPAddress:   req.IPAddress,
		Hostname:    req.Hostname,
		Model:       req.Model,
		Location:    req.Location,
		Description: req.Description,
		SSHUsername: req.SSHUsername,
		SSHPassword: req.SSHPassword,
		SSHKeyPath:  req.SSHKeyPath,
		SSHPort:     req.SSHPort,
		Status:      "offline", // Default status
	}
}

// ToModel - Update request'i mevcut model'e uygula
func (req *SwitchUpdateRequest) ToModel(existing *Switch) {
	existing.Name = req.Name
	existing.IPAddress = req.IPAddress
	existing.Hostname = req.Hostname
	existing.Model = req.Model
	existing.Location = req.Location
	existing.Status = req.Status
	existing.Description = req.Description
	existing.SSHUsername = req.SSHUsername
	existing.SSHPassword = req.SSHPassword
	existing.SSHKeyPath = req.SSHKeyPath
	existing.SSHPort = req.SSHPort
	existing.UpdatedAt = time.Now()
}

// FromModel - Model'den response'a dönüştür
func FromSwitchModel(switch_ *Switch) *SwitchResponse {
	return &SwitchResponse{
		ID:          switch_.ID,
		Name:        switch_.Name,
		IPAddress:   switch_.IPAddress,
		Hostname:    switch_.Hostname,
		Model:       switch_.Model,
		Location:    switch_.Location,
		Description: switch_.Description,
		SSHUsername: switch_.SSHUsername,
		SSHKeyPath:  switch_.SSHKeyPath,
		SSHPort:     switch_.SSHPort,
		Status:      switch_.Status,
		LastSeen:    switch_.LastSeen,
		IOSVersion:  switch_.IOSVersion,
		PortCount:   len(switch_.Ports),
		VLANCount:   len(switch_.VLANS),
		CreatedAt:   switch_.CreatedAt,
		UpdatedAt:   switch_.UpdatedAt,
	}
}

// FromSwitchModelList - Model listesinden response listesine dönüştür
func FromSwitchModelList(switches []Switch) []SwitchListResponse {
	responses := make([]SwitchListResponse, len(switches))
	for i, switch_ := range switches {
		responses[i] = SwitchListResponse{
			ID:        switch_.ID,
			Name:      switch_.Name,
			IPAddress: switch_.IPAddress,
			Hostname:  switch_.Hostname,
			Model:     switch_.Model,
			Location:  switch_.Location,
			Status:    switch_.Status,
			LastSeen:  switch_.LastSeen,
			PortCount: len(switch_.Ports),
			VLANCount: len(switch_.VLANS),
			CreatedAt: switch_.CreatedAt,
		}
	}
	return responses
}
