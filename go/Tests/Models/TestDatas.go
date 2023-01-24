package test

import (
	"time"

	"github.com/google/uuid"
	"github.com/kabos0809/slack_bot/go/Models"
)

var format string = "2006-01-02"
var testdate string = "20230117"

var date, _ = time.Parse(testdate, format)

var testSchool_1 = Models.School{
	ID: uuid.New(),
	Name: "Test_1",
}

var testSchool_2 = Models.School{
	ID: uuid.New(),
	Name: "Test_2",
}

var testStudents_1 = []Models.Student{
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

var testStudents_2 =[]Models.Student{
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

var testRestDates = []Models.RestDate{
	{
		ID: uuid.New(),
		StudentID: testStudents_2[0].ID,
		Subject: "国語",
		Date: date,
	},
	{
		ID: uuid.New(),
		StudentID: testStudents_2[0].ID,
		Subject: "数学",
		Date: date,
	},
	{
		ID: uuid.New(),
		StudentID: testStudents_2[0].ID,
		Subject: "数学",
		Date: date,
	},
	{
		ID: uuid.New(),
		StudentID: testStudents_2[2].ID,
		Subject: "英語",
		Date: date,
	},
}