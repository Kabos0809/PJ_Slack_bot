package Models

import (
	"time"
	"github.com/google/uuid"
)

type RestDate struct {
	ID uuid.UUID `gorm:"primaryKey; type:uuid;"`
	StudentID uuid.UUID
	Date time.Time `gorm:"not null;"`
	Subject string `gorm:"not null;"`
}

func (r *RestDate) TableName() string {
	return "restdate"
}