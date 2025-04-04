package repository

import (
	"github.com/prppoomw/blog-api/internal/domain"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PostRepository struct {
	db         *mongo.Database
	collection string
}

func NewPostRepository(db *mongo.Database, collection string) domain.PostRepository {
	return &PostRepository{
		db:         db,
		collection: collection,
	}
}
