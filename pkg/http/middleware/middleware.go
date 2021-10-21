package middleware

import "github.com/gin-gonic/gin"

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizedToken := c.Request.Header["Token"]

		if authorizedToken == nil {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "no authentication token found",
			})
			return
		}
		c.Set("token", authorizedToken)
		c.Next()
	}
}
