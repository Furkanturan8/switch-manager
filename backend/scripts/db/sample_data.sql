-- Switch Manager - Örnek Veri Ekleme Scripti
-- PostgreSQL için hazırlanmıştır

-- Önce mevcut verileri temizle (isteğe bağlı)
-- DELETE FROM ports;
-- DELETE FROM switches;
-- DELETE FROM vlans;

-- Switches tablosuna örnek veriler



INSERT INTO switches (id, name, ip_address, hostname, model, location, description, status, created_at, updated_at) VALUES
(1, 'Core-SW-01', '192.168.1.1', 'core-sw-01', 'Cisco Catalyst 9300', 'Ana Veri Merkezi', 'Ana core switch - 1. kat', 'online', NOW(), NOW()),
(2, 'Core-SW-02', '192.168.1.2', 'core-sw-02', 'Cisco Catalyst 9300', 'Ana Veri Merkezi', 'Ana core switch - 2. kat', 'online', NOW(), NOW()),
(3, 'Access-SW-01', '192.168.2.1', 'access-sw-01', 'Cisco Catalyst 2960', '1. Kat', '1. kat access switch', 'online', NOW(), NOW()),
(4, 'Access-SW-02', '192.168.2.2', 'access-sw-02', 'Cisco Catalyst 2960', '2. Kat', '2. kat access switch', 'online', NOW(), NOW()),
(5, 'Edge-SW-01', '192.168.3.1', 'edge-sw-01', 'HP ProCurve 5406', 'Giriş Katı', 'Giriş katı edge switch', 'online', NOW(), NOW());

-- VLANS tablosuna örnek veriler
INSERT INTO vlans (id, switch_id, vlan_id, name, description, admin_status, oper_status, status, mtu, stp_enabled, priority, created_at, updated_at) VALUES
(1, 1, 1, 'Management', 'Switch yönetim ağı', 'enabled', 'up', 'active', 1500, true, 32768, NOW(), NOW()),
(2, 1, 10, 'Data', 'Veri ağı', 'enabled', 'up', 'active', 1500, true, 32768, NOW(), NOW()),
(3, 1, 20, 'Voice', 'Ses ağı', 'enabled', 'up', 'active', 1500, true, 32768, NOW(), NOW()),
(4, 1, 30, 'Guest', 'Misafir ağı', 'enabled', 'up', 'active', 1500, true, 32768, NOW(), NOW()),
(5, 2, 1, 'Management', 'Switch yönetim ağı', 'enabled', 'up', 'active', 1500, true, 32768, NOW(), NOW()),
(6, 2, 10, 'Data', 'Veri ağı', 'enabled', 'up', 'active', 1500, true, 32768, NOW(), NOW()),
(7, 2, 20, 'Voice', 'Ses ağı', 'enabled', 'up', 'active', 1500, true, 32768, NOW(), NOW()),
(8, 3, 1, 'Management', 'Switch yönetim ağı', 'enabled', 'up', 'active', 1500, true, 32768, NOW(), NOW()),
(9, 3, 10, 'Data', 'Veri ağı', 'enabled', 'up', 'active', 1500, true, 32768, NOW(), NOW()),
(10, 4, 1, 'Management', 'Switch yönetim ağı', 'enabled', 'up', 'active', 1500, true, 32768, NOW(), NOW()),
(11, 4, 10, 'Data', 'Veri ağı', 'enabled', 'up', 'active', 1500, true, 32768, NOW(), NOW()),
(12, 5, 1, 'Management', 'Switch yönetim ağı', 'enabled', 'up', 'active', 1500, true, 32768, NOW(), NOW()),
(13, 5, 30, 'Guest', 'Misafir ağı', 'enabled', 'up', 'active', 1500, true, 32768, NOW(), NOW());

