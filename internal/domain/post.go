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
	ID        bson.ObjectID `bson:"_id" json:"id"`
	UserId    string        `bson:"userId" json:"userId"`
	Img       string        `bson:"img" json:"img"`
	Title     string        `bson:"title" json:"title"`
	Slug      string        `bson:"slug" json:"slug"`
	Desc      string        `bson:"desc" json:"desc"`
	Category  []string      `bson:"category" json:"category"`
	Content   string        `bson:"content" json:"content"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"updateAt"`
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
	GetPost(c context.Context, slug string) (*Post, error)
	CreatePost(c context.Context, post *Post) (*Post, error)
	DeletePost(c context.Context, id bson.ObjectID, userId string) error
	GetPostList(c context.Context, req *PostListQueryRequest) (*PostListResponse, error)
}

type PostRepository interface {
	FindBySlug(c context.Context, slug string) (*Post, error)
	Create(c context.Context, post *Post) (*Post, error)
	Delete(c context.Context, id bson.ObjectID, userId string) error
	FindByQuery(c context.Context, query *PostListQueryRequest) (*PostListResponse, error)
}
