package test

import (
	"Song_API/pkg/middleware"
	"Song_API/pkg/ratelimit"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRateLimit(t *testing.T) {
	assert := assert.New(t)
	rateRule := []ratelimit.Rule{
		{Path: "/api/test", Method: "GET", Capacity: 5, Rate: 1},
		{Path: "/api/test", Method: "POST", Capacity: 3, Rate: 1},
	}
	globalRule := ratelimit.Rule{Capacity: 7, Rate: 1}
	router := gin.Default()
	router.GET("/api/test", middleware.RateLimit(rateRule, globalRule), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	})
	router.POST("/api/test", middleware.RateLimit(rateRule, globalRule), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	})
	overallCounter := 0
	for i := 0; i < 100; i++ {
		req, _ := http.NewRequest("GET", "/api/test", nil)
		overallCounter++
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		if i < 5 && overallCounter <= 7 {
			assert.Equal(http.StatusOK, resp.Code)
		} else {
			assert.Equal(http.StatusTooManyRequests, resp.Code)
		}
		req, _ = http.NewRequest("POST", "/api/test", nil)
		overallCounter++
		resp = httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		if i < 3 && overallCounter <= 7 {
			assert.Equal(http.StatusOK, resp.Code)
		} else {
			assert.Equal(http.StatusTooManyRequests, resp.Code)
		}
	}
}
