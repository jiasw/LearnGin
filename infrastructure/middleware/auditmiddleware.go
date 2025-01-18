package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type AuditMiddleware struct {
}

func MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		fmt.Println("请求路径:", c.Request.URL.Path)
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}
