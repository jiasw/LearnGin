package routers

import (
	"fmt"
	"reflect"
	"visiontest/controllers"
	"visiontest/infrastructure/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	mid := middleware.TestMiddleware{}
	mid2 := new(middleware.TestMiddleware)

	arr := [10]int{}
	arr[0] = 10
	fmt.Println("arr type:", reflect.TypeOf(arr))
	fmt.Println("mid type:", reflect.TypeOf(mid))
	fmt.Println("mid2 type:", reflect.TypeOf(mid2))

	router := gin.Default()
	userctl := controllers.UserInfoController{
		Db:     db,
		Router: router,
	}
	router.Use(middleware.MiddlewareFunc())
	router.Use(mid.Test())
	router.GET("/", controllers.Home)
	router.GET("/userinfo", userctl.GetUserInfo)
	router.POST("/userinfo", userctl.SaveUserInfo)
	router.DELETE("/userinfo", userctl.DeleteUserInfo)
	return router
}
