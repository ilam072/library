package books

import (
	"github.com/gin-gonic/gin"
	"islamic-library/internal/api/auth"
	"islamic-library/internal/myerrors"
	"islamic-library/pkg/response"
	"net/http"
	"strconv"
)

func (api *API) Delete(c *gin.Context) {
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

	if err := api.bookService.Delete(bookId, username); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		api.Logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK())
}
