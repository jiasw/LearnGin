package main

import (
	"fmt"
	"visiontest/infrastructure/configger"
	"visiontest/infrastructure/logger"
	"visiontest/routers"
)

func main() {
	logger.Info(configger.Conf.Appname + "Starting server...")
	port := configger.Conf.Hostport
	fmt.Println("Listen and serve on " + port)
	r := routers.InitRouter()
	r.Run(port)
}
