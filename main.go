package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"gin-gonic-gorm-boilerplate/configs"
	"gin-gonic-gorm-boilerplate/internal/database"
	"gin-gonic-gorm-boilerplate/internal/middleware"
	"gin-gonic-gorm-boilerplate/internal/routing"
	"gin-gonic-gorm-boilerplate/internal/util/logger"
	"gin-gonic-gorm-boilerplate/internal/util/parser"
	"gin-gonic-gorm-boilerplate/internal/ws"
)

func main() {
	// Loading Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		logger.Error(fmt.Sprintf("error reading config file, %s", err))
	}

	viper.AutomaticEnv()

	var config configs.Config
	parser.BindEnvs("", &config, "")

	if err := viper.Unmarshal(&config); err != nil {
		logger.Error(fmt.Sprintf("unable to decode into struct, %v", err))
	}

	if *parser.Replicas() != nil {
		config.DB.Replicas = *parser.Replicas()
	}

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

	// DB Init
	db := database.NewManager()
	if config.DB.Master.DBName != "" {
		db.Init(config.DB.Master, config.DB.Replicas)
		defer func(dbManager *database.Manager) {
			err := dbManager.Close()
			if err != nil {
				logger.Error("db close error")
				logger.Error(err)
			}
		}(db)
	}

	// Init WebSocket
	hub := ws.NewHub()
	go hub.Run()

	// Init Gin Engine
	gin.ForceConsoleColor()
	r := gin.Default()

	// Register Middleware
	r.Use(middleware.AddDbToContext(db))

	// Routing
	routing.Route(r, db, hub)

	go func() {
		if err := r.Run(fmt.Sprintf(":%d", config.Server.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	// Wait OS Signal
	quit := make(chan os.Signal)
	// kill (no param) default sends syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Warning("Shutting down server...")

	// Graceful shutdown
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logger.Warning("Server exiting")
}
