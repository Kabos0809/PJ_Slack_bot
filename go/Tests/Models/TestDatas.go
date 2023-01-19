package test

import (
	"github.com/google/uuid"
	"github.com/kabos0809/slack_bot/go/Models"
)

var testSchool_1 = Models.School{
	ID: uuid.New(),
	Name: "Test_1",
}

var testSchool_2 = Models.School{
	ID: uuid.New(),
	Name: "Test_2",
}

var testStudents = []Models.Student{
	{
		ID: uuid.New(),
		FirstName: "test",
		LastName: "1",
		Name: "test_1",
		Grade: "テスト",
		SchoolID: testSchool_1.ID,
	},
	{
		ID: uuid.New(),
		FirstName: "test",
		LastName: "2",
		Name: "test_2",
		Grade: "テスト",
		SchoolID: testSchool_1.ID,
	},
	{
		ID: uuid.New(),
		FirstName: "test",
		LastName: "3",
		Name: "test_3",
		Grade: "テスト",
		SchoolID: testSchool_2.ID,
	},
	{
		ID: uuid.New(),
		FirstName: "test",
		LastName: "4",
		Name: "test_4",
		Grade: "テスト",
		SchoolID: testSchool_2.ID,
	},
}