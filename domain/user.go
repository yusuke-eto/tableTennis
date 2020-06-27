package domain

import (
	"github.com/jinzhu/gorm"
)

// User struct is base of user model
type User struct {
	gorm.Model
	Name    string `gorm:"not null"`
	StyleID uint8  `gorm:"not null"`
	Sex     uint8  `gorm:"not null"`
	Email   string `gorm:"type:varchar(100);unique_index;not null"`
}
