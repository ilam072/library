package users

import (
	"github.com/gin-gonic/gin"
	"islamic-library/internal/api/auth"
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

	user, err := api.service.Get(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, user)
}
