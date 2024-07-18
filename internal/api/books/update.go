package books

import (
	"errors"
	"github.com/gin-gonic/gin"
	"islamic-library/internal/api/auth"
	"islamic-library/internal/model"
	"islamic-library/internal/myerrors"
	"islamic-library/pkg/response"
	"net/http"
	"strconv"
)

//books/2

func (api *API) Update(c *gin.Context) {
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

	var input model.UpdateBookDTO
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("invalid request body"))
		api.Logger.Error(err)
		return
	}

	if err := api.bookService.Update(bookId, username, input); err != nil {
		if errors.Is(err, myerrors.ErrBookNotFound) {
			c.JSON(http.StatusNotFound, "book id not found")
			api.Logger.Error(err)
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		api.Logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK())
}
