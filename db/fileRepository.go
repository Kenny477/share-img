package db

import (
	"mime/multipart"

	"github.com/Kenny477/share-img/models"
	"github.com/google/uuid"
)

func SaveFile(file *multipart.FileHeader, fileId uuid.UUID) {
	fileRecord := models.File{
		FileId:  fileId,
		Name:    file.Filename,
		Caption: "",
	}

	DB.Create(&fileRecord)
}

func GetFile(id uuid.UUID) (models.File, error) {
	var file models.File
	res := DB.First(&file, "file_id = ?", &id)

	if res.Error != nil {
		return models.File{}, res.Error
	}

	return file, nil
}
