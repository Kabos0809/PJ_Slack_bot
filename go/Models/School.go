package Models

import (
	"gorm.io/gorm"
)

func GetAllSchool(db *gorm.DB) (*[]School, error) {
	var schools []School
	tx := db.Begin()
	if err := tx.Find(&schools).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &schools, nil
}

func GetSchool(id uint64, db *gorm.DB) (*School, error) {
	var school *School
	tx := db.Begin()
	if err := tx.Where("id = ?", id).Find(school).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return school, nil
}

func AddSchool(school *School, db *gorm.DB) error {
	tx := db.Begin()
	if err := tx.Create(school).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func DeleteSchool(id uint64, db *gorm.DB) error {
	tx := db.Begin()
	if err := tx.Where("id = ?", id).Delete(&School{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}