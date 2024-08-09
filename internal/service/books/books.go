package books

import (
	"github.com/sirupsen/logrus"
	"islamic-library/internal/model"
	"islamic-library/internal/myerrors"
	"net/http"
	"os"
	"path/filepath"
)

const (
	booksPath     = "books"
	privateAccess = "private"
	publicAccess  = "public"
)

type bookRepo interface {
	Create(book model.Book) error
	GetAll(username string) ([]model.Book, error)
	GetPublic(username string) ([]model.Book, error)
	GetById(bookId int) (model.Book, error)
	Update(bookId int, username string, input model.Book) error
	Delete(bookId int, username string) error
}

type userRepo interface {
	GetUserId(username string) (int, error)
}

type Service struct {
	userRepo
	bookRepo
	*logrus.Logger
}

func New(userRepo userRepo, bookRepo bookRepo, log *logrus.Logger) *Service {
	return &Service{
		userRepo,
		bookRepo,
		log,
	}
}

func (s *Service) Create(book model.CreateBookDTO, username string) error {
	if book.Access != privateAccess && book.Access != publicAccess {
		return myerrors.ErrInvalidBookAccess
	}

	userId, err := s.userRepo.GetUserId(username)
	if err != nil {
		return err
	}

	return s.bookRepo.Create(model.Book{
		Title:     book.Title,
		Author:    book.Author,
		IssueYear: book.IssueYear,
		UserId:    userId,
		FileName:  book.FileName,
		Access:    book.Access,
	})
}

func (s *Service) GetAll(username string) ([]model.GetBookDTO, error) {
	books, err := s.bookRepo.GetAll(username)
	if err != nil {
		return nil, err
	}

	responseBooks := make([]model.GetBookDTO, len(books))
	for i, book := range books {
		responseBooks[i] = model.GetBookDTO{
			Title:     book.Title,
			Author:    book.Author,
			IssueYear: book.IssueYear,
		}
	}
	return responseBooks, nil
}

func (s *Service) GetPublic(username string) ([]model.GetBookDTO, error) {
	books, err := s.bookRepo.GetPublic(username)
	if err != nil {
		return nil, err
	}

	responseBooks := make([]model.GetBookDTO, len(books))
	for i, book := range books {
		responseBooks[i] = model.GetBookDTO{
			Title:     book.Title,
			Author:    book.Author,
			IssueYear: book.IssueYear,
		}
	}
	return responseBooks, nil
}

func (s *Service) GetById(bookId int, username string) (model.DownloadBookDTO, error) {
	book, err := s.bookRepo.GetById(bookId)
	if err != nil {
		return model.DownloadBookDTO{}, err
	}

	userId, err := s.userRepo.GetUserId(username)
	if err != nil {
		return model.DownloadBookDTO{}, err
	}

	if userId != book.UserId && book.Access == privateAccess {
		return model.DownloadBookDTO{}, myerrors.ErrBookNotFound
	}

	return model.DownloadBookDTO{
		Title:     book.Title,
		Author:    book.Author,
		IssueYear: book.IssueYear,
		FileName:  book.FileName,
		Access:    book.Access,
	}, nil
}

func (s *Service) Update(bookId int, username string, input model.UpdateBookDTO) error {
	if err := s.bookRepo.Update(
		bookId,
		username,
		model.Book{
			Title:     input.Title,
			Author:    input.Author,
			IssueYear: input.IssueYear,
			Access:    input.Access,
		}); err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(bookId int, username string) error {
	return s.bookRepo.Delete(bookId, username)
}

func (s *Service) Download(fileName string) (string, []byte, error) {
	filePath := booksPath + "/" + filepath.Base(fileName)
	b, err := os.ReadFile(filePath)
	if err != nil {
		//pathError := err.(*fs.PathError)
		//if pathError.Err == syscall.ERROR_FILE_NOT_FOUND {
		//	return "", nil, myerrors.ErrBookFileNotFound
		//}
		return "", nil, err
	}
	contentType := http.DetectContentType(b[:512])

	return contentType, b, nil
}
