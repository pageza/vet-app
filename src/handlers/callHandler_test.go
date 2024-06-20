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

func TestGetCalls(t *testing.T) {
    setup()
    defer tearDown()

    // Create a user first
    user := models.User{Name: "Test User", Email: "testuser@example.com"}
    testDB.Create(&user)

    // Create a call associated with the user
    testDB.Create(&models.Call{Title: "Test Call", Content: "Test Content", Status: "open", UserID: user.ID})

    req, err := http.NewRequest("GET", "/calls", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(GetCalls)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var calls []models.Call
    if err := json.NewDecoder(rr.Body).Decode(&calls); err != nil {
        t.Errorf("handler returned invalid body: %v", err)
    }

    if len(calls) != 1 || calls[0].Title != "Test Call" {
        t.Errorf("handler returned unexpected body: got %v", calls)
    }
}

func TestCreateCall(t *testing.T) {
    setup()
    defer tearDown()

    // Create a user first
    user := models.User{Name: "Test User", Email: "testuser@example.com"}
    testDB.Create(&user)

    newCall := &models.Call{
        Title:   "Test Call",
        Content: "Test Content",
        Status:  "open",
        UserID:  user.ID,
    }
    jsonCall, _ := json.Marshal(newCall)

    req, err := http.NewRequest("POST", "/calls", bytes.NewBuffer(jsonCall))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(CreateCall)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var call models.Call
    if err := json.NewDecoder(rr.Body).Decode(&call); err != nil {
        t.Errorf("handler returned invalid body: %v", err)
    }

    if call.Title != newCall.Title || call.Content != newCall.Content || call.Status != newCall.Status {
        t.Errorf("handler returned unexpected body: got %v want %v", call, newCall)
    }
}

func TestGetCall(t *testing.T) {
    setup()
    defer tearDown()

    // Create a user first
    user := models.User{Name: "Test User", Email: "testuser@example.com"}
    testDB.Create(&user)

    // Create a call associated with the user
    testDB.Create(&models.Call{Title: "Test Call", Content: "Test Content", Status: "open", UserID: user.ID})

    req, err := http.NewRequest("GET", "/calls/1", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    router := mux.NewRouter()
    router.HandleFunc("/calls/{id}", GetCall)
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var call models.Call
    if err := json.NewDecoder(rr.Body).Decode(&call); err != nil {
        t.Errorf("handler returned invalid body: %v", err)
    }

    if call.Title != "Test Call" || call.Content != "Test Content" || call.Status != "open" {
        t.Errorf("handler returned unexpected body: got %v", call)
    }
}

func TestUpdateCall(t *testing.T) {
    setup()
    defer tearDown()

    // Create a user first
    user := models.User{Name: "Test User", Email: "testuser@example.com"}
    testDB.Create(&user)

    // Create a call associated with the user
    testDB.Create(&models.Call{Title: "Test Call", Content: "Test Content", Status: "open", UserID: user.ID})

    updatedCall := &models.Call{
        Title:   "Updated Call",
        Content: "Updated Content",
        Status:  "closed",
        UserID:  user.ID,
    }
    jsonCall, _ := json.Marshal(updatedCall)

    req, err := http.NewRequest("PUT", "/calls/1", bytes.NewBuffer(jsonCall))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    router := mux.NewRouter()
    router.HandleFunc("/calls/{id}", UpdateCall)
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var call models.Call
    if err := json.NewDecoder(rr.Body).Decode(&call); err != nil {
        t.Errorf("handler returned invalid body: %v", err)
    }

    if call.Title != updatedCall.Title || call.Content != updatedCall.Content || call.Status != updatedCall.Status {
        t.Errorf("handler returned unexpected body: got %v want %v", call, updatedCall)
    }
}

func TestDeleteCall(t *testing.T) {
    setup()
    defer tearDown()

    // Create a user first
    user := models.User{Name: "Test User", Email: "testuser@example.com"}
    testDB.Create(&user)

    // Create a call associated with the user
    testDB.Create(&models.Call{Title: "Test Call", Content: "Test Content", Status: "open", UserID: user.ID})

    req, err := http.NewRequest("DELETE", "/calls/1", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    router := mux.NewRouter()
    router.HandleFunc("/calls/{id}", DeleteCall)
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusNoContent {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
    }
}
