package Models

import (
	"testing"
	"time"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/kabos0809/slack_bot/go/Config"
)

var layout string = "2000-01-01"

func Test_CD_RestDate(t *testing.T) {
	dsn := Config.DbUrl()
	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	
	err = db.AutoMigrate(&RestDate{}, &Student{}, &TransferSchedule{})
	if err != nil {
		panic(err)
	}

	date, _ := time.Parse(layout, "2022-12-08")
	testdate_ok := &RestDate{
		ID: 1,
		Name: "TEST USER",
		Date: date,
		Subject: "数学",
	}
	err = CreateRestDate(testdate_ok, db)
	if err != nil {
		t.Error("[ERROR] RestDate could not Create.")
	}
	err = DeleteRestDate(testdate_ok.ID, db)
	if err != nil {
		t.Error("[ERROR] RestDate could not Delete")
	}
	testdate_fail := &RestDate{
		ID:2,
		Name: "FAIL USER",
		Date: date,
		Subject: "保健",
	}
	err = CreateRestDate(testdate_fail, db)
	if err != nil {
		log.Println(err)
		t.Error("[ERROR] RestDate could not Create.")
	}
}


/*func GetRestDate_Test(t *testing.T) {
	date, _ := time.Parse(layout, "2022-12-08")
	testdate := &
}*/