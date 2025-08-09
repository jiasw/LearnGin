package middleware

import (
	"net/http"
	"strings"

	jwthelper "visiontest/infrastructure/jwtHelper"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 权限认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "未提供token",
			})
			c.Abort()
			return
		}

		// 提取Token（格式: Bearer token）
		tokenParts := strings.SplitN(authHeader, " ", 2)
		if !(len(tokenParts) == 2 && strings.EqualFold(tokenParts[0], "Bearer")) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token格式错误",
			})
			c.Abort()
			return
		}

		// 验证Token
		claims, err := jwthelper.ParseToken(tokenParts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token无效或已过期",
			})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}
