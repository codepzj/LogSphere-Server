package middleware

import (
	"server/models/common/response"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GetToken(c)
		if token == "" {
			response.NoAuth("未经授权访问", c)
			c.Abort()
			return
		}
		if _, err := utils.ParseToken(token); err != nil {
			response.NoAuth("令牌失效", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
