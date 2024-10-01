package handler

import (
	pbaccom "api_gateway/genproto/accommodation"
	pbfav "api_gateway/genproto/favorites"
	pbpay "api_gateway/genproto/payment"
	pbrev "api_gateway/genproto/reviews"
	pbtarf "api_gateway/genproto/tariff"
	pbtop "api_gateway/genproto/top-properties"
	pbuser "api_gateway/genproto/user"
	"log/slog"
)

type Handler struct {
	UserService pbuser.UserClient
	FavoriteService pbfav.FavoritesClient
	ReviewService pbrev.ReviewsClient
	TarifService pbtarf.TariffServiceClient
	PaymentService pbpay.PaymentServiceClient
	AccommodationService pbaccom.AccommodationServiceClient
	TopPropertiesService pbtop.TopPropertiesServiceClient
	Log *slog.Logger
}

func NewHandler(UserService pbuser.UserClient, FavoriteService pbfav.FavoritesClient, ReviewService pbrev.ReviewsClient, TarifService pbtarf.TariffServiceClient, PaymentService pbpay.PaymentServiceClient, AccommodationService pbaccom.AccommodationServiceClient, TopPropertiesService pbtop.TopPropertiesServiceClient, logs *slog.Logger) *Handler {
	return &Handler{
		UserService: UserService,
		FavoriteService: FavoriteService,
		ReviewService: ReviewService,
		TarifService: TarifService,
		PaymentService: PaymentService,
		AccommodationService: AccommodationService,
		TopPropertiesService: TopPropertiesService,
		Log: logs,
	}
}