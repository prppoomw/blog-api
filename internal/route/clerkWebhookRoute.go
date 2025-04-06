package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prppoomw/blog-api/internal/config"
	"github.com/prppoomw/blog-api/internal/controller"
	"github.com/prppoomw/blog-api/internal/domain"
	"github.com/prppoomw/blog-api/internal/repository"
	"github.com/prppoomw/blog-api/internal/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewClerkRoute(timeout time.Duration, db mongo.Database, publicGroup *gin.RouterGroup, cfg *config.Config) {
	userCollection := db.Collection(domain.CollectionUsers)
	postCollection := db.Collection(domain.CollectionPosts)
	clerkWebhookSecret := cfg.ClerkWebhookSecret

	r := repository.NewClerkWebhookRepository(userCollection, postCollection)
	s := service.NewClerkWebhookService(r, timeout)
	c := controller.NewClerkWebhookController(s, clerkWebhookSecret)

	publicGroup.POST("/clerk", c.HandleWebhook)
}
