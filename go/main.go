package main

import (
	"net/http"
	"log"
	"os"

	"gorm.io/gorm"
    "gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/kabos0809/slack_bot/go/Models"
	"github.com/kabos0809/slack_bot/go/Config"
	"github.com/kabos0809/slack_bot/go/Mentioned_Message"
	"github.com/kabos0809/slack_bot/go/Middleware"
)

func main() {
	//mux := http.NewServeMux()
	godotenv.Load("../.env")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	dsn := Config.DbUrl()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	
	err = db.AutoMigrate(&Models.RestDate{}, &Models.Student{}, &Models.School{})
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/slack/mentioned", Middleware.Verify(func (w http.ResponseWriter, r *http.Request) {
		Mentioned_Message.MentionedHandler(w, r, api, db)
	}))
	/*http.HandleFunc("slack/Ineractive", Middleware.Verify(func (w http.ResponseWriter, r *http.Request) {
		Interactive_Message.InteractiveHandler(w, r, api, db)
	}))*/

	//middleware := Middleware.NewSecretsVerifierMiddleware(mux)

	log.Println("[INFO] Server listening")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
