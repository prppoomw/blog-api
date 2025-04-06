package middleware

import (
	"net/http"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gin-gonic/gin"
	"github.com/prppoomw/blog-api/internal/config"
	"github.com/prppoomw/blog-api/internal/domain"
)

func ClerkAuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		clerk.SetKey(cfg.ClerkKey)
		sessionToken := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		claims, err := jwt.Verify(c.Request.Context(), &jwt.VerifyParams{
			Token: sessionToken,
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "unauthorized"})
		}
		usr, err := user.Get(c.Request.Context(), claims.Subject)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "failed to get user"})
			return
		}
		//c.Set("claims", claims)
		c.Set("userId", usr.ID)
		c.Next()
	}
}
