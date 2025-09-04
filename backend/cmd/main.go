package main

import (
	"errors"
	"log"
	"switch-manager/internal/api/handler"
	"switch-manager/internal/api/repository"
	"switch-manager/internal/api/service"
	"switch-manager/internal/config"
	"switch-manager/internal/core/monitoring"
	"switch-manager/internal/models"
	"switch-manager/pkg/database"
	"switch-manager/pkg/errorx"
	"switch-manager/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Logger'ı başlat
	logger := logger.New()
	logger.Info("Switch Manager başlatılıyor...")

	// Konfigürasyonu yükle
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Konfigürasyon yüklenemedi:", err)
	}

	// Veritabanı bağlantısını kur
	db, err := database.Connect(cfg.Database)
	if err != nil {
		logger.Fatal("Veritabanı bağlantısı kurulamadı:", err)
	}
	defer database.Close()

	// Veritabanı migration'ını çalıştır
	if err = database.AutoMigrate(&models.Switch{}, &models.Port{}, models.VLAN{}); err != nil {
		logger.Fatal("Veritabanı migration hatası:", err)
	}
	logger.Info("Veritabanı migration tamamlandı")

	// Monitoring'i başlat
	monitor := monitoring.New(cfg, db)
	go monitor.Start()

	// Fiber app'i oluştur
	app := fiber.New(fiber.Config{
		AppName: "Switch Manager",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Custom error handling
			if appErr, ok := err.(*errorx.AppError); ok {
				return errorx.WrapErr(appErr, err)
			}

			// Default Fiber error handling
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			return errorx.New(code, "Internal Server Error", err)
		},
	})

	// Middleware'leri ekle
	app.Use(recover.New())

	// Custom logging middleware
	app.Use(func(c *fiber.Ctx) error {
		logger.Info("Request: ", c.Method(), " ", c.Path(), " ", "IP: ", c.IP())
		return c.Next()
	})

	// Route'ları tanımla
	setupRoutes(app, db)

	// HTTP sunucusunu başlat
	logger.Info("HTTP sunucusu başlatılıyor... port:", cfg.Server.Port)
	if err = app.Listen(":" + cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}

func setupRoutes(app *fiber.App, db *database.DB) {
	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Repository ve service'leri oluştur
	switchRepo := repository.NewSwitchRepository(db)
	switchService := service.NewSwitchService(switchRepo)
	switchHandler := handler.NewSwitchHandler(switchService)

	portRepo := repository.NewPortRepository(db)
	portService := service.NewPortService(portRepo)
	portHandler := handler.NewPortHandler(portService)

	// API v1
	api := app.Group("/api/v1")
	{
		// Switch yönetimi
		switches := api.Group("/switches")
		switches.Get("/", switchHandler.GetAllSwitches)
		switches.Post("/", switchHandler.CreateSwitch)
		switches.Get("/:id", switchHandler.GetSwitch)
		switches.Put("/:id", switchHandler.UpdateSwitch)
		switches.Delete("/:id", switchHandler.DeleteSwitch)

		// Port yönetimi (placeholder)
		ports := api.Group("/ports")
		ports.Get("/", portHandler.GetAllPortes)
		ports.Post("/", portHandler.CreatePort)
		ports.Get("/:id", portHandler.GetPort)
		ports.Put("/:id", portHandler.UpdatePort)
		ports.Delete("/:id", portHandler.DeletePort)

		// VLAN yönetimi (placeholder)
		api.Get("/vlans", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{"message": "VLAN listesi gelecek"})
		})
	}
}
