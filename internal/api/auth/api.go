package auth

import (
	"github.com/sirupsen/logrus"
	"islamic-library/internal/model"
)

type service interface {
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (string, error)
	Create(user model.RegisterUserDTO) error
}

type API struct {
	service service
	*logrus.Logger
}

func NewAPI(logger *logrus.Logger, service service) *API {
	return &API{
		service: service,
		Logger:  logger,
	}
}
