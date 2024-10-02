package api

import (
	"api_gateway/api/handler"
	"api_gateway/config"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	SetupRoutes(handler.Handler, *slog.Logger)
	StartServer(config.Config) error
}

type controllerImpl struct {
	Port   string
	Router *gin.Engine
}

func NewController(router *gin.Engine) Controller {
	return &controllerImpl{Router: router}
}

func (c *controllerImpl) StartServer(cfg config.Config) error {
	c.Port = cfg.API_GATEWAY
	return c.Router.Run(c.Port)
}

func (c *controllerImpl) SetupRoutes(h handler.Handler, logger *slog.Logger) {
	router := c.Router.Group("/api")

	properties := router.Group("/properties")
	{
		properties.POST("/propertiescreate",h.CreateHouse)
		properties.PUT("/propertiesupdate",h.UpdateHouse)
		properties.GET("/propertiesgetall/:limit/:page",h.GetAllHouse)
		properties.GET("/propertiesgetbyid/:id",h.GetByIdHouse)
		properties.DELETE("/propertiesdelete/:id",h.DeleteHouse)
	}
}