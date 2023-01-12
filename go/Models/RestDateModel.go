package Models

import (
	"time"
)

type RestDate struct {
	ID uint64 `gorm:"primaryKey; ;AUTO_INCREMENT;"`
	StudentID uint64 
	Date time.Time `gorm:"not null;"`
	Subject string `gorm:"not null;"`
}

func (r *RestDate) TableName() string {
	return "restdate"
}