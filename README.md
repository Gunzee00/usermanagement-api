#  User Management API

Mini project RESTful API menggunakan Golang 1.23, PostgreSQL, dan JWT Authentication, dengan fitur umum seperti register, login, dan CRUD user.

---

## Features

-  Register User
-  Login with JWT
-  Input Validation (Email & Password)
-  Password hashing dengan Bcrypt
-  Protected Routes menggunakan JWT Middleware
-  Dockerized dengan Alpine 3.20 + PostgreSQL
  
---

## Tech Stack

- Golang 1.23 (Alpine image)
- PostgreSQL (Docker latest)
- GORM ORM
- JWT Auth (`github.com/golang-jwt/jwt/v5`)
- Bcrypt untuk hashing password
- Validator (`github.com/go-playground/validator/v10`)

 ## Untuk Kolega & Klien
- Dokumentasi ini mencakup seluruh proses dari instalasi, setup, hingga penggunaan endpoint API.
- API dapat di-deploy di server manapun berkat dukungan Docker.
- Dapat dikembangkan lebih lanjut seperti reset password, otorisasi berbasis role, dll.
- Struktur modular memudahkan pengembangan dan debugging bersama tim.
 

 ## Instalasi dan Penggunaan

```bash
 1. Clone repository
git clone https://github.com/username/usermanagement-api.git
cd usermanagement-api

 2. Buat file .env
Buat file .env di root folder:
DB_USER=user
DB_PASSWORD=root
DB_NAME=usermanagement
DB_HOST=db
DB_PORT=5432
JWT_SECRET=supersecretkey


 3. Jalankan Docker
docker-compose up --build

---

 4. Port
API berjalan di: http://localhost:8080
Pastikan tidak ada konflik pada port 8080 dan 5432.

----------------------------------------------------------
 API Endpoint Guide

A. Register
Endpoint: POST /register

Body JSON:

{
  "name": "Gunawan",
  "email": "gunawan@mail.com",
  "password": "password123"
}

---

B. Login
Endpoint: POST /login

Body JSON:

{
  "email": "gunawan@mail.com",
  "password": "password123"
}

Response:

{
  "token": "eyJhbGciOiJIUzI1..."
}

Tambahkan token ke header:

Authorization: Bearer <token>

---

 C. Get All Users (Harus menggunakan Token)

Endpoint: GET /users

Header: 

Authorization: Bearer <token>

 D. Create Users (Harus menggunakan Token)

Endpoint: POST /api/users

Deskripsi: Menambahkan user baru

Header:

Authorization: Bearer <token>

Body JSON:

{
  "name": "Gunawan",
  "email": "gunawan@mail.com",
  "password": "password123"
}

 E. Get User by ID (Harus menggunakan Token)

Endpoint: GET /api/users/{id}

Deskripsi: Mendapatkan detail user berdasarkan ID

Header:
Authorization: Bearer <token>

 F. Update User by ID (Harus menggunakan Token)
Endpoint: PUT /api/users/{id}

Deskripsi: Mengupdate data user berdasarkan ID

Header:
Authorization: Bearer <token>

Body JSON:

{
  "name": "Gunawan Update",
  "email": "gunawanupdate@mail.com",
  "password": "newpassword123"
}

 G. Delete User by ID (Harus menggunakan Token)
Endpoint: DELETE /api/users/{id}

Deskripsi: Menghapus user berdasarkan ID

Header:

Authorization: Bearer <token>

---

## Docker Overview

Struktur docker-compose.yml

version: '3.8'
services:
  db:
    image: postgres:latest
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: root
      POSTGRES_DB: usermanagement
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data

  app:
    build: .
    container_name: usermanagement-api
    depends_on:
      - db
    ports:
      - "8080:8080"
    # Tambahkan environment from .env file jika diperlukan
    # env_file:
    #   - .env

volumes:
  pgdata:

 


