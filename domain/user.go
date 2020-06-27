package domain

import (
	"github.com/jinzhu/gorm"
)

// User struct is base of user model
type User struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null"`
	StyleID uint8  `json:"styled_id" gorm:"not null"`
	Sex     uint8  `json:"sex" gorm:"not null"`
	Email   string `json:"email" gorm:"type:varchar(100);unique_index;not null"`
}
