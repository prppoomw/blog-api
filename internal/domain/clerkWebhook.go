package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ClerkWebhookRepository interface {
	CreateUser(c context.Context, user *User) error
	DeleteUser(c context.Context, clerkUserId string) (*mongo.DeleteResult, error)
	DeletePostsByUser(c context.Context, userId string) error
}

type ClerkWebhookUsecase interface {
	HandleWebhook(c context.Context, payload map[string]interface{}) error
}
