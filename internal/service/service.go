package service

import (
	"api_gateway/genproto/accommodation"
	"api_gateway/genproto/favorites"
	"api_gateway/genproto/notification"
	"api_gateway/genproto/payment"
	"api_gateway/genproto/report"
	"api_gateway/genproto/request"
	"api_gateway/genproto/reviews"
	"api_gateway/genproto/tariff"
	"api_gateway/genproto/top_properties"
	"api_gateway/genproto/user"
	"api_gateway/internal/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	UserService() user.UserClient
	PropertiesService() accommodation.AccommodationServiceClient
	FavouriteService() favorites.FavoritesClient
	PaymentService() payment.PaymentServiceClient
	ReviewService() reviews.ReviewsClient
	TarifService() tariff.TariffServiceClient
	TopPropertiesService() top_properties.TopPropertiesServiceClient
	NotificationService() notification.NotificationServiceClient
	ReportService() report.ReportServiceClient
	RequestService() request.RequestServiceClient
}

type serviceManagerImpl struct {
	userClient          user.UserClient
	propertiesClient    accommodation.AccommodationServiceClient
	favoriteClient      favorites.FavoritesClient
	paymentClient       payment.PaymentServiceClient
	reviewClient        reviews.ReviewsClient
	tarifClient         tariff.TariffServiceClient
	topPropertiesClient top_properties.TopPropertiesServiceClient
	notificationClient  notification.NotificationServiceClient
	reportClient        report.ReportServiceClient
	requestClient       request.RequestServiceClient
}

func (s *serviceManagerImpl) UserService() user.UserClient {
	return s.userClient
}

func (s *serviceManagerImpl) PropertiesService() accommodation.AccommodationServiceClient {
	return s.propertiesClient
}

func (s *serviceManagerImpl) FavouriteService() favorites.FavoritesClient {
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

func (s *serviceManagerImpl) NotificationService() notification.NotificationServiceClient {
	return s.notificationClient
}

func (s *serviceManagerImpl) ReportService() report.ReportServiceClient {
	return s.reportClient
}

func (s *serviceManagerImpl) RequestService() request.RequestServiceClient {
	return s.requestClient
}

func NewServiceManager() (ServiceManager, error) {
	connUser, err := grpc.NewClient(
		config.Load().USER_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connAccom, err := grpc.NewClient(
		config.Load().ACCOMMODATION_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connAction, err := grpc.NewClient(
		config.Load().ACTION_BOARD,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &serviceManagerImpl{
		// userservice
		userClient: user.NewUserClient(connUser),

		// accommodationservice
		propertiesClient:    accommodation.NewAccommodationServiceClient(connAccom),
		paymentClient:       payment.NewPaymentServiceClient(connAccom),
		tarifClient:         tariff.NewTariffServiceClient(connAccom),
		topPropertiesClient: top_properties.NewTopPropertiesServiceClient(connAccom),

		// actioboardService
		reviewClient:       reviews.NewReviewsClient(connAction),
		favoriteClient:     favorites.NewFavoritesClient(connAction),
		reportClient:       report.NewReportServiceClient(connAction),
		requestClient:      request.NewRequestServiceClient(connAction),
		notificationClient: notification.NewNotificationServiceClient(connAction),
	}, nil
}
