package domain

import "context"

type ClerkWebhookRepository interface {
	CreateUser(c context.Context, user *User) error
	DeleteUser(c context.Context, clerkUserId string) error
	DeletePostsByUser(c context.Context, userId string) error
}

type ClerkWebhookUsecase interface {
	HandleWebhook(c context.Context, payload []byte, headers map[string][]string) error
}
