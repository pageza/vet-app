package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/pageza/vet-app/src/db"
    "github.com/pageza/vet-app/src/models"
    "gorm.io/gorm"
)

// GetCalls handles GET requests to fetch all calls
func GetCalls(w http.ResponseWriter, r *http.Request) {
    var calls []models.Call
    result := db.DB.Find(&calls)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(calls)
}

// GetCall handles GET requests to fetch a call by ID
func GetCall(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid call ID", http.StatusBadRequest)
        return
    }

    var call models.Call
    result := db.DB.First(&call, id)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            http.Error(w, "Call not found", http.StatusNotFound)
        } else {
            http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        }
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(call)
}

// CreateCall handles POST requests to create a new call
func CreateCall(w http.ResponseWriter, r *http.Request) {
    var call models.Call
    if err := json.NewDecoder(r.Body).Decode(&call); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result := db.DB.Create(&call)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(call)
}

// UpdateCall handles PUT requests to update a call
func UpdateCall(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid call ID", http.StatusBadRequest)
        return
    }

    var call models.Call
    if err := db.DB.First(&call, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            http.Error(w, "Call not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    if err := json.NewDecoder(r.Body).Decode(&call); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    db.DB.Model(&call).Updates(call)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(call)
}

// DeleteCall handles DELETE requests to delete a call
func DeleteCall(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid call ID", http.StatusBadRequest)
        return
    }

    var call models.Call
    if err := db.DB.First(&call, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            http.Error(w, "Call not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    db.DB.Delete(&call)

    w.WriteHeader(http.StatusNoContent)
}
