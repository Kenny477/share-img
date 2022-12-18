package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/Kenny477/share-img/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		fileId := uuid.New()
		file.Filename = fileId.String() + "_" + file.Filename

		// Upload the file to specific dst.
		dst := filepath.Base(file.Filename)
		c.SaveUploadedFile(file, "storage/"+dst)
		db.SaveFile(file, fileId)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func Retrieve(c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)

	if err != nil {
		c.String(http.StatusNotFound, id+" invalid")
		return
	}

	file, dbErr := db.GetFile(uuid)
	if dbErr != nil {
		c.String(http.StatusNotFound, id+" not found")
		return
	}
	fmt.Println(file)

	// c.String(http.StatusOK, id)
}
