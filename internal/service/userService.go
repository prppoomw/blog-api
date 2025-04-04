package service

import "github.com/prppoomw/blog-api/internal/model"

type UserService interface {
	GetUserSavedPostList(*model.User) (*[]string, error)
	SavePost(*model.User, *string) error
}
