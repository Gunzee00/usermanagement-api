package main

import (
	"log"
	"net/http"

	"github.com/Gunzee00/usermanagement_api/config"
	"github.com/Gunzee00/usermanagement_api/models"
	"github.com/Gunzee00/usermanagement_api/routes"
)

func main() {
	
	config.ConnectDB()
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Gagal migrasi: %v", err)
	}

	r := routes.SetupRoutes()

	log.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
