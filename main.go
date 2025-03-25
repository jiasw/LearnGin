package main

import (
	"fmt"
	"visiontest/infrastructure/configger"
	"visiontest/infrastructure/databasehelper"
	"visiontest/infrastructure/logger"
	"visiontest/models"
	"visiontest/routers"
)

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
	r.Run(configger.Conf.Hostport)
}
