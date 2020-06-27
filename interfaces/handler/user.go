package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tableTennis/usecase"
)

// Userに対するHandlerのインターフェース
type UserHandler interface {
	HandleUserSignup(c echo.Context) error
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

// Userデータに関するHandlerを生成
func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

func (uh userHandler) HandleUserSignup(c echo.Context) error {
	return c.String(http.StatusOK, "hoge")
}
