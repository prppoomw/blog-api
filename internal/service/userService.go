package service

import (
	"github.com/prppoomw/blog-api/internal/domain"
)

type UserService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) domain.UserUsecase {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) GetUserSavedPostList(user domain.User) ([]string, error) {
	return nil, nil
}

func (s *UserService) SavePost(user domain.User, id string) error {
	return nil
}
