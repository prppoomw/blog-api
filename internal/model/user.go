package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID          bson.ObjectID `bson:"_id"`
	ClerkUserID string        `bson:"clerkUserId"`
	Username    string        `bson:"username"`
	Email       string        `bson:"email"`
	Img         string        `bson:"img"`
	SavedPosts  []string      `bson:"savedPosts"`
	CreatedAt   time.Time     `bson:"createdAt"`
	UpdatedAt   time.Time     `bson:"updatedAt"`
}
