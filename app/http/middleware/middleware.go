package middleware

import (
	"context"
	"locally/goback/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		firebaseAuth := config.SetupFirebase()
		authorizationToken := c.GetHeader("Authorization")

 		idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

		if idToken == "" {
  			c.JSON(http.StatusBadRequest, gin.H{"error": "Id token not available"})
  			c.Abort()
  			return
		}

		token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
		  
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		c.Set("token", token.UID)
		c.Next()
	}
}
