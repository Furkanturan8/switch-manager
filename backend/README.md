# Switch Manager Backend

Cisco switch'ler için Go tabanlı network otomasyon sistemi.

## Özellikler

- **SSH Bağlantı Yönetimi**: Switch'lere güvenli SSH bağlantısı
- **Konfigürasyon Otomasyonu**: YAML template'ler ile otomatik konfigürasyon
- **Port Tracking**: Port durumlarının gerçek zamanlı takibi
- **Backup & Versioning**: Git destekli konfigürasyon yedekleme
- **Monitoring**: SNMP tabanlı performans izleme
- **REST API**: Modern web API ile entegrasyon

## Kurulum

### Gereksinimler

- Go 1.21+
- PostgreSQL 15+
- Docker & Docker Compose (opsiyonel)

### Yerel Kurulum

1. **Bağımlılıkları indir:**
```bash
go mod download
```

2. **Environment dosyasını oluştur:**
```bash
cp .env.example .env
# .env dosyasını düzenle
```

3. **Veritabanını başlat:**
```bash
docker-compose up -d postgres
```

4. **Uygulamayı çalıştır:**
```bash
go run cmd/main.go
```

### Docker ile Kurulum

```bash
docker-compose up -d
```

## API Endpoints

### Health Check
- `GET /health` - Servis durumu

### API v1
- `GET /api/v1/switches` - Switch listesi
- `GET /api/v1/ports` - Port listesi
- `GET /api/v1/vlans` - VLAN listesi

## Proje Yapısı

```
backend/
├── cmd/                    # Ana uygulama giriş noktaları
├── internal/              # Proje içi paketler
│   ├── config/           # Konfigürasyon yönetimi
│   ├── ssh/              # SSH bağlantı yönetimi
│   ├── backup/           # Backup & versioning modülü
│   ├── porttracker/      # Port takip modülü
│   ├── versiontracker/   # IOS versiyon takip modülü
│   ├── usersync/         # Kullanıcı senkronizasyon modülü
│   └── monitoring/       # Monitoring agent modülü
├── pkg/                  # Dışa açık paketler
│   ├── models/           # Veri modelleri
│   ├── utils/            # Yardımcı fonksiyonlar
│   ├── logger/           # Loglama sistemi
│   └── database/         # Veritabanı bağlantısı
├── configs/              # Konfigürasyon dosyaları
│   ├── switches/         # Switch konfigürasyonları
│   └── templates/        # YAML template'leri
├── scripts/              # Yardımcı scriptler
└── docs/                 # Dokümantasyon
```

## Geliştirme

### Test Çalıştırma
```bash
go test ./...
```

### Lint Kontrolü
```bash
golangci-lint run
```

### Mock Oluşturma
```bash
./generate-mock.sh
```

## Katkıda Bulunma

1. Fork yapın
2. Feature branch oluşturun (`git checkout -b feature/amazing-feature`)
3. Commit yapın (`git commit -m 'Add amazing feature'`)
4. Push yapın (`git push origin feature/amazing-feature`)
5. Pull Request oluşturun

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır.
