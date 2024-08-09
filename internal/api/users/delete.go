package users

import (
	"github.com/gin-gonic/gin"
	"islamic-library/internal/api/auth"
	"islamic-library/pkg/response"
	"net/http"
)

func (api *API) Delete(c *gin.Context) {
	username, err := auth.GetUsername(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		api.Logger.Error(err)
		return
	}

	if err := api.service.Delete(username); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		api.Logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK())
}
