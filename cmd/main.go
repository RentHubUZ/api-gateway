package main

import (
	"api_gateway/api"
	"api_gateway/api/handler"
	"api_gateway/internal/config"
	logger "api_gateway/internal/logs"
	"api_gateway/internal/service"
	"log/slog"
	"os"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	logs := logger.NewLogger()

	path, err := os.Getwd()
	if err != nil {
		logs.Error("Failed to get working directory: ", slog.Any("error", err))
	}

	casbinEnforcer, err := casbin.NewEnforcer(path+"/internal/casbin/model.conf", path+"/internal/casbin/policy.csv")
	if err != nil {
		logs.Error("Failed to load Casbin enforcer: ", slog.Any("error", err))
	}

	services, err := service.NewServiceManager()
	if err != nil {
		logs.Error("Failed to initialize services: ", slog.Any("error", err))
	}

	handler := handler.NewHandler(services, logs)

	controller := api.NewController(gin.Default())
	controller.SetupRoutes(handler, logs, casbinEnforcer)

	if err := controller.StartServer(cfg); err != nil {
		logs.Error("Failed to start server: ", slog.Any("error", err))
	}
}
