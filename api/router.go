package api

import (
	"api_gateway/api/handler"
	"api_gateway/api/middleware"
	"api_gateway/internal/config"
	"log/slog"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controller interface {
	SetupRoutes(*handler.Handler, *slog.Logger, *casbin.Enforcer)
	StartServer(*config.Config) error
}

type controllerImpl struct {
	Port   string
	Router *gin.Engine
}

func NewController(router *gin.Engine) Controller {
	return &controllerImpl{Router: router}
}

func (c *controllerImpl) StartServer(cfg *config.Config) error {
	c.Port = cfg.API_GATEWAY
	return c.Router.Run(c.Port)
}

func (c *controllerImpl) SetupRoutes(h *handler.Handler, logger *slog.Logger, enf *casbin.Enforcer) {
	c.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	c.Router.Use(middleware.CORSMiddleware())
	c.Router.Use(middleware.PermissionMiddleware(enf))

	router := c.Router.Group("/api")

	properties := router.Group("/properties")
	{
		properties.POST("/propertiescreate", h.CreateHouse)
		properties.PUT("/propertiesupdate", h.UpdateHouse)
		properties.GET("/propertiesgetall/:limit/:page", h.GetAllHouse)
		properties.GET("/propertiesgetbyid/:id", h.GetByIdHouse)
		properties.DELETE("/propertiesdelete/:id", h.DeleteHouse)
	}
}
