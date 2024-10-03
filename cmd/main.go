package main

import (
	"api_gateway/api"
	"api_gateway/api/handler"
	"api_gateway/internal/config"
	logger "api_gateway/internal/logs"
	"api_gateway/internal/service"
	"log"
	"os"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	logs := logger.NewLogger()
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	casbinEnforcer, err := casbin.NewEnforcer(path+"/internal/casbin/model.conf", path+"/internal/casbin/policy.csv")
	if err != nil {
		log.Fatal(err)
	}
	services, err := service.NewServiceManager()

	handler := handler.NewHandler(services, logs)

	controller := api.NewController(gin.Default())
	if err := controller.StartServer(cfg); err != nil {
		log.Fatal(err)
	}

	controller.SetupRoutes(handler, logs, casbinEnforcer)

}
