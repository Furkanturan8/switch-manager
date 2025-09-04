# Switch Manager API Dokümantasyonu

Bu dokümantasyon Switch Manager API'sinin tüm endpoint'lerini ve kullanım örneklerini içerir.

## Base URL
```
http://localhost:8080/api/v1
```

## Authentication
Şu anda authentication gerektirmiyor. Gelecekte JWT token kullanılabilir.

---

## Switch Endpoints

### 1. Switch Oluşturma
**POST** `/switches`

**Request Body:**
```json
{
  "name": "Core-Switch-01",
  "ip_address": "192.168.1.1",
  "hostname": "core-switch-01",
  "model": "Cisco Catalyst 2960",
  "location": "Data Center Rack 1",
  "description": "Ana core switch",
  "ssh_username": "admin",
  "ssh_password": "password123",
  "ssh_key_path": "/path/to/key",
  "ssh_port": 22
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "name": "Core-Switch-01",
  "ip_address": "192.168.1.1",
  "hostname": "core-switch-01",
  "model": "Cisco Catalyst 2960",
  "location": "Data Center Rack 1",
  "description": "Ana core switch",
  "ssh_username": "admin",
  "ssh_key_path": "/path/to/key",
  "ssh_port": 22,
  "status": "offline",
  "last_seen": "0001-01-01T00:00:00Z",
  "ios_version": "",
  "port_count": 0,
  "vlan_count": 0,
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:00:00Z"
}
```

### 2. Switch Detayını Getirme
**GET** `/switches/{id}`

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "Core-Switch-01",
  "ip_address": "192.168.1.1",
  "hostname": "core-switch-01",
  "model": "Cisco Catalyst 2960",
  "location": "Data Center Rack 1",
  "description": "Ana core switch",
  "ssh_username": "admin",
  "ssh_key_path": "/path/to/key",
  "ssh_port": 22,
  "status": "online",
  "last_seen": "2024-01-01T12:30:00Z",
  "ios_version": "15.2",
  "port_count": 24,
  "vlan_count": 5,
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:30:00Z"
}
```

### 3. Tüm Switch'leri Listeleme
**GET** `/switches`

**Response (200 OK):**
```json
{
  "switches": [
    {
      "id": 1,
      "name": "Core-Switch-01",
      "ip_address": "192.168.1.1",
      "hostname": "core-switch-01",
      "model": "Cisco Catalyst 2960",
      "location": "Data Center Rack 1",
      "status": "online",
      "last_seen": "2024-01-01T12:30:00Z",
      "port_count": 24,
      "vlan_count": 5,
      "created_at": "2024-01-01T12:00:00Z"
    }
  ],
  "count": 1
}
```

### 4. Switch Güncelleme
**PUT** `/switches/{id}`

**Request Body:**
```json
{
  "name": "Updated Switch Name",
  "location": "New Location",
  "description": "Updated description",
  "ssh_port": 2222
}
```

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "Updated Switch Name",
  "ip_address": "192.168.1.1",
  "hostname": "core-switch-01",
  "model": "Cisco Catalyst 2960",
  "location": "New Location",
  "description": "Updated description",
  "ssh_username": "admin",
  "ssh_key_path": "/path/to/key",
  "ssh_port": 2222,
  "status": "online",
  "last_seen": "2024-01-01T12:30:00Z",
  "ios_version": "15.2",
  "port_count": 24,
  "vlan_count": 5,
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:35:00Z"
}
```

### 5. Switch Durum Güncelleme
**PATCH** `/switches/{id}/status`

**Request Body:**
```json
{
  "status": "online"
}
```

**Response (200 OK):**
```json
{
  "message": "Switch status updated successfully",
  "status": "online"
}
```

### 6. Switch Silme
**DELETE** `/switches/{id}`

**Response (204 No Content)**

---

## Port Endpoints

### 1. Port Oluşturma
**POST** `/ports`

