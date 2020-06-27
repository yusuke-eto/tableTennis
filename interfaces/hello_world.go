package interfaces

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// 公開したい propertyは public にしておくこと
type hello struct {
	Name string
	Age  uint8
}

func helloWorldHandler(c echo.Context) error {
	h := &hello{Name: "hoge", Age: 33}
	return c.JSON(http.StatusOK, h)
}
