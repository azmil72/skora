# Backend API - Ujian Kompetensi

Backend API untuk sistem ujian kompetensi dengan fitur CRUD lengkap.

## Setup

### Prerequisites
- Go 1.21+
- Docker & Docker Compose
- PostgreSQL (via Docker)

### Installation

1. Clone repository dan masuk ke direktori backend
```bash
cd backend
```

2. Install dependencies
```bash
go mod tidy
```

3. Setup database dengan Docker
```bash
docker-compose up -d
```

4. Copy environment file
```bash
cp .env.example .env
```

5. Jalankan aplikasi
```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## Database Configuration

Database menggunakan PostgreSQL dengan konfigurasi:
- Host: localhost
- Port: 5432
- Database: mydb
- Username: postgres
- Password: password

## API Endpoints

### Users
- `POST /api/v1/users` - Create user
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### Rooms
- `POST /api/v1/rooms` - Create room
- `GET /api/v1/rooms` - Get all rooms
- `GET /api/v1/rooms/:id` - Get room by ID (UUID)
- `PUT /api/v1/rooms/:id` - Update room
- `DELETE /api/v1/rooms/:id` - Delete room

### Pertanyaans
- `POST /api/v1/pertanyaans` - Create pertanyaan
- `GET /api/v1/pertanyaans` - Get all pertanyaans
- `GET /api/v1/pertanyaans/:id` - Get pertanyaan by ID
- `PUT /api/v1/pertanyaans/:id` - Update pertanyaan
- `DELETE /api/v1/pertanyaans/:id` - Delete pertanyaan

### Sesi Ujians
- `POST /api/v1/sesi-ujians` - Create sesi ujian
- `GET /api/v1/sesi-ujians` - Get all sesi ujians
- `GET /api/v1/sesi-ujians/:id` - Get sesi ujian by ID
- `PUT /api/v1/sesi-ujians/:id` - Update sesi ujian
- `DELETE /api/v1/sesi-ujians/:id` - Delete sesi ujian

### Answers
- `POST /api/v1/answers` - Create answer
- `GET /api/v1/answers` - Get all answers
- `GET /api/v1/answers/:id` - Get answer by ID
- `PUT /api/v1/answers/:id` - Update answer
- `DELETE /api/v1/answers/:id` - Delete answer

### Hasil Ujians
- `POST /api/v1/hasil-ujians` - Create hasil ujian
- `GET /api/v1/hasil-ujians` - Get all hasil ujians
- `GET /api/v1/hasil-ujians/:id` - Get hasil ujian by ID
- `PUT /api/v1/hasil-ujians/:id` - Update hasil ujian
- `DELETE /api/v1/hasil-ujians/:id` - Delete hasil ujian

## Models

### User
```json
{
  "id_users": 1,
  "nama": "John Doe",
  "email": "john@example.com",
  "password_hash": "hashed_password",
  "created_at": "2024-01-01T00:00:00Z"
}
```

### Room
```json
{
  "id_room": "uuid-v7",
  "room_name": "Room Test",
  "durasi": 60,
  "created_by": 1,
  "created_at": "2024-01-01T00:00:00Z"
}
```

### Pertanyaan
```json
{
  "id": 1,
  "room_id": "uuid-v7",
  "pertanyaan_text": "Apa itu Go?",
  "type_pertanyaan": "multiple_choice"
}
```

## Features

- ✅ CRUD operations untuk semua model
- ✅ UUID v7 untuk Room ID
- ✅ PostgreSQL database
- ✅ GORM ORM
- ✅ Gin web framework
- ✅ Environment configuration
- ✅ Database auto-migration
- ✅ Foreign key relationships
- ✅ JSON API responses

## Environment Variables

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=mydb
PORT=8080
```