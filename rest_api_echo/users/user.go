package users

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserResponse struct {
	Message string `json:"message"`
}

func GetUserHandler(service *UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := UserResponse{
			Message: service.GetAll(),
		}
		return c.JSON(http.StatusOK, res)
	}
}
