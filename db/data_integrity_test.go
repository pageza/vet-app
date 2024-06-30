package db

import (
	"testing"

	"github.com/pageza/vet-app/models"
	"github.com/stretchr/testify/assert"
)

func setupDataIntegrity(t *testing.T) {
	SetupDB(t, &models.User{}, &models.Call{}, &models.Response{})
}

func TestForeignKeyConstraints(t *testing.T) {
	setupDataIntegrity(t)

	// Create a user
	user := models.User{Name: "John Doe", Email: "john@example.com"}
	result := DB.Create(&user)
	assert.NoError(t, result.Error)
	assert.NotZero(t, user.ID)

	// Create a call associated with the user
	call := models.Call{UserID: user.ID, Desc: "Help needed"}
	result = DB.Create(&call)
	assert.NoError(t, result.Error)
	assert.NotZero(t, call.ID)

	// Create a response associated with the call and user
	response := models.Response{CallID: call.ID, UserID: user.ID, Msg: "Response message"}
	result = DB.Create(&response)
	assert.NoError(t, result.Error)
	assert.NotZero(t, response.ID)

	// Delete the user and ensure cascading deletes
	result = DB.Delete(&user)
	assert.NoError(t, result.Error)

	var count int64
	DB.Model(&models.Call{}).Where("user_id = ?", user.ID).Count(&count)
	assert.Equal(t, int64(0), count)

	DB.Model(&models.Response{}).Where("user_id = ?", user.ID).Count(&count)
	assert.Equal(t, int64(0), count)

	DB.Model(&models.Response{}).Where("call_id = ?", call.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}
