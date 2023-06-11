package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Authorization is a middleware to authenticate the user.
func Authorization(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		tokenString := strings.Split(authHeader, "Bearer ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("AUTH_KEY")), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		user := token.Claims.(jwt.MapClaims)["user"].(string)
		userRole := token.Claims.(jwt.MapClaims)["role"].(string)
		if !strings.Contains(strings.Join(roles, ","), userRole) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid role"})
			c.Abort()
			return
		}
		ctx := context.WithValue(c.Request.Context(), "user", user)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
