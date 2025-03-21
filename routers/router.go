package routers

import (
	"github.com/gin-gonic/gin"
	"visiontest/controllers"
	"visiontest/infrastructure/databasehelper"
	"visiontest/infrastructure/middleware"
)

func InitRouter(helper *databasehelper.DatabaseHelper) *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("DB", helper.DB)
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
