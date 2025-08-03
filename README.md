# 📦 User Management API

Mini project RESTful API menggunakan Golang 1.23, PostgreSQL, dan JWT Authentication, dengan fitur umum seperti register, login, dan CRUD user.

---

## 🧩 Features

- ✅ Register User
- ✅ Login with JWT
- ✅ Input Validation (Email & Password)
- ✅ Password hashing dengan Bcrypt
- ✅ Protected Routes menggunakan JWT Middleware
- ✅ Dockerized dengan Alpine 3.20 + PostgreSQL
- ✅ Clean Architecture

---

## 🛠 Tech Stack

- Golang 1.23 (Alpine image)
- PostgreSQL (Docker latest)
- GORM ORM
- JWT Auth (`github.com/golang-jwt/jwt/v5`)
- Bcrypt untuk hashing password
- Validator (`github.com/go-playground/validator/v10`)

---

## 📥 Instalasi

 1. Clone repository

```bash
git clone https://github.com/username/usermanagement-api.git
cd usermanagement-api

 2. Clone repository
Buat file .env di root folder:


env
Salin
Edit
DB_USER=user
DB_PASSWORD=root
DB_NAME=usermanagement
DB_HOST=db
DB_PORT=5432
JWT_SECRET=rahasia_kamu

3. Jalankan Docker
bash
Salin
Edit
docker-compose up --build