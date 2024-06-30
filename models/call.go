package models

type Call struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `gorm:"not null"`
	Desc   string `gorm:"size:255"`
	User   User   `gorm:"constraint:OnDelete:CASCADE;"`
}
