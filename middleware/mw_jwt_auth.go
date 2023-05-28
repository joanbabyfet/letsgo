package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joanbabyfet/letsgo/common"
	"github.com/joanbabyfet/letsgo/service"
)

// jwt认证中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// handle before request

		auth := c.GetHeader("Authorization")
		token := strings.Split(auth, "Bearer ")[1]

		//验证token
		_, valid := service.Parse(token)

		if !valid {
			common.Error(c, "invalid token", -1, nil)
			return
		}
		c.Next()

		// handle after request
	}
}
