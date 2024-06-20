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

// CreateResponse creates a new response
func CreateResponse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	callID, err := strconv.Atoi(vars["call_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var response models.Response
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response.CallID = uint(callID)
	if err := db.DB.Create(&response).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(response)
}

// GetResponses gets all responses for a call
func GetResponses(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	callID, err := strconv.Atoi(vars["call_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var responses []models.Response
	if err := db.DB.Where("call_id = ?", callID).Preload("User").Find(&responses).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(responses)
}

// GetResponsesForUser gets all responses for a user
func GetResponsesForUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var responses []models.Response
	if err := db.DB.Where("user_id = ?", userID).Preload("Call").Find(&responses).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(responses)
}

// GetResponse gets a single response by ID
func GetResponse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var response models.Response
	if err := db.DB.Preload("User").Preload("Call").First(&response, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(response)
}

// UpdateResponse updates a response by ID
func UpdateResponse(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid response ID", http.StatusBadRequest)
        return
    }

    var existingResponse models.Response
    if err := db.DB.First(&existingResponse, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            http.Error(w, "Response not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    var updatedResponse models.Response
    if err := json.NewDecoder(r.Body).Decode(&updatedResponse); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    existingResponse.Content = updatedResponse.Content

    if err := db.DB.Save(&existingResponse).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := db.DB.Preload("User").Preload("Call").First(&existingResponse, id).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(existingResponse)
}




// DeleteResponse deletes a response by ID
func DeleteResponse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := db.DB.Delete(&models.Response{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
