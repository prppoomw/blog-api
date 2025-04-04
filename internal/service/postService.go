package service

import (
	"github.com/prppoomw/blog-api/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PostService struct {
	postRepository domain.PostRepository
}

func NewPostService(postRepository domain.PostRepository) domain.PostUsecase {
	return &PostService{
		postRepository: postRepository,
	}
}

func (s *PostService) GetPost(slug string) (domain.PostResponse, error) {
	return domain.PostResponse{}, nil
}

func (s *PostService) GetPostList(req domain.PostListQueryRequest) (domain.PostListResponse, error) {
	return domain.PostListResponse{}, nil
}

func (s *PostService) CreatePost(post domain.Post) (domain.PostResponse, error) {
	return domain.PostResponse{}, nil
}

func (s *PostService) DeletePost(id bson.ObjectID) error {
	return nil
}
