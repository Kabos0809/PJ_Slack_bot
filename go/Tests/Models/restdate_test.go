package test

import (
	"log"
	"time"
	"testing"
	"database/sql"
	
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"github.com/google/uuid"
	"github.com/kabos0809/slack_bot/go/Config"
	"github.com/kabos0809/slack_bot/go/Models"
)
var format string = "2006-01-02"
var testdate string = "20230117"
type metadata struct {
	schid uuid.UUID
	stuid uuid.UUID
	resid uuid.UUID
}

var d = metadata{}
var m Models.Model
var sqldb *sql.DB

func TestCreate(t *testing.T) {
	m, sqldb = ConnDB()
	
	testSchool := Models.School{
		ID: uuid.New(),
		Name: "TestSchool1",
	}
	if err := m.AddSchool(&testSchool); err != nil {
		t.Fatalf("[FAIL] Failed to AddSchool: %s", err)
	}

	school, err := m.TestGetFirstSchool(testSchool.ID)
	if err != nil {
		t.Fatalf("[FAIL] Failed to Get School: %s", err)
	}

	testStudent := Models.Student{
		ID: uuid.New(),
		FirstName: "hoge",
		LastName: "henyo",
		Name: "hogehenyo",
		Grade: "テスト",
		SchoolID: school.ID,
	}
	
	if err := m.CreateStudent(&testStudent); err != nil {
	t.Fatalf("[FAIL] Failed to CreateStudent: %s", err)
	}


	
	if err := m.AddStudent4School(&testStudent, testStudent.SchoolID); err != nil {
	t.Fatalf("[FAIL] Failed to Add Student for School: %s", err)
	}


	date, _ := time.Parse(testdate, format)

	testRestDate := Models.RestDate{
		ID: uuid.New(),
		StudentID: testStudent.ID,
		Date: date,
		Subject: "英語",
	}

	
	if err := m.CreateRestDate(&testRestDate); err != nil {
	t.Fatalf("[FAIL] Failed to CreateRestDate: %s", err)
	}


	
	if err := m.AddRestDate4Student(&testRestDate, testRestDate.StudentID); err != nil {
	t.Fatalf("[FAIL] Failed to Add RestDate for Student: %s", err)
	}


	d = metadata{
		schid: testSchool.ID,
		stuid: testStudent.ID,
		resid: testRestDate.ID,
	}
}

func TestGetRestDate(t *testing.T) {
	restdate, err := m.GetRestDatebyID(d.resid)
	if err != nil {
		t.Fatalf("[FAIL] Failed to Get RestDate by ID: %s", err)
	}
	log.Printf("StudentID:%s, Date:%s, Subject:%s", restdate.StudentID.String(), restdate.Date.Format(format), restdate.Subject)
}

func TestDeleteRestDate(t *testing.T) {
	restdate, err := m.GetRestDatebyID(d.resid)
	if err != nil {
		t.Fatalf("[FAIL] Failed to Get RestDate by ID: %s", err)
	}

	
	if err := m.DeleteRestFromStudent(restdate, restdate.StudentID); err != nil {
	t.Fatalf("[FAIL] Failed to Delete RestDate from Student Association: %s", err)
	}

	time.Sleep(1 * time.Second)
	
	if err := m.DeleteRestDate(restdate.ID); err != nil {
	t.Fatalf("[FAIL] Failed to Delete RestDate: %s", err)
	}

}

func ConnDB() (Models.Model, *sql.DB) {
	godotenv.Load("../../.env")
	
	dsn := Config.DbUrl()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	
	err = db.AutoMigrate(&Models.RestDate{}, &Models.Student{}, &Models.School{})
	if err != nil {
		panic(err)
	}

	m := Models.Model{Db: db}
	return m, sqlDB
}