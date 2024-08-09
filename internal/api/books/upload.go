package books

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
)

const booksPath = "books"

func (api *API) Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	if filepath.Ext(file.Filename) != ".pdf" {
		c.JSON(http.StatusBadRequest, "file extension must be .pdf")
		return
	}
	fileUUID, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "oops, something went wrong")
		return
	}

	fileName := fileUUID.String() + filepath.Ext(file.Filename)
	file.Filename = fileName
	savePath := booksPath + "/" + filepath.Base(file.Filename)
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "file uploading failed")
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("'%s' uploaded!", fileName))
}
