package models

type Response struct {
	ID     uint   `gorm:"primaryKey"`
	CallID uint   `gorm:"not null"`
	UserID uint   `gorm:"not null"`
	Msg    string `gorm:"size:255"`
	Call   Call   `gorm:"constraint:OnDelete:CASCADE;"`
	User   User   `gorm:"constraint:OnDelete:CASCADE;"`
}
