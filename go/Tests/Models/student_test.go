package test

import (
	"fmt"
	"testing"
)

/*type metadata struct {
	schid uuid.UUID
	stuid uuid.UUID
	resid uuid.UUID
}*/


func TestCreateStudent(t *testing.T) {
	for _, s := range testStudents_1 {
		if err := m.CreateStudent(&s); err != nil {
			t.Fatalf("[FAIL] Failed to Create Student: %s", err)
		}
		if err := m.AddStudent4School(&s, s.SchoolID); err != nil {
			t.Fatalf("[FAIL] Failed to Add Student for Schools Association: %s", err)
		}
	}
}

func TestGetStudent(t *testing.T) {
	students_1, err := m.GetStudentbySchoolAndGrade(testSchool_1.ID, "テスト")
	if err != nil {
		t.Fatalf("[FAIL] Failed to Get Student From School: %s", err)
	}
	students_2, err := m.GetStudentbySchoolAndGrade(testSchool_2.ID, "はずれ")
	if err != nil {
		t.Fatalf("[FAIL] Failed to Get Student From School: %s", err)
	}

	fmt.Println(students_1)
	fmt.Println(students_2)

	student, err := m.GetStudentbyID(testStudents_1[2].ID)
	if err != nil {
		t.Fatalf("[FAIL] Failed to Get Student by ID: %s", err)
	}

	fmt.Println(*student)

	transfer, err := m.TransferCount(testStudents_2[0].ID)
	if err != nil {
		t.Fatalf("[FAIL] Failed to Get Transfer Count: %s", err)
	}

	fmt.Println(transfer)
}

func TestDeleteStudent(t *testing.T) {
	defer sqldb.Close()

	for _, s := range testStudents_1 {
		if err := m.DeleteStudentFromSchool(&s, s.SchoolID); err != nil {
			t.Fatalf("[FAIL] Failed to Delete Student from School: %s", err)
		}
		if err := m.DeleteStudent(s.ID); err != nil {
			t.Fatalf("[FAIL] Failed to Delete Student: %s", err)
		}
	}
}