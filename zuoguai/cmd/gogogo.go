package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"time"
	"zuoguai/internal/api"
	"zuoguai/internal/config"
	"zuoguai/internal/db"
	"zuoguai/internal/entity"
	"zuoguai/internal/service"
)

const DevMode = true
const DevConfigPath = "C:\\Users\\作怪\\Desktop\\app.yaml"

//func Init() {
//	gin.SetMode(gin.DebugMode)
//	if DEV_MODE {
//		gin.SetMode(gin.DebugMode)
//	} else {
//		gin.SetMode(gin.ReleaseMode)
//	}
//	gin.DefaultWriter = io.Discard
//}

func main() {

	var path, sep string
	for i := 1; i < len(os.Args); i++ {
		path += sep + os.Args[i]
		sep = " "
	}
	if path == "" {
		path = DevConfigPath
	}

	r := gin.Default()
	_ = r.SetTrustedProxies([]string{"127.0.0.1"})
	conf := config.GetConfigs(path)
	entity.GetEntities()
	service.GetServices()
	api.RegisterRouter(r)
	_ = db.GetDB().AutoMigrate(&entity.User{}, &entity.Schedule{})

	//printAppInfo(conf, 10)
	if err := r.Run(":" + strconv.Itoa(conf.App.ServerPort)); err != nil {
		fmt.Println("start server err:" + err.Error())
	}
}
func printAppInfo(conf *config.Configs, n int) {
	fmt.Println(fmt.Sprintf("PRINT X %d", n))
	for i := 0; i < n; i++ {
		fmt.Println(fmt.Sprintf("NAME: %s, VSERSION: %s, SERVER START AT PORT: %d, TIME: %s",
			conf.App.AppName, conf.App.AppVersion, conf.App.ServerPort, time.Now().Format("2006-01-02 15:04:05")))
	}
}
