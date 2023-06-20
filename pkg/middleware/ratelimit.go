package middleware

import (
	"Song_API/pkg/ratelimit"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RateLimit(rateRules []ratelimit.Rule, globalRule ratelimit.Rule) gin.HandlerFunc {
	return func(c *gin.Context) {
		client := getClient(c)
		rule := getRule(c, rateRules)
		bucket := ratelimit.GetBucket(client, rule)
		globalBucket := ratelimit.GetBucket("global", globalRule)
		if !bucket.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"message": "Too many requests"})
			c.Abort()
			return
		}
		if !globalBucket.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"message": "Server is busy"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func getClient(c *gin.Context) string {
	ip := c.ClientIP()
	path := c.Request.URL.Path
	method := c.Request.Method
	data := fmt.Sprintf("%s-%s-%s", ip, path, method)
	return data
}

func isPathMatched(c *gin.Context, rulePath string) bool {
	actualPath := c.Request.URL.Path
	ruleSegments := strings.Split(rulePath, "/")
	actualSegments := strings.Split(actualPath, "/")
	if len(ruleSegments) != len(actualSegments) {
		return false
	}
	for i := 0; i < len(ruleSegments); i++ {
		if ruleSegments[i] == actualSegments[i] {
			continue
		}
		if strings.HasPrefix(ruleSegments[i], ":") {
			param := strings.TrimPrefix(ruleSegments[i], ":")
			actualValue := actualSegments[i]
			if value, exists := c.Params.Get(param); exists && value == actualValue {
				continue
			}
		}
		return false
	}
	return true
}

func getRule(c *gin.Context, rateRules []ratelimit.Rule) ratelimit.Rule {
	path := c.Request.URL.Path
	method := c.Request.Method
	for _, rule := range rateRules {
		if isPathMatched(c, rule.Path) && rule.Method == method {
			return rule
		}
	}
	// Default Rule
	return ratelimit.Rule{
		Capacity: 10,
		Rate:     1,
		Path:     path,
		Method:   method,
	}
}
