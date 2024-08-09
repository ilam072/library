package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"islamic-library/internal/model"
	"islamic-library/internal/myerrors"
	"islamic-library/pkg/util"
	"time"
)

type authorization interface {
	Get(username string) (model.User, error)
	Create(user model.User) error
}

const (
	tokenTTL   = time.Hour * 12
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
)

type tokenClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

type Service struct {
	authorization
	*logrus.Logger
}

func New(repo authorization, log *logrus.Logger) *Service {
	return &Service{
		repo,
		log,
	}
}

func (s *Service) GenerateToken(username, password string) (string, error) {
	user, err := s.authorization.Get(username)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", myerrors.ErrInvalidUserData
		}
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Username: username,
	})

	signedToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil

}
func (s *Service) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.Username, nil
}

func (s *Service) Create(user model.RegisterUserDTO) error {
	s.Logger.Debugln("hashing user password")
	hashPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	s.Logger.Debugln("user password was successfully hashed")

	return s.authorization.Create(model.User{
		Username:     user.Username,
		PasswordHash: hashPassword,
		Email:        user.Email,
	})
}
