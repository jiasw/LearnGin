package routers

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"time"
	"visiontest/controllers"
	"visiontest/infrastructure/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // 开发环境可以使用 *
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-Requested-With",
		},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(func(c *gin.Context) {
		fmt.Println("全局中间件")
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
