package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os"
	"student-manage/common"
)

func main() {
	InitConfig()
	_ = common.GetDB()

	router := gin.Default()
	_ = router.SetTrustedProxies([]string{"localhost"})
	router = collectRouter(router)

	// listen and serve on 0.0.0.0:port
	port := viper.GetString("server.port")
	if port != "" {
		if err := router.Run(":" + port); err != nil {
			log.Fatalln(err.Error())
		}
	}

	// listen and serve on 0.0.0.0:8080
	if err := router.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}

func InitConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalln("InitConfig() get error: ", err.Error())
	}

	viper.SetConfigName("applicationDev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalln("reading config get error: ", err.Error())
	}
}
