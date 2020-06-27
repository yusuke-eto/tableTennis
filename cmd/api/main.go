package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"tableTennis/infrastructure/persistence"
	"tableTennis/interfaces/handler"
	"tableTennis/usecase"
)

func main() {
	// 依存関係を注入
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	e := echo.New()

	// Middlewareの設定 時間、メソッド、uri、ステータスコードは出しておく
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time: ${time_rfc3339}, method: ${method}, uri: ${uri}, status: ${status}",
	}))
	e.Use(middleware.Recover())

	// Routes
	e.POST("/user", userHandler.HandleUserSignup)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
