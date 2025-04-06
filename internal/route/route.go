package route

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prppoomw/blog-api/internal/config"
	"github.com/prppoomw/blog-api/internal/middleware"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func Setup(cfg *config.Config, timeout time.Duration, db *mongo.Database, gin *gin.Engine) {
	gin.Use(GlobalErrorHandler())

	privateRouter := gin.Group("")
	publicRouter := gin.Group("")

	privateRouter.Use(cors.Default())
	publicRouter.Use(cors.Default())

	privateRouter.Use(middleware.ClerkAuthMiddleware(cfg))

	NewClerkRoute(timeout, *db, publicRouter, cfg)
	NewPostRoute(timeout, *db, privateRouter, publicRouter)
}

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			statusCode := c.Writer.Status()
			if statusCode == http.StatusOK {
				statusCode = http.StatusInternalServerError
			}

			c.JSON(statusCode, gin.H{
				"message": err.Error(),
				"status":  statusCode,
				"stack":   err,
			})
		}
	}
}
