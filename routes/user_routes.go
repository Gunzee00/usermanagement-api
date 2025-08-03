package routes

import (
	"github.com/Gunzee00/usermanagement_api/controllers"
	"github.com/Gunzee00/usermanagement_api/middlewares"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Auth routes
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Protected routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.JWTMiddleware)

	api.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	api.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	api.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	api.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	return r
}
