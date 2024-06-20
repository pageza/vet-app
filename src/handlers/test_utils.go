package handlers

import (
    "github.com/DATA-DOG/go-sqlmock"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func SetupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
    db, mock, err := sqlmock.New()
    if err != nil {
        return nil, nil, err
    }

    dsn := "user=mock dbname=mock sslmode=disable"
    gdb, err := gorm.Open(postgres.New(postgres.Config{
        DSN:                  dsn,
        DriverName:           "postgres",
        Conn:                 db,
        PreferSimpleProtocol: true,
    }), &gorm.Config{})
    if err != nil {
        return nil, nil, err
    }

    return gdb, mock, nil
}
