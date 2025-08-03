package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/Gunzee00/usermanagement_api/config"
	"github.com/Gunzee00/usermanagement_api/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// validasi email
func isValidEmail(email string) bool {
	
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// CreateUser
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, `{"message": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	// Validasi field kosong
	if user.Name == "" || user.Email == "" || user.Password == "" {
		http.Error(w, `{"message": "All fields are required"}`, http.StatusBadRequest)
		return
	}

	// Validasi format email
	if !isValidEmail(user.Email) {
		http.Error(w, `{"message": "Invalid email format"}`, http.StatusBadRequest)
		return
	}

	// Validasi password minimal 6 karakter
	if len(user.Password) < 6 {
		http.Error(w, `{"message": "Password must be at least 6 characters"}`, http.StatusBadRequest)
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		http.Error(w, `{"message": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUsers
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		http.Error(w, `{"message": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// GetUserByID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"message": "Invalid ID"}`, http.StatusBadRequest)
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, `{"message": "User not found"}`, http.StatusNotFound)
		} else {
			http.Error(w, `{"message": "`+err.Error()+`"}`, http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(user)
}

// UpdateUser
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"message": "Invalid ID"}`, http.StatusBadRequest)
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		http.Error(w, `{"message": "User not found"}`, http.StatusNotFound)
		return
	}

	var updated models.User
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, `{"message": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	// Validasi
	if updated.Name == "" || updated.Email == "" || updated.Password == "" {
		http.Error(w, `{"message": "All fields are required"}`, http.StatusBadRequest)
		return
	}

	if !isValidEmail(updated.Email) {
		http.Error(w, `{"message": "Invalid email format"}`, http.StatusBadRequest)
		return
	}

	if len(updated.Password) < 6 {
		http.Error(w, `{"message": "Password must be at least 6 characters"}`, http.StatusBadRequest)
		return
	}

	user.Name = updated.Name
	user.Email = updated.Email
	user.Password = updated.Password

	if err := config.DB.Save(&user).Error; err != nil {
		http.Error(w, `{"message": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// DeleteUser
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"message": "Invalid ID"}`, http.StatusBadRequest)
		return
	}

	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		http.Error(w, `{"message": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
