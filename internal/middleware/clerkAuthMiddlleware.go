package middleware

import (
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/gin-gonic/gin"
)

func ClerkAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

		protectedHandler := clerkhttp.WithHeaderAuthorization()(dummyHandler)

		protectedHandler.ServeHTTP(c.Writer, c.Request)

		claims, ok := clerk.SessionClaimsFromContext(c.Request.Context())
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
