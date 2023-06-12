package middleware

import (
	"Song_API/pkg/cache"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Authorization is a middleware to authenticate the user.
func Authorization(roles []string, cache cache.Cache) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		var tokenString string
		var cacheErr error
		if authHeader == "" {
			tokenString, cacheErr = cache.Get("token")
			if cacheErr != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
				c.Abort()
				return
			}
		} else {
			tokenString = strings.Split(authHeader, "Bearer ")[1]
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("AUTH_KEY")), nil
		})
		// token.Valid takes standard claims into consideration. So we don't need to check for expiration explicitly.
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid/Expired token"})
			c.Abort()
			return
		}
		claims, _ := token.Claims.(jwt.MapClaims)
		user := claims["user"].(string)
		userRole := claims["role"].(string)
		exp := claims["exp"].(float64)
		if !strings.Contains(strings.Join(roles, ","), userRole) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid role"})
			c.Abort()
			return
		}
		remTime := time.Until(time.Unix(int64(exp), 0))
		cache.Set("token", tokenString, remTime)
		ctx := context.WithValue(c.Request.Context(), "user", user)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
