package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Name    string
	Caption string
	FileId  uuid.UUID
}