-- Ports tablosuna örnek veriler
-- Core-SW-01 için portlar
INSERT INTO ports (id, switch_id, name, interface, status, admin_status, oper_status, speed, mode, duplex, mtu, mac_address, description, access_vlan, trunk_vlans, poe, max_mac, error_count, last_change, created_at, updated_at) VALUES
(1, 1, 'GigabitEthernet0/1', 'Gi0/1', 'up', 'up', 'up', 1000, 'trunk', 'full', 1500, '00:1B:63:84:45:E6', 'Core uplink to Core-SW-02', 0, '[1,10,20,30]', false, 0, 0, NOW(), NOW(), NOW()),
(2, 1, 'GigabitEthernet0/2', 'Gi0/2', 'up', 'up', 'up', 1000, 'access', 'full', 1500, '00:1B:63:84:45:E7', 'Server 1 connection', 10, '[]', false, 1, 0, NOW(), NOW(), NOW()),
(3, 1, 'GigabitEthernet0/3', 'Gi0/3', 'up', 'up', 'up', 1000, 'access', 'full', 1500, '00:1B:63:84:45:E8', 'Server 2 connection', 10, '[]', false, 1, 0, NOW(), NOW(), NOW()),
(4, 1, 'GigabitEthernet0/4', 'Gi0/4', 'down', 'up', 'down', 1000, 'access', 'full', 1500, '00:1B:63:84:45:E9', 'Unused port', 10, '[]', false, 0, 0, NOW(), NOW(), NOW()),
(5, 1, 'GigabitEthernet0/5', 'Gi0/5', 'up', 'up', 'up', 1000, 'trunk', 'full', 1500, '00:1B:63:84:45:EA', 'Uplink to Access-SW-01', 0, '[1,10,20]', false, 0, 0, NOW(), NOW(), NOW()),
(6, 1, 'GigabitEthernet0/6', 'Gi0/6', 'up', 'up', 'up', 1000, 'trunk', 'full', 1500, '00:1B:63:84:45:EB', 'Uplink to Access-SW-02', 0, '[1,10,20]', false, 0, 0, NOW(), NOW(), NOW());

-- Core-SW-02 için portlar
INSERT INTO ports (id, switch_id, name, interface, status, admin_status, oper_status, speed, mode, duplex, mtu, mac_address, description, access_vlan, trunk_vlans, poe, max_mac, error_count, last_change, created_at, updated_at) VALUES
(7, 2, 'GigabitEthernet0/1', 'Gi0/1', 'up', 'up', 'up', 1000, 'trunk', 'full', 1500, '00:1B:63:84:45:EC', 'Core uplink to Core-SW-01', 0, '[1,10,20,30]', false, 0, 0, NOW(), NOW(), NOW()),
(8, 2, 'GigabitEthernet0/2', 'Gi0/2', 'up', 'up', 'up', 1000, 'access', 'full', 1500, '00:1B:63:84:45:ED', 'Backup Server 1', 10, '[]', false, 1, 0, NOW(), NOW(), NOW()),
(9, 2, 'GigabitEthernet0/3', 'Gi0/3', 'up', 'up', 'up', 1000, 'access', 'full', 1500, '00:1B:63:84:45:EE', 'Backup Server 2', 10, '[]', false, 1, 0, NOW(), NOW(), NOW()),
(10, 2, 'GigabitEthernet0/4', 'Gi0/4', 'up', 'up', 'up', 1000, 'access', 'full', 1500, '00:1B:63:84:45:EF', 'Voice Gateway', 20, '[]', false, 1, 0, NOW(), NOW(), NOW());

-- Access-SW-01 için portlar
INSERT INTO ports (id, switch_id, name, interface, status, admin_status, oper_status, speed, mode, duplex, mtu, mac_address, description, access_vlan, trunk_vlans, poe, max_mac, error_count, last_change, created_at, updated_at) VALUES
(11, 3, 'FastEthernet0/1', 'Fa0/1', 'up', 'up', 'up', 100, 'access', 'full', 1500, '00:1B:63:84:45:F0', 'PC 1', 10, '[]', false, 1, 0, NOW(), NOW(), NOW()),
(12, 3, 'FastEthernet0/2', 'Fa0/2', 'up', 'up', 'up', 100, 'access', 'full', 1500, '00:1B:63:84:45:F1', 'PC 2', 10, '[]', false, 1, 0, NOW(), NOW(), NOW()),
(13, 3, 'FastEthernet0/3', 'Fa0/3', 'up', 'up', 'up', 100, 'access', 'full', 1500, '00:1B:63:84:45:F2', 'IP Phone 1', 20, '[]', true, 1, 0, NOW(), NOW(), NOW()),
(14, 3, 'FastEthernet0/4', 'Fa0/4', 'up', 'up', 'up', 100, 'access', 'full', 1500, '00:1B:63:84:45:F3', 'IP Phone 2', 20, '[]', true, 1, 0, NOW(), NOW(), NOW()),
(15, 3, 'GigabitEthernet0/1', 'Gi0/1', 'up', 'up', 'up', 1000, 'trunk', 'full', 1500, '00:1B:63:84:45:F4', 'Uplink to Core-SW-01', 0, '[1,10,20]', false, 0, 0, NOW(), NOW(), NOW());

