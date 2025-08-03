package main

import (
	"log"
	"net/http"

	"github.com/Gunzee00/usermanagement_api/config"
	"github.com/Gunzee00/usermanagement_api/models"
	"github.com/Gunzee00/usermanagement_api/routes"
)

func main() {
	// Koneksi ke database dan simpan di variabel global config.DB
	config.ConnectDB()

	// Auto migrate model ke database menggunakan koneksi global
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Gagal migrasi: %v", err)
	}

	// Routing
	r := routes.SetupRoutes()

	log.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
