package handler

import (
	pbaccom "api_gateway/genproto/accommodation"
	pbfav "api_gateway/genproto/favorites"
	pbnotif "api_gateway/genproto/notification"
	pbpay "api_gateway/genproto/payment"
	pbreport "api_gateway/genproto/report"
	pbreq "api_gateway/genproto/request"
	pbrev "api_gateway/genproto/reviews"
	pbtarf "api_gateway/genproto/tariff"
	pbtop "api_gateway/genproto/top_properties"
	pbuser "api_gateway/genproto/user"
	"api_gateway/internal/service"
	"errors"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService          pbuser.UserClient
	FavoriteService      pbfav.FavoritesClient
	ReviewService        pbrev.ReviewsClient
	TarifService         pbtarf.TariffServiceClient
	PaymentService       pbpay.PaymentServiceClient
	AccommodationService pbaccom.AccommodationServiceClient
	TopPropertiesService pbtop.TopPropertiesServiceClient
	NotificationService  pbnotif.NotificationServiceClient
	ReportService        pbreport.ReportServiceClient
	RequestService       pbreq.RequestServiceClient
	Log                  *slog.Logger
}

func NewHandler(service service.ServiceManager, logs *slog.Logger) *Handler {
	return &Handler{
		UserService:          service.UserService(),
		FavoriteService:      service.FavoriteService(),
		ReviewService:        service.ReviewService(),
		TarifService:         service.TarifService(),
		PaymentService:       service.PaymentService(),
		AccommodationService: service.PropertiesService(),
		TopPropertiesService: service.TopPropertiesService(),
		NotificationService:  service.NotificationService(),
		ReportService:        service.ReportService(),
		RequestService:       service.RequestService(),
		Log:                  logs,
	}
}

func getUserID(c *gin.Context) (string, error) {
	id, ok := c.Get("user_id")
	if !ok {
		return "", errors.New("user id not found")
	}

	idStr, ok := id.(string)
	if !ok {
		return "", errors.New("invalid user id")
	}

	return idStr, nil
}
