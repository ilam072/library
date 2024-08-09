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

func (api *API) Create(c *gin.Context) {
	username, err := auth.GetUsername(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		api.Logger.Error(err)
		return
	}

	var input model.CreateBookDTO
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("invalid request body"))
		api.Logger.Error(err)
		return
	}

	if err := api.bookService.Create(input, username); err != nil {
		if errors.Is(err, myerrors.ErrInvalidBookAccess) {
			c.JSON(http.StatusBadRequest, response.Error("set valid book access: private or public"))
			api.Logger.Error(err)
			return
		}
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		api.Logger.Error(err)
		return
	}
	c.JSON(http.StatusOK, response.OK())
}