**Request Body:**
```json
{
  "switch_id": 1,
  "name": "GigabitEthernet0/1",
  "interface": "Gi0/1",
  "description": "Server Port",
  "speed": 1000,
  "mode": "access",
  "duplex": "full",
  "mtu": 1500,
  "access_vlan": 10,
  "trunk_vlans": [10, 20, 30],
  "poe": true,
  "max_mac": 5
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "switch_id": 1,
  "name": "GigabitEthernet0/1",
  "interface": "Gi0/1",
  "description": "Server Port",
  "status": "down",
  "admin_status": "down",
  "oper_status": "down",
  "speed": 1000,
  "mode": "access",
  "duplex": "full",
  "mtu": 1500,
  "mac_address": "",
  "access_vlan": 10,
  "trunk_vlans": [10, 20, 30],
  "poe": true,
  "max_mac": 5,
  "last_change": "0001-01-01T00:00:00Z",
  "error_count": 0,
  "switch_name": "Core-Switch-01",
  "switch_ip": "192.168.1.1",
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:00:00Z"
}
```

### 2. Port Detayını Getirme
**GET** `/ports/{id}`

**Response (200 OK):**
```json
{
  "id": 1,
  "switch_id": 1,
  "name": "GigabitEthernet0/1",
  "interface": "Gi0/1",
  "description": "Server Port",
  "status": "up",
  "admin_status": "up",
  "oper_status": "up",
  "speed": 1000,
  "mode": "access",
  "duplex": "full",
  "mtu": 1500,
  "mac_address": "00:11:22:33:44:55",
  "access_vlan": 10,
  "trunk_vlans": [10, 20, 30],
  "poe": true,
  "max_mac": 5,
  "last_change": "2024-01-01T12:30:00Z",
  "error_count": 0,
  "switch_name": "Core-Switch-01",
  "switch_ip": "192.168.1.1",
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:30:00Z"
}
```

### 3. Tüm Port'ları Listeleme
**GET** `/ports`

**Response (200 OK):**
```json
{
  "ports": [
    {
      "id": 1,
      "switch_id": 1,
      "name": "GigabitEthernet0/1",
      "interface": "Gi0/1",
      "status": "up",
      "admin_status": "up",
      "oper_status": "up",
      "speed": 1000,
      "mode": "access",
      "access_vlan": 10,
      "poe": true,
      "switch_name": "Core-Switch-01",
      "switch_ip": "192.168.1.1",
      "last_change": "2024-01-01T12:30:00Z",
      "created_at": "2024-01-01T12:00:00Z"
    }
  ],
  "count": 1
}
```

### 4. Port Güncelleme
**PUT** `/ports/{id}`

**Request Body:**
```json
{
  "name": "Updated Port Name",
  "description": "Updated description",
  "speed": 10000,
  "mode": "trunk",
  "access_vlan": 20
}
```

**Response (200 OK):**
```json
{
  "id": 1,
  "switch_id": 1,
  "name": "Updated Port Name",
  "interface": "Gi0/1",
  "description": "Updated description",
  "status": "up",
  "admin_status": "up",
  "oper_status": "up",
  "speed": 10000,
  "mode": "trunk",
  "duplex": "full",
  "mtu": 1500,
  "mac_address": "00:11:22:33:44:55",
  "access_vlan": 20,
  "trunk_vlans": [10, 20, 30],
  "poe": true,
  "max_mac": 5,
  "last_change": "2024-01-01T12:30:00Z",
  "error_count": 0,
  "switch_name": "Core-Switch-01",
  "switch_ip": "192.168.1.1",
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:35:00Z"
}
```

### 5. Port Durum Güncelleme
**PATCH** `/ports/{id}/status`

**Request Body:**
```json
{
  "status": "up",
  "admin_status": "up"
}
```

**Response (200 OK):**
```json
{
  "message": "Port status updated successfully",
  "status": "up",
  "admin_status": "up"
}
```

### 6. Port Silme
**DELETE** `/ports/{id}`

**Response (204 No Content)**

---

## VLAN Endpoints

### 1. VLAN Oluşturma
**POST** `/vlans`

**Request Body:**
```json
{
  "switch_id": 1,
  "vlan_id": 10,
  "name": "VLAN-10",
  "description": "Server VLAN",
  "mtu": 1500,
  "stp_enabled": true,
  "priority": 32768
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "switch_id": 1,
  "vlan_id": 10,
  "name": "VLAN-10",
  "description": "Server VLAN",
  "admin_status": "enabled",
  "oper_status": "down",
  "status": "active",
  "mtu": 1500,
  "stp_enabled": true,
  "priority": 32768,
  "port_count": 0,
  "switch_name": "Core-Switch-01",
  "switch_ip": "192.168.1.1",
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:00:00Z"
}
```

