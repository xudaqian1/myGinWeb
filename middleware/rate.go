package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"myGinWeb/pkg/utils"
	"net/http"
	"time"
)

func RateMiddleware(limiter *utils.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果ip请求连接数在两秒内超过5次，返回429并抛出error
		if !limiter.Allow(c.ClientIP(), 5, 2*time.Second) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			log.Println("too many requests")
			return
		}
		c.Next()
	}
}
