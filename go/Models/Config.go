package Models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Model struct {
	Db *gorm.DB
}