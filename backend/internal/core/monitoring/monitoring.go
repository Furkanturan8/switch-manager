package monitoring

import (
	"switch-manager/internal/config"
	"switch-manager/pkg/database"
	"switch-manager/pkg/logger"
	"time"
)

type Monitor struct {
	config *config.Config
	db     *database.DB
	ticker *time.Ticker
	logger logger.Logger
}

func New(cfg *config.Config, db *database.DB) *Monitor {
	return &Monitor{
		config: cfg,
		db:     db,
		ticker: time.NewTicker(30 * time.Second),
		logger: logger.New(),
	}
}

func (m *Monitor) Start() {
	m.logger.Info("Monitoring başlatılıyor...")

	for range m.ticker.C {
		m.checkSwitches()
		m.collectMetrics()
	}
}

func (m *Monitor) checkSwitches() {
	// Switch durumlarını kontrol et
	// Bu fonksiyon daha sonra implement edilecek
	m.logger.Debug("Switch durumları kontrol ediliyor...")
}

func (m *Monitor) collectMetrics() {
	// Performans metriklerini topla
	// Bu fonksiyon daha sonra implement edilecek
	m.logger.Debug("Performans metrikleri toplanıyor...")
}

func (m *Monitor) Stop() {
	m.logger.Info("Monitoring durduruluyor...")
	m.ticker.Stop()
}

func (m *Monitor) GetStatus() string {
	return "running"
}
