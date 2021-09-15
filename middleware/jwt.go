package middleware

import (
	"github.com/gin-gonic/gin"
	"myGinWeb/pkg/e"
	"myGinWeb/pkg/utils"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.SUCCESS
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := utils.ParseToken(token)

			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
