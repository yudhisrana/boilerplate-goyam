# Boilerplate Go YAML MySQL

Terima kasih telah menggunakan **Boilerplate Go YAML MySQL** ini! 🎉

Boilerplate ini dirancang untuk membantu Anda memulai proyek Go dengan integrasi MySQL dan konfigurasi menggunakan YAML. Template ini menyederhanakan pengaturan awal proyek, memungkinkan Anda untuk fokus pada pengembangan fitur utama aplikasi Anda.

## Fitur:

-   **Go Modules**: Sudah dikonfigurasi untuk menggunakan Go modules dalam mengelola dependensi.
-   **Integrasi MySQL**: Sudah termasuk driver MySQL (`github.com/go-sql-driver/mysql`) untuk koneksi ke database.
-   **Konfigurasi YAML**: Menggunakan (`github.com/stretchr/testify/assert/yaml`) untuk memuat konfigurasi dari file YAML, memudahkan manajemen pengaturan aplikasi.
-   **Struktur Folder Sederhana**: Struktur yang bersih dan sederhana agar Anda bisa lebih fokus pada pengembangan aplikasi, bukan pengaturan boilerplate.

## Kenapa Menggunakan Template Ini?

-   **Membantu Membuat Awal Struktur Project**: Mulai proyek Anda dengan struktur dan dependensi yang sudah siap.
-   **Dukungan MySQL**: Sudah dikonfigurasi untuk bekerja dengan MySQL, membuatnya lebih mudah untuk menghubungkan dan melakukan query ke database.
-   **Konfigurasi Mudah**: Menggunakan file YAML untuk pengaturan, yang mudah dibaca dan dipelihara.

## Cara Menggunakan:

1. **Clone repository ini**:

    ```bash
    git clone https://github.com/yourusername/boilerplate-goyam.git nama-proyek-anda
    cd nama-proyek-anda
    ```

2. **Inisialisasi Go module**

    ```bash
    go mod init github.com/yourusername/nama-proyek-anda
    go mod tidy
    ```

3. **Rubah Konfigurasi Database**
   Buka dan sesuaikan konfigurasi database pada file:

    ```
    your-project/config.yaml
    ```

4. **Jalankan Aplikasi**
   Jalankan aplikasi dengan perintah berikut:
    ```bash
    go run .\cmd\main.go
    ```
