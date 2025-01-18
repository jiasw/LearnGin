package main

import (
	"fmt"
	"os"
	"visiontest/infrastructure/configger"
	"visiontest/infrastructure/logger"
	"visiontest/models"
	"visiontest/routers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	file, _ := os.Create("gin.log")
	defer file.Close()

	dsn := "root:123456@tcp(localhost:3306)/hdk_db?charset=utf8mb4&parseTime=True&loc=Local"
	// 连接到数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("无法连接到数据库: ", err)
	}
	db.AutoMigrate(&models.UserInfo{})

	logger.Info(configger.Conf.Appname + "Starting server...")
	port := configger.Conf.Hostport
	fmt.Println("Listen and serve on " + port)
	r := routers.InitRouter(db)
	r.Run(port)
}
