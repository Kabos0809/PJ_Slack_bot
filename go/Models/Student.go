package Models

import "github.com/google/uuid"

func (m Model) GetStudentbySchoolAndGrade(schoolID uuid.UUID, grade string) (*[]Student, error) {
	var students []Student
	var school *School
	tx := m.Db.Preload("Students").Begin()
	if err := tx.Where("id = ?", schoolID).First(&school).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	for _, s := range school.Students {
		if s.Grade == grade {
			students = append(students, s)
		}
	}

	return &students, nil
}

//IDから生徒情報取得
func (m Model) GetStudentbyID(id uuid.UUID) (*Student, error) {
	var student *Student
	tx := m.Db.Preload("RestDates").Begin()
	if err := tx.Where("id = ?", id).First(&student).Error; err != nil {
		tx.Rollback()
		return student, err
	}
	tx.Commit()
	return student, nil
}

//生徒情報の登録
func (m Model) CreateStudent(student *Student) error {
	student.Name = student.LastName + student.FirstName
	tx := m.Db.Begin()
	if err := tx.Create(student).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

//生徒情報の削除
func (m Model) DeleteStudent(id uuid.UUID) error {
	tx := m.Db.Preload("RestDates").Begin()
	if err := tx.Where("id = ?", id).Delete(&Student{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

//残り振替回数の追加
func IncrementCount(student *Student, sub string) *Student {
	switch sub {
	case "国語": student.JpnCounts = student.JpnCounts + 1
	case "数学": student.MathCounts = student.MathCounts + 1
	case "英語": student.EngCounts = student.EngCounts + 1
	}
	return student
}

//残り振替回数を減らす
func DecrementCount(student *Student, sub string) *Student {
	switch sub {
	case "国語": student.JpnCounts = student.JpnCounts - 1
	case "数学": student.MathCounts = student.MathCounts - 1
	case "英語": student.EngCounts = student.EngCounts - 1
	}
	return student
}

//休んだ日の追加
func (m Model) AddRestDate4Student(rdate *RestDate, id uuid.UUID) error {
	var student *Student
	tx := m.Db.Preload("RestDates").Begin()

	if err := tx.Where("id = ?", id).First(&student).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := m.Db.Model(student).Association("RestDates").Append(rdate); err != nil {
		tx.Rollback()
		return err
	}

	student = IncrementCount(student, rdate.Subject)

	if err := tx.Save(student).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

//休んだ日の削除
func (m Model) DeleteRestFromStudent(rdate *RestDate, id uuid.UUID) error {
	var student *Student
	tx := m.Db.Preload("RestDates").Begin()

	if err := tx.Where("id = ?", id).First(&student).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := m.Db.Model(student).Association("RestDates").Delete(rdate); err != nil {
		tx.Rollback()
		return err
	}

	student = DecrementCount(student, rdate.Subject)
	
	if err := tx.Save(student).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

//振替回数を返す
func (m Model) TransferCount(id uuid.UUID) (TransferCounts, error) {
	var student *Student

	tx := m.Db.Preload("RestDates").Begin()
	if err := tx.Where("id = ?", id).First(&student).Error; err != nil {
		tx.Rollback()
		return TransferCounts{}, err
	}
	tx.Commit()

	counts := TransferCounts{
		JpnCounts: student.JpnCounts,
		MathCounts: student.MathCounts,
		EngCounts: student.EngCounts,
	}
	
	return counts, nil
}