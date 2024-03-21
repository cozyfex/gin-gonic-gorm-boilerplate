package main

import (
	"fmt"
	"gin-gonic-gorm-boilerplate/configs"
	"gin-gonic-gorm-boilerplate/internal/db"
	"gin-gonic-gorm-boilerplate/internal/routing"
	"gin-gonic-gorm-boilerplate/internal/util/logger"
	"gin-gonic-gorm-boilerplate/internal/util/parser"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 설정 로딩
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		logger.Error(fmt.Sprintf("error reading config file, %s", err))
	}

	viper.AutomaticEnv()

	var config configs.Config
	if err := viper.Unmarshal(&config); err != nil {
		logger.Error(fmt.Sprintf("unable to decode into struct, %v", err))
	}

	if *parser.ReplicaParser() != nil {
		config.DB.Replicas = *parser.ReplicaParser()
	}

	logger.Warning(config.DB.Master.Port)
	logger.Warning(config.DB.Replicas[0].Port)

	switch config.Mode {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "stg":
		gin.SetMode(gin.DebugMode)
	case "prd":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	logger.Info(fmt.Sprintf("Config Mode: %s", config.Mode))
	logger.Info(fmt.Sprintf("Gin Mode: %s", gin.Mode()))

	// DB 초기화
	dbManager := db.NewManager()
	dbManager.Init(config.DB.Master, config.DB.Replicas)

	r := gin.Default()

	// Routing
	routing.Route(r)

	r.Run(fmt.Sprintf(":%d", config.Server.Port)) // listen and serve on 0.0.0.0:8080
}