### 2. VLAN Detayını Getirme
**GET** `/vlans/{id}`

**Response (200 OK):**
```json
{
  "id": 1,
  "switch_id": 1,
  "vlan_id": 10,
  "name": "VLAN-10",
  "description": "Server VLAN",
  "admin_status": "enabled",
  "oper_status": "up",
  "status": "active",
  "mtu": 1500,
  "stp_enabled": true,
  "priority": 32768,
  "port_count": 5,
  "switch_name": "Core-Switch-01",
  "switch_ip": "192.168.1.1",
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:30:00Z"
}
```

### 3. Tüm VLAN'ları Listeleme
**GET** `/vlans`

**Response (200 OK):**
```json
{
  "vlans": [
    {
      "id": 1,
      "switch_id": 1,
      "vlan_id": 10,
      "name": "VLAN-10",
      "description": "Server VLAN",
      "admin_status": "enabled",
      "oper_status": "up",
      "status": "active",
      "port_count": 5,
      "switch_name": "Core-Switch-01",
      "switch_ip": "192.168.1.1",
      "created_at": "2024-01-01T12:00:00Z"
    }
  ],
  "count": 1
}
```

### 4. VLAN Güncelleme
**PUT** `/vlans/{id}`

**Request Body:**
```json
{
  "name": "Updated VLAN Name",
  "description": "Updated description",
  "mtu": 9000,
  "stp_enabled": false
}
```

**Response (200 OK):**
```json
{
  "id": 1,
  "switch_id": 1,
  "vlan_id": 10,
  "name": "Updated VLAN Name",
  "description": "Updated description",
  "admin_status": "enabled",
  "oper_status": "up",
  "status": "active",
  "mtu": 9000,
  "stp_enabled": false,
  "priority": 32768,
  "port_count": 5,
  "switch_name": "Core-Switch-01",
  "switch_ip": "192.168.1.1",
  "created_at": "2024-01-01T12:00:00Z",
  "updated_at": "2024-01-01T12:35:00Z"
}
```

### 5. VLAN Durum Güncelleme
**PATCH** `/vlans/{id}/status`

**Request Body:**
```json
{
  "admin_status": "disabled",
  "status": "suspended"
}
```

**Response (200 OK):**
```json
{
  "message": "VLAN status updated successfully",
  "admin_status": "disabled",
  "status": "suspended"
}
```

### 6. VLAN Silme
**DELETE** `/vlans/{id}`

**Response (204 No Content)**

---

## Error Responses

### 400 Bad Request
```json
{
  "error": "Invalid request",
  "message": "Validation failed"
}
```

### 404 Not Found
```json
{
  "error": "Not found",
  "message": "Switch not found"
}
```

### 500 Internal Server Error
```json
{
  "error": "Internal server error",
  "message": "Something went wrong"
}
```

---

## Postman Collection

### Environment Variables
```
base_url: http://localhost:8080/api/v1
```

### Headers
```
Content-Type: application/json
Accept: application/json
```

### Test Scripts (Her endpoint için)

**Status Code Test:**
```javascript
pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);
});
```

**Response Time Test:**
```javascript
pm.test("Response time is less than 1000ms", function () {
    pm.expect(pm.response.responseTime).to.be.below(1000);
});
```

**JSON Response Test:**
```javascript
pm.test("Response is JSON", function () {
    pm.response.to.be.json;
});
```

---

## Örnek Postman Workflow

1. **Switch Oluştur** → POST `/switches`
2. **Switch Detayını Getir** → GET `/switches/{id}`
3. **Port Oluştur** → POST `/ports` (switch_id ile)
4. **VLAN Oluştur** → POST `/vlans` (switch_id ile)
5. **Port Durumunu Güncelle** → PATCH `/ports/{id}/status`
6. **Switch Durumunu Güncelle** → PATCH `/switches/{id}/status`
7. **Tüm Verileri Listele** → GET `/switches`, GET `/ports`, GET `/vlans`

---

## Notlar

- Tüm tarih formatları ISO 8601 standardında (RFC 3339)
- Port modları: `access`, `trunk`, `routed`
- Duplex modları: `full`, `half`, `auto`
- Switch durumları: `online`, `offline`, `error`
- Port durumları: `up`, `down`, `disabled`
- VLAN durumları: `active`, `suspended`
- VLAN admin durumları: `enabled`, `disabled`
