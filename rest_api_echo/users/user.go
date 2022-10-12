package users

import (
	"fmt"
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

// Services -> Repository
type UserServiceInterface interface {
	GetAll() string
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAll() string {
	result, _ := s.repo.GetSth()
	return result
}

// Repository
type UserRepository struct {
}

func (r *UserRepository) GetSth() (string, error) {
	return fmt.Sprintf("Call get user"), fmt.Errorf("TODO next")
}
