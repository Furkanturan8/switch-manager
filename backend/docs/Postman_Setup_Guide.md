# Postman Kurulum ve Kullanım Kılavuzu

Bu kılavuz Switch Manager API'sini Postman ile test etmek için gerekli adımları içerir.

## 📋 İçindekiler

1. [Postman Kurulumu](#postman-kurulumu)
2. [Collection Import](#collection-import)
3. [Environment Setup](#environment-setup)
4. [Test Senaryoları](#test-senaryoları)
5. [Troubleshooting](#troubleshooting)

---

## 🚀 Postman Kurulumu

### 1. Postman İndirme
- [Postman Desktop App](https://www.postman.com/downloads/) indirin
- Veya [Postman Web](https://web.postman.com/) kullanın

### 2. Hesap Oluşturma (Opsiyonel)
- Postman hesabı oluşturarak collection'larınızı senkronize edebilirsiniz

---

## 📥 Collection Import

### 1. Collection Dosyasını İçe Aktarma
1. Postman'i açın
2. **Import** butonuna tıklayın
3. `Switch_Manager_Postman_Collection.json` dosyasını seçin
4. **Import** butonuna tıklayın

### 2. Collection Yapısı
```
Switch Manager API
├── Switches
│   ├── Create Switch
│   ├── Get Switch by ID
│   ├── Get All Switches
│   ├── Update Switch
│   ├── Update Switch Status
│   └── Delete Switch
├── Ports
│   ├── Create Port
│   ├── Get Port by ID
│   ├── Get All Ports
│   ├── Update Port
│   ├── Update Port Status
│   └── Delete Port
└── VLANs
    ├── Create VLAN
    ├── Get VLAN by ID
    ├── Get All VLANs
    ├── Update VLAN
    ├── Update VLAN Status
    └── Delete VLAN
```

---

## ⚙️ Environment Setup

### 1. Environment Dosyasını İçe Aktarma
1. Postman'de **Environments** sekmesine gidin
2. **Import** butonuna tıklayın
3. `Switch_Manager_Environment.json` dosyasını seçin
4. **Import** butonuna tıklayın

### 2. Environment Değişkenleri
```json
{
  "base_url": "http://localhost:8080/api/v1",
  "switch_id": "",
  "port_id": "",
  "vlan_id": ""
}
```

### 3. Environment'ı Aktif Etme
1. Sağ üst köşedeki environment dropdown'ından
2. **Switch Manager Environment**'ı seçin

---

## 🧪 Test Senaryoları

### Senaryo 1: Temel CRUD İşlemleri

#### 1.1 Switch Oluşturma ve Test Etme
1. **Create Switch** request'ini çalıştırın
2. Response'da `id` alanının geldiğini kontrol edin
3. Environment'da `switch_id` değişkeninin otomatik set edildiğini görün

#### 1.2 Switch Detayını Getirme
1. **Get Switch by ID** request'ini çalıştırın
2. Oluşturulan switch'in detaylarını görün
3. Status code 200 olduğunu kontrol edin

#### 1.3 Switch Güncelleme
1. **Update Switch** request'ini çalıştırın
2. Güncellenen alanları response'da kontrol edin
3. `updated_at` alanının değiştiğini görün

### Senaryo 2: Port İşlemleri

#### 2.1 Port Oluşturma
1. **Create Port** request'ini çalıştırın
2. `switch_id` environment değişkeninin kullanıldığını görün
3. Port'un başarıyla oluşturulduğunu kontrol edin

#### 2.2 Port Durum Güncelleme
1. **Update Port Status** request'ini çalıştırın
2. Status'un güncellendiğini kontrol edin

### Senaryo 3: VLAN İşlemleri

#### 3.1 VLAN Oluşturma
1. **Create VLAN** request'ini çalıştırın
2. VLAN'ın switch'e bağlandığını kontrol edin

#### 3.2 VLAN Durum Güncelleme
1. **Update VLAN Status** request'ini çalıştırın
2. Admin status ve status'un güncellendiğini kontrol edin

### Senaryo 4: Liste İşlemleri

#### 4.1 Tüm Verileri Listeleme
1. **Get All Switches** request'ini çalıştırın
2. **Get All Ports** request'ini çalıştırın
3. **Get All VLANs** request'ini çalıştırın
4. Her birinde `count` alanının doğru olduğunu kontrol edin

---

## 🔧 Troubleshooting

### Yaygın Hatalar ve Çözümleri

#### 1. Connection Refused
```
Error: connect ECONNREFUSED 127.0.0.1:8080
```
**Çözüm:**
- Backend server'ın çalıştığından emin olun
- Port 8080'in kullanılabilir olduğunu kontrol edin
- `base_url` environment değişkenini kontrol edin

#### 2. 404 Not Found
```
Error: 404 Not Found
```
**Çözüm:**
- API endpoint'lerinin doğru olduğunu kontrol edin
- Route'ların tanımlandığından emin olun
- Base URL'in doğru olduğunu kontrol edin

#### 3. 500 Internal Server Error
```
Error: 500 Internal Server Error
```
**Çözüm:**
- Backend log'larını kontrol edin
- Database bağlantısının çalıştığından emin olun
- Request body'nin doğru format'ta olduğunu kontrol edin

#### 4. Environment Variables Not Working
**Çözüm:**
- Environment'ın aktif olduğundan emin olun
- Değişken isimlerinin doğru olduğunu kontrol edin
- `{{variable_name}}` syntax'ını kullandığınızdan emin olun

---

## 📊 Test Sonuçları

### Test Scripts
Her request'te otomatik olarak çalışan test'ler:

1. **Status Code Test**: Response status code'un doğru olduğunu kontrol eder
2. **JSON Response Test**: Response'un JSON formatında olduğunu kontrol eder
3. **Required Fields Test**: Gerekli alanların response'da olduğunu kontrol eder
4. **Environment Variable Test**: ID'lerin environment'a kaydedildiğini kontrol eder

### Test Results
- ✅ **Passed**: Test başarılı
- ❌ **Failed**: Test başarısız
- ⚠️ **Skipped**: Test atlandı

---

## 🚀 Gelişmiş Kullanım

### 1. Collection Runner
1. Collection'ı seçin
2. **Run** butonuna tıklayın
3. Tüm request'leri sırayla çalıştırın
4. Test sonuçlarını görün

### 2. Newman (CLI)
```bash
# Newman kurulumu
npm install -g newman

# Collection'ı çalıştırma
newman run Switch_Manager_Postman_Collection.json -e Switch_Manager_Environment.json
```

### 3. CI/CD Integration
```yaml
# GitHub Actions örneği
- name: Run API Tests
  run: |
    newman run Switch_Manager_Postman_Collection.json \
      -e Switch_Manager_Environment.json \
      --reporters cli,json \
      --reporter-json-export results.json
```

---

## 📝 Notlar

- Environment değişkenleri otomatik olarak güncellenir
- Test script'leri her request'te çalışır
- Collection'ı team ile paylaşabilirsiniz
- Postman'de collection'ı fork edebilirsiniz

---

## 🔗 Faydalı Linkler

- [Postman Documentation](https://learning.postman.com/)
- [Newman Documentation](https://github.com/postmanlabs/newman)
- [Postman API](https://www.postman.com/postman/workspace/postman-public-workspace/collection/12959542-c8142d51-e97c-46b6-bd77-52bb66712c9a)
