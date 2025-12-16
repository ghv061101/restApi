package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID   uint      `json:"id" gorm:"primaryKey"`
	Name string    `json:"name" validate:"required,min=2"`
	Dob  time.Time `json:"dob" validate:"required" gorm:"type:date"`
}

func MigrateUsers(db *gorm.DB) error {
	return db.AutoMigrate(&Users{})
}

// Age calculates the current age in years based on DOB
func (u *Users) Age() int {
	if u == nil {
		return 0
	}
	now := time.Now()
	years := now.Year() - u.Dob.Year()
	// adjust if birthday hasn't happened yet this year
	birthdayThisYear := time.Date(now.Year(), u.Dob.Month(), u.Dob.Day(), 0, 0, 0, 0, time.UTC)
	if now.Before(birthdayThisYear) {
		years--
	}
	return years
}
