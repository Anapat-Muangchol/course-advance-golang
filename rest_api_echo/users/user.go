package users

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserResponse struct {
	Message string `json:"message"`
}

func UserHandler(c echo.Context) error {
	res := UserResponse{
		Message: "Get with user",
	}
	return c.JSON(http.StatusOK, res)
}
