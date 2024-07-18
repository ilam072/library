package books

import (
	"errors"
	"github.com/gin-gonic/gin"
	"islamic-library/internal/api/auth"
	"islamic-library/internal/model"
	"islamic-library/internal/myerrors"
	"islamic-library/pkg/response"
	"net/http"
)

func (api *API) Get(c *gin.Context) {
	username, err := auth.GetUsername(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		api.Logger.Error(err)
		return
	}

	var books []model.GetBookDTO

	URIUsername := c.Param("username")
	if _, err := api.userService.Get(URIUsername); errors.Is(err, myerrors.ErrInvalidUserData) {
		c.JSON(http.StatusNotFound, response.Error("user not found"))
		return
	}
	if username != URIUsername {
		books, err = api.bookService.GetPublic(URIUsername)
	} else {
		books, err = api.bookService.GetAll(URIUsername)
	}
	if err != nil {
		if errors.Is(err, myerrors.ErrBooksNotFound) {
			c.JSON(http.StatusNotFound, "books not found")
			api.Logger.Error(err)
			return
		}
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		api.Logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, books)

}
