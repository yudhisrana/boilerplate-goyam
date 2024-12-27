# Boilerplate Go YAML MySQL

Boilerplate ini sengaja saya buat untuk membantu saya dalam memulai proyek Go dengan integrasi MySQL dan konfigurasi menggunakan YAML.

## Fitur:

-   **Go Modules**: Sudah dikonfigurasi untuk menggunakan Go modules dalam mengelola dependensi.
-   **Integrasi MySQL**: Sudah termasuk driver MySQL (`github.com/go-sql-driver/mysql`) untuk koneksi ke database.
-   **Konfigurasi YAML**: Menggunakan (`github.com/stretchr/testify/assert/yaml`) untuk memuat konfigurasi dari file YAML, memudahkan manajemen pengaturan aplikasi.
-   **Dukungan MySQL**: Sudah dikonfigurasi dengan database MySQL.
-   **Konfigurasi Mudah**: Menggunakan file YAML untuk pengaturan.
-   **Fitur Lain Next Update**

## Cara Menggunakan:

1. **Clone repository ini**:

    ```bash
    git clone https://github.com/yudhisrana/boilerplate-goyam.git nama-proyek-anda
    cd nama-proyek-anda
    ```

2. **Setup Boilerplate**

    ```bash
    chmod +x setup.sh
    ./setup.sh
    ```

3. **Rubah Konfigurasi Database**

    ```
    nama-proyek-anda/config.yaml
    ```

4. **Jalankan Aplikasi**
    ```bash
    go run .\cmd\main.go
    ```
