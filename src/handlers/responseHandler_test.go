package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pageza/vet-app/src/models"
)

func TestCreateResponse(t *testing.T) {
	setup()
	defer tearDown()

	// Create a user first
	user := models.User{Name: "Test User", Email: "testuser@example.com"}
	testDB.Create(&user)

	// Create a call associated with the user
	call := models.Call{Title: "Test Call", Content: "Test Content", Status: "open", UserID: user.ID}
	testDB.Create(&call)

	newResponse := &models.Response{
		Content: "Test Response",
		UserID:  user.ID,
		CallID:  call.ID,
	}
	jsonResponse, _ := json.Marshal(newResponse)

	req, err := http.NewRequest("POST", "/calls/"+strconv.Itoa(int(call.ID))+"/responses", bytes.NewBuffer(jsonResponse))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/calls/{call_id}/responses", CreateResponse)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response models.Response
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("handler returned invalid body: %v", err)
	}

	if response.Content != newResponse.Content || response.UserID != newResponse.UserID || response.CallID != newResponse.CallID {
		t.Errorf("handler returned unexpected body: got %v want %v", response, newResponse)
	}
}

func TestGetResponses(t *testing.T) {
	setup()
	defer tearDown()

	// Create a user first
	user := models.User{Name: "Test User", Email: "testuser@example.com"}
	testDB.Create(&user)

	// Create a call associated with the user
	call := models.Call{Title: "Test Call", Content: "Test Content", Status: "open", UserID: user.ID}
	testDB.Create(&call)

	// Create a response associated with the call and user
	response := models.Response{Content: "Test Response", UserID: user.ID, CallID: call.ID}
	testDB.Create(&response)

	req, err := http.NewRequest("GET", "/calls/"+strconv.Itoa(int(call.ID))+"/responses", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/calls/{call_id}/responses", GetResponses)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responses []models.Response
	if err := json.NewDecoder(rr.Body).Decode(&responses); err != nil {
		t.Errorf("handler returned invalid body: %v", err)
	}

	if len(responses) != 1 || responses[0].Content != "Test Response" {
		t.Errorf("handler returned unexpected body: got %v", responses)
	}
}

func TestGetResponsesForUser(t *testing.T) {
	setup()
	defer tearDown()

	// Create a user first
	user := models.User{Name: "Test User", Email: "testuser@example.com"}
	testDB.Create(&user)

	// Create a call associated with the user
	call := models.Call{Title: "Test Call", Content: "Test Content", Status: "open", UserID: user.ID}
	testDB.Create(&call)

	// Create a response associated with the call and user
	response := models.Response{Content: "Test Response", UserID: user.ID, CallID: call.ID}
	testDB.Create(&response)

	req, err := http.NewRequest("GET", "/users/"+strconv.Itoa(int(user.ID))+"/responses", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{user_id}/responses", GetResponsesForUser)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responses []models.Response
	if err := json.NewDecoder(rr.Body).Decode(&responses); err != nil {
		t.Errorf("handler returned invalid body: %v", err)
	}

	if len(responses) != 1 || responses[0].Content != "Test Response" {
		t.Errorf("handler returned unexpected body: got %v", responses)
	}
}

func TestGetResponse(t *testing.T) {
	setup()
	defer tearDown()

	// Create a user first
	user := models.User{Name: "Test User", Email: "testuser@example.com"}
	testDB.Create(&user)

	// Create a call associated with the user
	call := models.Call{Title: "Test Call", Content: "Test Content", Status: "open", UserID: user.ID}
	testDB.Create(&call)

	// Create a response associated with the call and user
	response := models.Response{Content: "Test Response", UserID: user.ID, CallID: call.ID}
	testDB.Create(&response)

	req, err := http.NewRequest("GET", "/responses/"+strconv.Itoa(int(response.ID)), nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/responses/{id}", GetResponse)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var res models.Response
	if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
		t.Errorf("handler returned invalid body: %v", err)
	}

	if res.Content != "Test Response" || res.UserID != user.ID || res.CallID != call.ID {
		t.Errorf("handler returned unexpected body: got %v", res)
	}
}

func TestUpdateResponse(t *testing.T) {
    setup()
    defer tearDown()

    // Create a user first
    user := models.User{Name: "Test User", Email: "testuser@example.com"}
    testDB.Create(&user)

    // Create a call associated with the user
    call := models.Call{Title: "Test Call", Content: "Test Content", Status: "open", UserID: user.ID}
    testDB.Create(&call)

    // Create a response associated with the call and user
    response := models.Response{Content: "Test Response", UserID: user.ID, CallID: call.ID}
    testDB.Create(&response)

    // Ensure the response is created
    var createdResponse models.Response
    testDB.First(&createdResponse, response.ID)

    updatedResponse := &models.Response{
        Content: "Updated Response",
    }
    jsonResponse, _ := json.Marshal(updatedResponse)

    req, err := http.NewRequest("PUT", "/responses/"+strconv.Itoa(int(response.ID)), bytes.NewBuffer(jsonResponse))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    router := mux.NewRouter()
    router.HandleFunc("/responses/{id}", UpdateResponse)
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var res models.Response
    if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
        t.Errorf("handler returned invalid body: %v", err)
    }

    // Reload the response from the database to ensure it was updated
    var dbResponse models.Response
    if err := testDB.Preload("User").Preload("Call").First(&dbResponse, response.ID).Error; err != nil {
        t.Fatalf("failed to load response from the database: %v", err)
    }

    // Check if the response has been updated correctly
    if dbResponse.Content != "Updated Response" || dbResponse.UserID != user.ID || dbResponse.CallID != call.ID {
        t.Errorf("handler returned unexpected body: got %v want %v", dbResponse, updatedResponse)
    }
}



func TestDeleteResponse(t *testing.T) {
	setup()
	defer tearDown()

	// Create a user first
	user := models.User{Name: "Test User", Email: "testuser@example.com"}
	testDB.Create(&user)

	// Create a call associated with the user
	call := models.Call{Title: "Test Call", Content: "Test Content", Status: "open", UserID: user.ID}
	testDB.Create(&call)

	// Create a response associated with the call and user
	response := models.Response{Content: "Test Response", UserID: user.ID, CallID: call.ID}
	testDB.Create(&response)

	req, err := http.NewRequest("DELETE", "/responses/"+strconv.Itoa(int(response.ID)), nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/responses/{id}", DeleteResponse)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
}
