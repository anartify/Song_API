package middleware

import (
	"Song_API/pkg/cache"
	"Song_API/pkg/ratelimit"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RateLimit function is a middleware function that checks if the rate limit is reached for a particular client. If the quota of client isn't reached, it verifies the global quota (Global limit is set to protect the server from all clients across all endpoints). If the global quota is reached, it rejects the request, otherwise it allows the request to proceed.
func RateLimit(rateRules []ratelimit.Rule, globalRule ratelimit.Rule, bucketCache cache.Cache) gin.HandlerFunc {
	return func(c *gin.Context) {
		client := getClient(c)
		rule := getRule(c, rateRules)
		bucket := ratelimit.GetBucket(client, rule, bucketCache)
		globalBucket := ratelimit.GetBucket("global", globalRule, bucketCache)
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
		bucketCache.Set(client, bucket)
		bucketCache.Set("global", globalBucket)
		c.Next()
	}
}

// getClient function returns a string that uniquely identifies a client. It is a combination of client IP, path and method.
func getClient(c *gin.Context) string {
	ip := c.ClientIP()
	path := c.Request.URL.Path
	method := c.Request.Method
	data := fmt.Sprintf("%s-%s-%s", ip, path, method)
	return data
}

// isPathMatched function checks if the path in the request matches the path in the rate rule.
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

// getRule function returns the rate rule for a particular request. If the request doesn't match any rule, it returns the default rule.
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
