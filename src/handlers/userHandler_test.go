package handlers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gorilla/mux"
    "github.com/pageza/vet-app/src/models"
)

func TestGetUsers(t *testing.T) {
    setup()
    defer tearDown()

    testDB.Create(&models.User{Name: "Test User", Email: "testuser@example.com"})

    req, err := http.NewRequest("GET", "/users", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(GetUsers)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var users []models.User
    if err := json.NewDecoder(rr.Body).Decode(&users); err != nil {
        t.Errorf("handler returned invalid body: %v", err)
    }

    if len(users) != 1 || users[0].Name != "Test User" {
        t.Errorf("handler returned unexpected body: got %v", users)
    }
}

func TestCreateUser(t *testing.T) {
    setup()
    defer tearDown()

    newUser := &models.User{
        Name:  "Test User",
        Email: "testuser@example.com",
    }
    jsonUser, _ := json.Marshal(newUser)

    req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(CreateUser)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var user models.User
    if err := json.NewDecoder(rr.Body).Decode(&user); err != nil {
        t.Errorf("handler returned invalid body: %v", err)
    }

    if user.Name != newUser.Name || user.Email != newUser.Email {
        t.Errorf("handler returned unexpected body: got %v want %v", user, newUser)
    }
}

func TestGetUser(t *testing.T) {
    setup()
    defer tearDown()

    testDB.Create(&models.User{Name: "Test User", Email: "testuser@example.com"})

    req, err := http.NewRequest("GET", "/users/1", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    router := mux.NewRouter()
    router.HandleFunc("/users/{id}", GetUser)
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var user models.User
    if err := json.NewDecoder(rr.Body).Decode(&user); err != nil {
        t.Errorf("handler returned invalid body: %v", err)
    }

    if user.Name != "Test User" || user.Email != "testuser@example.com" {
        t.Errorf("handler returned unexpected body: got %v", user)
    }
}

func TestUpdateUser(t *testing.T) {
    setup()
    defer tearDown()

    testDB.Create(&models.User{Name: "Test User", Email: "testuser@example.com"})

    updatedUser := &models.User{
        Name:  "Updated User",
        Email: "updateduser@example.com",
    }
    jsonUser, _ := json.Marshal(updatedUser)

    req, err := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(jsonUser))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    router := mux.NewRouter()
    router.HandleFunc("/users/{id}", UpdateUser)
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var user models.User
    if err := json.NewDecoder(rr.Body).Decode(&user); err != nil {
        t.Errorf("handler returned invalid body: %v", err)
    }

    if user.Name != updatedUser.Name || user.Email != updatedUser.Email {
        t.Errorf("handler returned unexpected body: got %v want %v", user, updatedUser)
    }
}

func TestDeleteUser(t *testing.T) {
    setup()
    defer tearDown()

    testDB.Create(&models.User{Name: "Test User", Email: "testuser@example.com"})

    req, err := http.NewRequest("DELETE", "/users/1", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    router := mux.NewRouter()
    router.HandleFunc("/users/{id}", DeleteUser)
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusNoContent {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
    }
}
