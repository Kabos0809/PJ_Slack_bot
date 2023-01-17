package Models

import (
	"time"
	"github.com/google/uuid"
)

type RestDate struct {
	ID uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4();"`
	StudentID uuid.UUID
	Date time.Time `gorm:"not null;"`
	Subject string `gorm:"not null;"`
}

func (r *RestDate) TableName() string {
	return "restdate"
}