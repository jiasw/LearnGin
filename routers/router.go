package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"visiontest/controllers"
	"visiontest/infrastructure/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		fmt.Println("middleware")
		c.Next()
	})

	user := controllers.UserInfoController{

		Router: router,
	}
	router.Use(middleware.MiddlewareFunc())
	router.GET("/", controllers.Home)
	router.GET("/userinfo", user.GetUserInfo)
	router.POST("/userinfo", user.SaveUserInfo)
	router.DELETE("/userinfo", user.DeleteUserInfo)
	return router
}
