package handlers

import (
    "github.com/pageza/vet-app/src/db"
    "github.com/pageza/vet-app/src/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var testDB *gorm.DB

func setup() {
    var err error
    dsn := "host=db user=testuser password=testpassword dbname=testdb port=5432 sslmode=disable"
    testDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    testDB.AutoMigrate(&models.User{}, &models.Call{}, &models.Response{})

    // Set the global DB connection to the test DB
    db.DB = testDB
}

func tearDown() {
    testDB.Exec("DROP TABLE IF EXISTS users, calls, responses CASCADE")
}
