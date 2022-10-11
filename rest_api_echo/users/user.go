package users

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserResponse struct {
	Message string `json:"message"`
}

func UserHandler(c echo.Context) error {
	us := NewUserService()

	res := UserResponse{
		Message: us.GetAll(),
	}
	return c.JSON(http.StatusOK, res)
}

// Service
type UserServiceInterface interface {
	GetAll() string
}

type UserService struct {
}

func (us *UserService) GetAll() string {
	return fmt.Sprintf("Get with user")
}

func NewUserService() *UserService {
	return &UserService{}
}
