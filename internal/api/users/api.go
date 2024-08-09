package users

import (
	"github.com/sirupsen/logrus"
	"islamic-library/internal/model"
)

// TODO: 09.07: handle "/api/get" [GET]; "/api/update" [POST]; "/api/delete" [DELETE]

type service interface {
	Create(user model.RegisterUserDTO) error
	Get(username string) (model.GetUserDTO, error)
	Update(username string, input model.UpdateUserDTO) error
	Delete(username string) error
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
