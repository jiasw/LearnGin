package main

import (
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "visiontest/docs" // 重要：导入生成的文档
	"visiontest/infrastructure/configger"
	"visiontest/infrastructure/databasehelper"
	"visiontest/infrastructure/logger"
	"visiontest/models"
	"visiontest/routers"
)

// @title 测试
// @version 1.0
// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io
// @Tags 		users
// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
func main() {
	// 连接到数据库
	err := databasehelper.Initialize(configger.Conf.DBconfig)
	if err != nil {
		fmt.Println("连接数据库失败: ", err)
	}
	db := databasehelper.GetInstance()
	defer db.Close()
	// 自动迁移表结构
	db.DB.AutoMigrate(&models.UserInfo{})
	logger.Info(configger.Conf.Appname + "Starting server...")
	fmt.Println("Listen and serve on " + configger.Conf.Hostport)
	r := routers.InitRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(configger.Conf.Hostport)
}