-- Access-SW-02 için portlar
INSERT INTO ports (id, switch_id, name, interface, status, admin_status, oper_status, speed, mode, duplex, mtu, mac_address, description, access_vlan, trunk_vlans, poe, max_mac, error_count, last_change, created_at, updated_at) VALUES
(16, 4, 'FastEthernet0/1', 'Fa0/1', 'up', 'up', 'up', 100, 'access', 'full', 1500, '00:1B:63:84:45:F5', 'PC 3', 10, '[]', false, 1, 0, NOW(), NOW(), NOW()),
(17, 4, 'FastEthernet0/2', 'Fa0/2', 'up', 'up', 'up', 100, 'access', 'full', 1500, '00:1B:63:84:45:F6', 'PC 4', 10, '[]', false, 1, 0, NOW(), NOW(), NOW()),
(18, 4, 'FastEthernet0/3', 'Fa0/3', 'up', 'up', 'up', 100, 'access', 'full', 1500, '00:1B:63:84:45:F7', 'IP Phone 3', 20, '[]', true, 1, 0, NOW(), NOW(), NOW()),
(19, 4, 'GigabitEthernet0/1', 'Gi0/1', 'up', 'up', 'up', 1000, 'trunk', 'full', 1500, '00:1B:63:84:45:F8', 'Uplink to Core-SW-01', 0, '[1,10,20]', false, 0, 0, NOW(), NOW(), NOW());

-- Edge-SW-01 için portlar
INSERT INTO ports (id, switch_id, name, interface, status, admin_status, oper_status, speed, mode, duplex, mtu, mac_address, description, access_vlan, trunk_vlans, poe, max_mac, error_count, last_change, created_at, updated_at) VALUES
(20, 5, 'FastEthernet0/1', 'Fa0/1', 'up', 'up', 'up', 100, 'access', 'full', 1500, '00:1B:63:84:45:F9', 'Guest WiFi AP', 30, '[]', false, 1, 0, NOW(), NOW(), NOW()),
(21, 5, 'FastEthernet0/2', 'Fa0/2', 'up', 'up', 'up', 100, 'access', 'full', 1500, '00:1B:63:84:45:FA', 'Guest WiFi AP 2', 30, '[]', false, 1, 0, NOW(), NOW(), NOW()),
(22, 5, 'FastEthernet0/3', 'Fa0/3', 'down', 'down', 'down', 100, 'access', 'full', 1500, '00:1B:63:84:45:FB', 'Unused port', 30, '[]', false, 0, 0, NOW(), NOW(), NOW()),
(23, 5, 'GigabitEthernet0/1', 'Gi0/1', 'up', 'up', 'up', 1000, 'trunk', 'full', 1500, '00:1B:63:84:45:FC', 'Uplink to Core-SW-01', 0, '[1,30]', false, 0, 0, NOW(), NOW(), NOW());

-- Sequence'leri güncelle (eğer auto-increment kullanıyorsanız)
-- SELECT setval('switches_id_seq', (SELECT MAX(id) FROM switches));
-- SELECT setval('vlans_id_seq', (SELECT MAX(id) FROM vlans));
-- SELECT setval('ports_id_seq', (SELECT MAX(id) FROM ports));

-- Veri ekleme işlemi tamamlandı
SELECT 'Sample data inserted successfully!' as message;
SELECT 'Switches: ' || COUNT(*) as switches_count FROM switches;
SELECT 'VLANs: ' || COUNT(*) as vlans_count FROM vlans;
SELECT 'Ports: ' || COUNT(*) as ports_count FROM ports;
