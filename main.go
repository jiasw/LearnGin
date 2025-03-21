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
	db, err := databasehelper.NewDatabaseHelper(configger.Conf.DBconfig)
	if err != nil {
		fmt.Println("无法连接到数据库: ", err)
	}
	defer db.Close()
	// 自动迁移表结构
	db.DB.AutoMigrate(&models.UserInfo{})
	logger.Info(configger.Conf.Appname + "Starting server...")
	fmt.Println("Listen and serve on " + configger.Conf.Hostport)
	r := routers.InitRouter(db)
	r.Run(configger.Conf.Hostport)
}
