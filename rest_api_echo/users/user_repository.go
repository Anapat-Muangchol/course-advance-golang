package users

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository
type IRepository interface {
	GetSth() (string, error)
}

type IRepositoryV2 interface {
	IRepository
	GetSth2() (string, error)
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
