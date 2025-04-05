package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prppoomw/blog-api/internal/domain"
)

type ClerkWebhookController struct {
	clerkWebhookService domain.ClerkWebhookUsecase
	clerkWebhookSecret  string
}

func NewClerkWebhookController(clerkWebhookService domain.ClerkWebhookUsecase, clerkWebhookSecret string) *ClerkWebhookController {
	return &ClerkWebhookController{
		clerkWebhookService: clerkWebhookService,
		clerkWebhookSecret:  clerkWebhookSecret,
	}
}

func (ctrl *ClerkWebhookController) HandleWebhook(c *gin.Context) {

}
