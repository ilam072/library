package users

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"islamic-library/internal/api/auth"
	"islamic-library/internal/model"
	"islamic-library/internal/myerrors"
	"islamic-library/pkg/response"
	"net/http"
)

func (api *API) Update(c *gin.Context) {
	username, err := auth.GetUsername(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		api.Logger.Error(err)
		return
	}

	var input model.UpdateUserDTO

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("invalid request body"))
		api.Logger.Error(err)
		return
	}

	if err := validator.New().Struct(&input); err != nil {
		validateErr := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, response.ValidationError(validateErr))
		api.Logger.Error(err)
		return
	}

	if err := api.service.Update(username, input); err != nil {
		if errors.Is(err, myerrors.ErrUsernameExists) {
			c.JSON(http.StatusConflict, response.Error("username already exists"))
			api.Logger.Error(err)
			return
		}

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		api.Logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK())
}
