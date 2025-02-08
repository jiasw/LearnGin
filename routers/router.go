package routers

import (
	"visiontest/controllers"
	"visiontest/infrastructure/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	user := controllers.UserInfoController{
		Db:     db,
		Router: router,
	}
	router.Use(middleware.MiddlewareFunc())

	router.GET("/", controllers.Home)
	router.GET("/userinfo", user.GetUserInfo)
	router.POST("/userinfo", user.SaveUserInfo)
	router.DELETE("/userinfo", user.DeleteUserInfo)
	return router
}
