package Models

import (
	"time"
)

type TransferSchedule struct {
	ID uint64 `gorm:"ForiegnKey; AUTO_INCREMENT;"`
	Name string `gorm:"not null;"`
	Date time.Time `gorm:"not null;"`
	Time string `gorm:"not null;"`
	Subject string `gorm:"not null;"`
}

func (t *TransferSchedule) TableName() string {
	return "transferschedule"
}