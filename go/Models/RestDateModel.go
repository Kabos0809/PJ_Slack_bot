package Models

import (
	"time"
	"github.com/google/uuid"
)

type RestDate struct {
	RestID uuid.UUID
	Name string `gorm:"not null;"`
	Date time.Time `gorm:"not null;"`
	Subject string `gorm:"not null;"`
}

func (r *RestDate) TableName() string {
	return "restdate"
}