package users

import (
	"github.com/sirupsen/logrus"
	"islamic-library/internal/model"
	"islamic-library/pkg/util"
)

type userRepo interface {
	Create(user model.User) error
	Get(username string) (model.User, error)
	Update(username string, input model.User) error
	Delete(username string) error
}

type Service struct {
	userRepo
	*logrus.Logger
}

func New(repo userRepo, log *logrus.Logger) *Service {
	return &Service{
		repo,
		log,
	}
}

func (s *Service) Create(user model.RegisterUserDTO) error {
	s.Logger.Debugln("hashing user password")
	hashPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	s.Logger.Debugln("user password was successfully hashed")

	return s.userRepo.Create(model.User{
		Username:     user.Username,
		PasswordHash: hashPassword,
		Email:        user.Email,
	})
}

func (s *Service) Get(username string) (model.GetUserDTO, error) {
	user, err := s.userRepo.Get(username)
	if err != nil {
		return model.GetUserDTO{}, err
	}

	return model.GetUserDTO{
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s *Service) Update(username string, input model.UpdateUserDTO) error {
	s.Logger.Debugln("hashing user password")
	hashPassword, err := util.HashPassword(input.Password)
	if err != nil {
		return err
	}
	s.Logger.Debugln("user password was successfully hashed")

	return s.userRepo.Update(username, model.User{
		Username:     input.Username,
		PasswordHash: hashPassword,
		Email:        input.Email,
	})
}

func (s *Service) Delete(username string) error {
	return s.userRepo.Delete(username)
}
