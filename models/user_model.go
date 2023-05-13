package models

import "time"

type User struct {
	ID        uint `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string `json:"email" gorm:"unique;not null;default:null"`
	Password  string `json:"password" gorm:"not null;default:null"`
}
