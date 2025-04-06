package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prppoomw/blog-api/internal/domain"
	svix "github.com/svix/svix-webhooks/go"
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
	headers := c.Request.Header
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid payload"})
		return
	}

	if ctrl.clerkWebhookSecret == "" {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Webhook secret needed!"})
		return
	}

	wh, err := svix.NewWebhook(ctrl.clerkWebhookSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Invalid webhook secret"})
		log.Fatal(err)
		return
	}

	err = wh.Verify(payload, headers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Webhook verification failed!"})
		log.Fatal(err)
		return
	}

	var payloadData map[string]interface{}
	err = json.Unmarshal(payload, &payloadData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Failed to parse webhook payload"})
		log.Fatal(err)
		return
	}

	err = ctrl.clerkWebhookService.HandleWebhook(c, payloadData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Webhook received"})
}
