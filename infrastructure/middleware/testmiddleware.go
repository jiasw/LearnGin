package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type TestMiddleware struct {
}

func (tm *TestMiddleware) Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("TestMiddleware")
		c.Next()
		fmt.Println("TestMiddleware end")
	}

}

func (tm *TestMiddleware) Test2(c *gin.Context) {
	fmt.Println("TestMiddleware2")
	c.Next()
}

func (tm *TestMiddleware) Test3(c *gin.Context) {
	fmt.Println("TestMiddleware3")
	c.Next()
}

func (tm *TestMiddleware) Test4(c *gin.Context) {
	fmt.Println("TestMiddleware4")
	c.Next()
}

func (tm *TestMiddleware) Test5(c *gin.Context) {
	fmt.Println("TestMiddleware5")
	c.Next()
}

func (tm *TestMiddleware) Test6(c *gin.Context) {
	fmt.Println("TestMiddleware6")
	c.Next()
}
