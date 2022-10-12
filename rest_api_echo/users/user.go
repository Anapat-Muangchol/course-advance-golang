package users

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
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
type IService interface {
	GetAll() string
}

type UserService struct {
	repo IRepository
}

func NewUserService(repo IRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAll() string {
	result, _ := s.repo.GetSth()
	return result
}

// Repository
type IRepository interface {
	GetSth() (string, error)
}

type UserRepository struct {
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client) UserRepository {
	return UserRepository{
		client: client,
	}
}

func (r *UserRepository) GetSth() (string, error) {

	return "TODO next", fmt.Errorf("TODO next")
}
