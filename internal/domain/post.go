package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionPosts = "posts"
)

type Post struct {
	ID        bson.ObjectID `bson:"_id"`
	User      bson.ObjectID `bson:"user"`
	Img       string        `bson:"img"`
	Title     string        `bson:"title"`
	Slug      string        `bson:"slug"`
	Desc      string        `bson:"desc"`
	Category  []string      `bson:"category"`
	Content   string        `bson:"content"`
	CreatedAt time.Time     `bson:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt"`
}

type PostListResponse struct {
	Posts   []Post `json:"posts"`
	HasMore bool   `json:"hasMore"`
}

type PostListQueryRequest struct {
	Page     int
	Limit    int
	Category string
	Author   string
	Search   string
}

type PostUsecase interface {
	GetPost(slug string) (Post, error)
	GetPostList(req PostListQueryRequest) (PostListResponse, error)
	CreatePost(post Post) (Post, error)
	DeletePost(id bson.ObjectID, user bson.ObjectID) error
}

type PostRepository interface {
	FindBySlug(c context.Context, slug string) (*Post, error)
	Create(c context.Context, post *Post) (*Post, error)
	Delete(c context.Context, id bson.ObjectID, user bson.ObjectID) error
	FindByQuery(c context.Context, query *PostListQueryRequest) (*PostListResponse, error)
}
