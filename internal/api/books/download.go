package books

import (
	"errors"
	"github.com/gin-gonic/gin"
	"islamic-library/internal/api/auth"
	"islamic-library/internal/myerrors"
	"islamic-library/pkg/response"
	"net/http"
	"strconv"
)

func (api *API) Download(c *gin.Context) {
	username, err := auth.GetUsername(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		api.Logger.Error(err)
		return
	}

	bookId, err := strconv.Atoi(c.Param("bookId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, myerrors.ErrInvalidBookId)
		api.Logger.Error(err)
		return
	}

	book, err := api.bookService.GetById(bookId, username)
	if err != nil {
		if errors.Is(err, myerrors.ErrBookNotFound) {
			c.JSON(http.StatusNotFound, response.Error("book not found"))
			api.Logger.Error(err)
			return
		}

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		api.Logger.Error(err)
		return
	}

	ct, bytes, err := api.bookService.Download(book.FileName)
	if err != nil {
		if errors.Is(err, myerrors.ErrBookFileNotFound) {
			c.JSON(http.StatusNotFound, response.Error("book is not uploaded"))
			api.Logger.Error(err)
			return
		}

		c.JSON(http.StatusNotFound, response.Error(err.Error()))
		api.Logger.Error(err)
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+book.Title)
	c.Data(http.StatusOK, ct, bytes)
}
