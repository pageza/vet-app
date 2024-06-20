package models

import (
	"time"
	"gorm.io/gorm"
)

type Call struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	Title         string         `json:"title"`
	Content       string         `json:"content"`
	Status        string         `json:"status"`
	IsOpenThread  bool           `json:"is_open_thread"`
	UserID        uint           `json:"user_id"`
	User          User           `json:"user"`
	Responses     []Response     `json:"responses"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
