package routers

import (
	"fmt"
	"visiontest/controllers"
	"visiontest/infrastructure/middleware"

	"github.com/gin-gonic/gin"
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
	router.Use(middleware.AuditMiddleware())
	api := router.Group("/api/v1")
	api.GET("/", controllers.Home)
	api.POST("/login", controllers.Login)
	api.Use(middleware.AuthMiddleware())
	{

		api.GET("/userlist", user.GetUserInfoList)
		api.GET("/userinfo", user.GetUserInfoByID)
		api.POST("/createUser", user.CreateUserInfo)
		api.POST("/delUser", user.DeleteUserByID)
		api.POST("/updateUser", user.UpdateUserInfo)
	}

	return router
}
