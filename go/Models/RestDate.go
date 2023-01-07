package Models

import (
	"errors"

	"gorm.io/gorm"
)

//以下二つはWebページに対応した場合に使うかも
/*
func GetAllRestDate() (*[]RestDate, error) {
	var restdates []RestDate
	tx := db.Begin()
	if err := tx.Find(&restdates).Error; err != nil {
		tx.Rollback()
		return &restdates, err
	}
	tx.Commit()
	return &restdates, nil
}
*/

func GetRestDatebyID(id uint64, db *gorm.DB) (*RestDate, error) {
	var restdate *RestDate
	tx := db.Begin()
	if err := tx.Where("id = ?", id).First(restdate).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return restdate, nil
}

//休んだ日の登録
func CreateRestDate(rdate *RestDate, student *Student, db *gorm.DB) error {
	tx := db.Begin()
	if rdate.Subject != "国語" && rdate.Subject != "数学" && rdate.Subject != "英語" {
		err := errors.New("[ERROR] Subjects must be Japanese, Mathematics or English.")
		return err
	}
	if err := tx.Create(rdate).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := AddRestDate4Student(student, rdate, db); err != nil {
		return err
	}
	tx.Commit()
	return nil
}

//休んだ日の削除
func DeleteRestDate(id uint64, db *gorm.DB) error {
	tx := db.Begin()
	restdate, err := GetRestDatebyID(id, db)
	if err != nil {
		return err
	}

	student, err := GetStudentbyID(id, db)
	if err != nil {
		return err
	}

	if err := DeleteRestFromStudent(student, restdate, db); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("id = ?", id).Delete(&RestDate{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}