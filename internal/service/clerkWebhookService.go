package service

import (
	"context"
	"time"

	"github.com/prppoomw/blog-api/internal/domain"
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

func (s *ClerkWebhookService) HandleWebhook(c context.Context, payload []byte, headers map[string][]string) error {
	return nil
}
