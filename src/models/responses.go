package models

import (
	"time"
	"gorm.io/gorm"
)

type Response struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Content   string         `json:"content"`
	CallID    uint           `json:"call_id"`
	Call      Call           `json:"call"`
	UserID    uint           `json:"user_id"`
	User      User           `json:"user"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
