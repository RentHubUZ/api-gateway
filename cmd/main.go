package main

import (
	"api_gateway/api"
	"api_gateway/api/handler"
	"api_gateway/internal/config"
	"api_gateway/internal/pkg/logger"
	"api_gateway/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	logs := logger.NewLogger()
	config := config.Load()

	servicemanger, err := service.NewServiceManager()

	if err != nil {
		log.Println("Error initializing service manager", "error", err.Error())
		logs.Error("Error initializing service manager", "error", err.Error())
		return
	}

	handler := handler.NewHandler(servicemanger.UserService(),servicemanger.FavoriteService(),
							servicemanger.ReviewService(),servicemanger.TarifService(),servicemanger.PaymentService(),
							servicemanger.PropertiesService(),servicemanger.TopPropertiesService(),logs)
	controller := api.NewController(gin.Default())
	controller.SetupRoutes(*handler, logs)
	controller.StartServer(config)
}
