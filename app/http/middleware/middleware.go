package middleware

import (
	"context"
	"fmt"
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
  			c.JSON(http.StatusUnauthorized, gin.H{"error": "Id token not available"})
  			c.Abort()
  			return
		}

		token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
		  
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		// fmt.Print(token.Audience)
		// fmt.Print(token.Claims)
		// fmt.Print(token.Firebase)
		// fmt.Print(token.Issuer)
		// fmt.Print(token.Subject)
		fmt.Print(token.Firebase.Identities["phone"])
		// fmt.Print(token.Firebase.SignInProvider)
		// fmt.Print(token.Firebase.Tenant)
		c.Set("token", token.UID)
		c.Next()
	}
}
