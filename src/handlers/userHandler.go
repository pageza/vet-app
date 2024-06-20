package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/pageza/vet-app/src/models"
    "github.com/pageza/vet-app/src/db"
    "github.com/gorilla/mux"
    "strconv"
    "gorm.io/gorm"
)

// GetUsers handles GET requests to fetch all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
    var users []models.User
    result := db.DB.Find(&users)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

// GetUser handles GET requests to fetch a user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var user models.User
    result := db.DB.First(&user, id)
    if result.Error != nil {
        if gorm.IsRecordNotFoundError(result.Error) {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        }
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

// CreateUser handles POST requests to create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result := db.DB.Create(&user)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

// UpdateUser handles PUT requests to update a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var user models.User
    if err := db.DB.First(&user, id).Error; err != nil {
        if gorm.IsRecordNotFoundError(err) {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    db.DB.Model(&user).Updates(user)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}

// DeleteUser handles DELETE requests to delete a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var user models.User
    if err := db.DB.First(&user, id).Error; err != nil {
        if gorm.IsRecordNotFoundError(err) {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    db.DB.Delete(&user)

    w.WriteHeader(http.StatusNoContent)
}
