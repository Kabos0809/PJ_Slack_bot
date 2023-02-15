package Models

import (
	"gorm.io/gorm"
)

type Model struct {
	Db *gorm.DB
}