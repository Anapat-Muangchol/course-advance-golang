package users

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
