package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"islamic-library/internal/api/auth"
	"islamic-library/internal/api/books"
	"islamic-library/internal/api/users"
	"islamic-library/internal/model"
)

type authService interface {
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (string, error)
	Create(user model.RegisterUserDTO) error
}

type usersService interface {
	Create(user model.RegisterUserDTO) error
	Get(username string) (model.GetUserDTO, error)
	Update(username string, input model.UpdateUserDTO) error
	Delete(username string) error
}

type booksService interface {
	Create(book model.CreateBookDTO, username string) error
	GetAll(username string) ([]model.GetBookDTO, error)
	GetPublic(username string) ([]model.GetBookDTO, error)
	GetById(bookId int, username string) (model.DownloadBookDTO, error)
	Update(bookId int, username string, input model.UpdateBookDTO) error
	Delete(bookId int, username string) error
	Download(fileName string) (string, []byte, error)
}

func New(logger *logrus.Logger, authService authService, usersService usersService, booksService booksService) *gin.Engine {

	r := gin.New()
	r.Use(gin.Recovery())

	authAPI := auth.NewAPI(logger, authService)
	userAuth := r.Group("/user-auth")
	{
		userAuth.POST("/sign-up", authAPI.SignUp)
		userAuth.POST("/sign-in", authAPI.SignIn)
	}

	usersAPI := users.NewAPI(logger, usersService)
	api := r.Group("/api", authAPI.UserIdentify)
	{
		api.GET("/user", usersAPI.Get)
		api.POST("/user", usersAPI.Update)
		api.DELETE("/user", usersAPI.Delete)
	}

	booksAPI := books.NewAPI(logger, booksService, usersService)
	{
		api.GET("/:username/books", booksAPI.Get)
		api.GET("/books/:bookId", booksAPI.Download)
		api.POST("/book", booksAPI.Create)
		api.POST("/books/upload", booksAPI.Upload)
		api.POST("/books/:bookId", booksAPI.Update)
		api.DELETE("/books/:bookId", booksAPI.Delete)
	}

	return r
}
