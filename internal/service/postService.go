package service

import (
	"context"
	"time"

	"github.com/prppoomw/blog-api/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PostService struct {
	postRepository domain.PostRepository
	contextTimeout time.Duration
}

func NewPostService(postRepository domain.PostRepository, timeout time.Duration) domain.PostUsecase {
	return &PostService{
		postRepository: postRepository,
		contextTimeout: timeout,
	}
}

func (s *PostService) GetPost(c context.Context, slug string) (*domain.Post, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.postRepository.FindBySlug(ctx, slug)
}

func (s *PostService) CreatePost(c context.Context, post *domain.Post) (*domain.Post, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.postRepository.Create(ctx, post)
}

func (s *PostService) DeletePost(c context.Context, id bson.ObjectID, userId string) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.postRepository.Delete(ctx, id, userId)
}

func (s *PostService) GetPostList(c context.Context, req *domain.PostListQueryRequest) (*domain.PostListResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.postRepository.FindByQuery(ctx, req)
}
