package service

import (
	"context"
	"errors"
	"time"

	"github.com/prppoomw/blog-api/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ClerkWebhookService struct {
	clerkWebhookRepository domain.ClerkWebhookRepository
	contextTimeout         time.Duration
}

func NewClerkWebhookService(clerkWebhookRepository domain.ClerkWebhookRepository, timeout time.Duration) domain.ClerkWebhookUsecase {
	return &ClerkWebhookService{
		clerkWebhookRepository: clerkWebhookRepository,
		contextTimeout:         timeout,
	}
}

func (s *ClerkWebhookService) HandleWebhook(c context.Context, payload map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	eventType, ok := payload["type"].(string)
	if !ok {
		return errors.New("cant extract event type")
	}

	data, ok := payload["data"].(map[string]interface{})
	if !ok {
		return errors.New("invalid data object")
	}
	switch eventType {
	case "user.created":
		clerkUserId, _ := data["id"].(string)
		username, _ := data["username"].(string)
		profileImageURL, _ := data["profile_image_url"].(string)
		email := ""
		if emailAddresses, ok := data["email_addresses"].([]interface{}); ok && len(emailAddresses) > 0 {
			if firstEmail, ok := emailAddresses[0].(map[string]interface{}); ok {
				email, _ = firstEmail["email_address"].(string)
			}
		}
		if username == "" {
			username = email
		}
		id := bson.NewObjectID()
		now := time.Now()
		newUser := domain.User{
			Id:          id,
			ClerkUserId: clerkUserId,
			Username:    username,
			Email:       email,
			Img:         profileImageURL,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		return s.clerkWebhookRepository.CreateUser(ctx, &newUser)

	case "user.deleted":
		clerkUserId, _ := data["id"].(string)
		res, err := s.clerkWebhookRepository.DeleteUser(c, clerkUserId)
		if err != nil {
			return err
		}

		if res.DeletedCount > 0 {
			err := s.clerkWebhookRepository.DeletePostsByUser(c, clerkUserId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
