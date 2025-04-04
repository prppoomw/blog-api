package repository

import (
	"github.com/prppoomw/blog-api/internal/domain"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	db         *mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database, collection string) domain.UserRepository {
	return &UserRepository{
		db:         db,
		collection: collection,
	}
}
