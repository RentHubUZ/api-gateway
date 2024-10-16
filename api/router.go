package api

import (
	"api_gateway/api/handler"
	"api_gateway/api/middleware"
	"api_gateway/internal/config"
	"log/slog"

	_ "api_gateway/api/docs"

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

// @title Api Gateway
// @version 1.0
// @description This is a sample server for Api-gateway Service
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
// @schemes http
func (c *controllerImpl) SetupRoutes(h *handler.Handler, logger *slog.Logger, casbinEnforcer *casbin.Enforcer) {
	c.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	c.Router.Use(middleware.CORSMiddleware())
	c.Router.Use(middleware.PermissionMiddleware(casbinEnforcer))

	properties := c.Router.Group("/properties")
	{
		properties.POST("/propertiescreate", h.CreateHouse)
		properties.PUT("/propertiesupdate", h.UpdateHouse)
		properties.GET("/propertiesgetall/:limit/:page", h.GetAllHouse)
		properties.GET("/propertiesgetbyid/:properties_id", h.GetByIdHouse)
		properties.DELETE("/propertiesdelete/:properties_id", h.DeleteHouse)
	}

	payment := c.Router.Group("/payment")
	{
		payment.POST("/createpayment", h.CreatePayment)
		payment.GET("/getbyidpayment/:payment_id", h.GetPayment)
		payment.GET("/getallpayment/:limit/:page", h.GetAllPaymet)
		payment.DELETE("/deletepayment/:payment_id", h.DeletePayment)
	}

	tarif := c.Router.Group("/tarif")
	{
		tarif.POST("/createtarif", h.CreateTarif)
		tarif.PUT("/updatetarif", h.UpdateTarif)
		tarif.GET("/getbyidtarif/:tarif_id", h.GetByIdTarif)
		tarif.GET("/getalltarif/:limit/:page", h.GetAllTarif)
		tarif.DELETE("/deletetarif/:tarif_id", h.DeleteTarif)
	}

	topProperties := c.Router.Group("/topproperties")
	{
		topProperties.POST("/createtopproperties", h.CreateTopProperties)
		topProperties.PUT("/updatetopproperties", h.UpdateTopProperties)
		topProperties.GET("/getbyidtopproperties/:top_properties_id", h.GetByIdTopProperties)
		topProperties.GET("/getalltopproperties/:limit/:page", h.GetAllTopProperties)
		topProperties.DELETE("/deletetopproperties/:top_properties_id", h.DeleteTopProperties)
	}
	review := c.Router.Group("/review")
	{
		review.POST("/create", h.CreateReview)
		review.GET("/getall", h.GetAllReviews)
		review.GET("/getbyid/:id", h.GetReviewById)
		review.DELETE("/delete/:id", h.DeleteReview)
	}
	favorites := c.Router.Group("/favorites")
	{
		favorites.POST("/create", h.CreateFavorites)
		favorites.GET("/getall", h.GetAllFavorites)
		favorites.GET("/getbyid/:id", h.GetByIdFavorites)
		favorites.DELETE("/delete/:id", h.DeleteFavorites)
	}

	request := c.Router.Group("/request")
	{
		request.POST("/create", h.CreateRequest)
		request.GET("/getbyid/:id", h.GetRequestById)
		request.DELETE("/delete/:id", h.DeleteRequest)
	}

	report := c.Router.Group("/report")
	{
		report.POST("/create", h.CreateReport)
		report.GET("/getbyid/:id", h.GetReportById)
		report.DELETE("/delete/:id", h.DeleteReport)
	}

	notification := c.Router.Group("/notification")
	{
		notification.POST("/create", h.CreateNotification)
		notification.GET("/get/:id", h.GetNotification)
	}

	UploadMedia := c.Router.Group("/upload")
	{
		UploadMedia.POST("/imagesandvideo",h.UploadMedia)
	}
	
	user := c.Router.Group("/user")
	{
		user.GET("/profile",h.GetProfile)
		user.PUT("/profile/update",h.UpdateProfile)
		user.DELETE("/profile/delete",h.DeleteProfile)
		user.PUT("/password",h.ChangePassword)

	}
}
