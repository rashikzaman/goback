package middleware

import (
	"context"

	"locally/goback/app/domain"
	"locally/goback/config"
	"locally/goback/db"
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
			
		phoneNumber := token.Claims["phone_number"].(string)
		var users []domain.User
		usersWithThisPhoneNumber := db.GetDb().Where("phone_number = ?", phoneNumber).Find(&users)

		if(usersWithThisPhoneNumber.RowsAffected <= 0) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			c.Abort()
			return
		}
		
		c.Set("token", token.UID)
		c.Next()
	}
}
