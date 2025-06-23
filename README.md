# ewallet-ums

Sistem e-wallet User Management Service (UMS) berbasis Go. Proyek ini menyediakan layanan manajemen user, autentikasi, dan validasi token untuk aplikasi e-wallet, serta integrasi dengan layanan wallet eksternal.

## Fitur

- Registrasi user
- Login user
- Logout user
- Refresh token
- Validasi token (gRPC)
- Healthcheck endpoint

## Struktur Direktori

- `cmd/` : Entry point aplikasi (HTTP & gRPC server)
- `internal/api/` : Handler API (HTTP & gRPC)
- `internal/services/` : Bisnis logic utama
- `internal/models/` : Model data (User, Session)
- `internal/repository/` : Akses database
- `helpers/` : Helper (config, database, JWT, logger, response)
- `external/` : Integrasi ke layanan eksternal (wallet)

## Instalasi

1. **Clone repository**
   ```bash
   git clone <repo-url>
   cd ewallet-ums
   ```
2. **Install dependency**
   ```bash
   go mod download
   ```
3. **Buat file .env**
   Contoh isi file `.env`:
   ```env
   PORT=8083
   GRPC_PORT=7001
   DB_USER=root
   DB_PASS=yourpassword
   DB_HOST=127.0.0.1
   DB_NAME=ewallet_ums
   APP_SECRET=your_secret_key
   APP_NAME=ewallet-ums
   WALLET_HOST=http://127.0.0.1:8082
   WALLET_ENDPOINT_CREATE=/wallet/v1
   ```

## Menjalankan Aplikasi

### HTTP Server

```bash
go run main.go
```

Akan berjalan di port sesuai variabel `PORT` (default: 8083).

### gRPC Server

Uncomment baris `go cmd.ServeGRPC()` di `main.go` jika ingin menjalankan gRPC server.

## Endpoint Utama

- `GET /health` : Cek status aplikasi
- `POST /user/v1/register` : Registrasi user
- `POST /user/v1/login` : Login user
- `DELETE /user/v1/logout` : Logout user (butuh Authorization header)
- `PUT /user/v1/refresh-token` : Refresh token (butuh Authorization header)

## Lisensi

MIT
