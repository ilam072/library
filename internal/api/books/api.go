package books

import (
	"github.com/sirupsen/logrus"
	"islamic-library/internal/model"
)

type bookService interface {
	Create(book model.CreateBookDTO, username string) error
	GetAll(username string) ([]model.GetBookDTO, error)
	GetPublic(username string) ([]model.GetBookDTO, error)
	GetById(bookId int, username string) (model.DownloadBookDTO, error)
	Update(bookId int, username string, input model.UpdateBookDTO) error
	Delete(bookId int, username string) error
	Download(fileName string) (string, []byte, error)
}

type userService interface {
	Get(username string) (model.GetUserDTO, error)
}

type API struct {
	bookService bookService
	userService userService
	*logrus.Logger
}

func NewAPI(logger *logrus.Logger, bookService bookService, userService userService) *API {
	return &API{
		bookService: bookService,
		userService: userService,
		Logger:      logger,
	}
}
