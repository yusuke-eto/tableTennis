package handler

import (
	"github.com/labstack/echo/v4"
	"tableTennis/domain"
	"tableTennis/usecase"
)

// UserHandler Userに対するHandlerのインターフェース
type UserHandler interface {
	HandleUserSignup(c echo.Context) error
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

// NewUserHandler Userデータに関するHandlerを生成
func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

func (uh userHandler) HandleUserSignup(c echo.Context) error {
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	return uh.userUseCase.Insert(*user)
}
