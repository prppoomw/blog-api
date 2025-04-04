package service

import (
	"github.com/prppoomw/blog-api/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PostResponse struct {
}

type PostListResponse struct {
}

type PostListQueryRequest struct {
}

type PostService interface {
	GetPost(*string) (*PostResponse, error)
	GetPostList(*PostListQueryRequest) (*PostListResponse, error)
	CreatePost(*model.Post) (*PostResponse, error)
	DeletePost(*bson.ObjectID) error
}
