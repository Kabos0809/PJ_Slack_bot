package Models

import (
	"time"
)

type RestDate struct {
	ID uint64 `gorm: AUTO_INCREMENT;`
	Name string `gorm:"not null;"`
	Date time.Time `gorm:"not null;"`
	Subject string `gorm:"not null;"`
}

func (r *RestDate) TableName() string {
	return "restdate"
}