package routers

import (
	"visiontest/controllers"
	"visiontest/infrastructure/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.MiddlewareFunc())
	router.GET("/", controllers.Home)
	router.GET("/User", controllers.Usercontroller.GetUserInfo)
	router.POST("/CreateUser", controllers.Usercontroller.SaveUserInfo)

	return router
}
