package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//IDから生徒情報取得
func GetStudentbyID(id uuid.UUID, db *gorm.DB) (*Student, error) {
	var student *Student
	tx := db.Preload("RestDates").Begin()
	if err := tx.Where("student_id = ?", id).First(student).Error; err != nil {
		tx.Rollback()
		return student, err
	}
	tx.Commit()
	return student, nil
}

//生徒情報の登録
func CreateStudent(student *Student, school *School, db *gorm.DB) error {
	tx := db.Begin()
	if err := tx.Create(student).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := db.Model(school).Association("Students").Append(student); err != nil {
		return err
	}
	tx.Commit()
	return nil
}

//生徒情報の削除
func DeleteStudent(student *Student, school *School, db *gorm.DB) error {
	tx := db.Preload("RestDate").Begin()
	if err := tx.Where("student_id = ?", student.StudentID).Delete(&Student{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

//残り振替回数の追加
func IncrementCount(student *Student, sub string, db *gorm.DB) error {
	tx := db.Begin()
	switch sub {
	case "国語": student.JpnCounts = student.JpnCounts + 1
	case "数学": student.MathCounts = student.MathCounts + 1
	case "英語": student.EngCounts = student.EngCounts + 1
	}
	if err := tx.Save(student).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

//残り振替回数を減らす
func DecrementCount(student *Student, sub string, db *gorm.DB) error {
	tx := db.Begin()
	switch sub {
	case "国語": student.JpnCounts = student.JpnCounts - 1
	case "数学": student.MathCounts = student.MathCounts - 1
	case "英語": student.EngCounts = student.EngCounts - 1
	}
	if err := tx.Save(student).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

//休んだ日の追加
func AddRestDate4Student(student *Student, rdate *RestDate, db *gorm.DB) error {
	tx := db.Begin()
	if err := db.Model(student).Association("RestDate").Append(rdate); err != nil {
		tx.Rollback()
		return err
	}
	if err := IncrementCount(student, rdate.Subject, db); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

//休んだ日の削除
func DeleteRestFromStudent(student *Student, rdate *RestDate, db *gorm.DB) error {
	tx := db.Begin()
	if err := db.Model(student).Association("RestDate").Delete(rdate); err != nil {
		tx.Rollback()
		return err
	}
	if err := DecrementCount(student, rdate.Subject, db); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

//振替回数を返す
func TransferCount(studentID uuid.UUID, db *gorm.DB) (*TransferCounts, error) {
	student, err := GetStudentbyID(studentID, db)
	if err != nil {
		return nil, err
	}
	counts := &TransferCounts{
		JpnCounts: student.JpnCounts,
		MathCounts: student.MathCounts,
		EngCounts: student.EngCounts,
	}
	return counts, nil
}