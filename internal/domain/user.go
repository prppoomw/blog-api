package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionUsers = "users"
)

type User struct {
	Id          bson.ObjectID `bson:"_id"`
	ClerkUserId string        `bson:"clerkUserId"`
	Username    string        `bson:"username"`
	Email       string        `bson:"email"`
	Img         string        `bson:"img"`
	SavedPosts  []string      `bson:"savedPosts"`
	CreatedAt   time.Time     `bson:"createdAt"`
	UpdatedAt   time.Time     `bson:"updatedAt"`
}

type UserUsecase interface {
	GetUserSavedPostList(User) ([]string, error)
	SavePost(User, string) error
}

type UserRepository interface {
}
