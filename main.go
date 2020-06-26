package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"tableTennis/interfaces"
)

type User struct {
	gorm.Model
	Name    string `gorm:"not null"`
	StyleID uint8  `gorm:"not null"`
	Sex     uint8  `gorm:"not null"`
	Email   string `gorm:"type:varchar(100);unique_index;not null"`
}

func main() {
	db, err := gorm.Open("mysql", "root:password@/table_tennis?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
	interfaces.HandleRoutes()
}
