package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"tableTennis/infrastructure/persistence"
	"tableTennis/interfaces/handler"
	"tableTennis/usecase"
)

type User struct {
	gorm.Model
	Name    string `gorm:"not null"`
	StyleID uint8  `gorm:"not null"`
	Sex     uint8  `gorm:"not null"`
	Email   string `gorm:"type:varchar(100);unique_index;not null"`
}

func main() {
	// 依存関係を注入
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/user", userHandler.HandleUserSignup)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
