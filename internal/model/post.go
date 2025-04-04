package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
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
