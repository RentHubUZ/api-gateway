package service

import (
	"api_gateway/config"
	"api_gateway/genproto/accommodation"
	"api_gateway/genproto/favorites"
	"api_gateway/genproto/payment"
	"api_gateway/genproto/reviews"
	"api_gateway/genproto/tariff"
	"api_gateway/genproto/top-properties"
	"api_gateway/genproto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	UserService() user.UserClient
	PropertiesService() accommodation.AccommodationServiceClient
	FavoriteService() favorites.FavoritesClient
	PaymentService() payment.PaymentServiceClient
	ReviewService() reviews.ReviewsClient
	TarifService() tariff.TariffServiceClient
	TopPropertiesService() top_properties.TopPropertiesServiceClient
}

type serviceManagerImpl struct {
	userClient    user.UserClient
	propertiesClient accommodation.AccommodationServiceClient
	favoriteClient favorites.FavoritesClient
	paymentClient payment.PaymentServiceClient
	reviewClient reviews.ReviewsClient
	tarifClient tariff.TariffServiceClient
	topPropertiesClient top_properties.TopPropertiesServiceClient
}

func (s *serviceManagerImpl) UserService() user.UserClient {
	return s.userClient
}

func (s *serviceManagerImpl) PropertiesService() accommodation.AccommodationServiceClient {
	return s.propertiesClient
}

func (s *serviceManagerImpl) FavoriteService() favorites.FavoritesClient {
	return s.favoriteClient
}

func (s *serviceManagerImpl) PaymentService() payment.PaymentServiceClient {
	return s.paymentClient
}

func (s *serviceManagerImpl) ReviewService() reviews.ReviewsClient {
	return s.reviewClient
} 

func (s *serviceManagerImpl) TarifService() tariff.TariffServiceClient {
	return s.tarifClient
}

func (s *serviceManagerImpl) TopPropertiesService() top_properties.TopPropertiesServiceClient {
	return s.topPropertiesClient
}


func NewServiceManager() (ServiceManager, error) {
	connUser, err := grpc.Dial(
		config.Load().USER_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connAccom, err := grpc.Dial(
		config.Load().ACCOMMODATION_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connAction, err := grpc.Dial(
		config.Load().ACTION_BOARD,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &serviceManagerImpl{
		// userservice
		userClient:user.NewUserClient(connUser),

		// accommodationservice
		propertiesClient: accommodation.NewAccommodationServiceClient(connAccom),
		paymentClient: payment.NewPaymentServiceClient(connAccom),
		tarifClient: tariff.NewTariffServiceClient(connAccom),
		topPropertiesClient: top_properties.NewTopPropertiesServiceClient(connAccom),

		// actioboardService
		reviewClient: reviews.NewReviewsClient(connAction),
		favoriteClient: favorites.NewFavoritesClient(connAction),
	}, nil
}