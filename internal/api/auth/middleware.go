package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"islamic-library/pkg/response"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "username"
)

func (api *API) UserIdentify(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		c.JSON(http.StatusUnauthorized, response.Error("empty auth header"))
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, response.Error("invalid auth header"))
		return
	}

	if len(headerParts[1]) == 0 {
		c.JSON(http.StatusUnauthorized, response.Error("empty token"))
		return
	}
	username, err := api.service.ParseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, username)
}

func GetUsername(c *gin.Context) (string, error) {
	username, exists := c.Get(userCtx)
	if !exists {
		return "", errors.New("username not found")
	}

	usernameString, ok := username.(string)
	if !ok {
		return "", errors.New("username is of invalid type")
	}

	return usernameString, nil
}
