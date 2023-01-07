package Interactive_Message

import (
	"os"
	"io"
	"log"
	"net/http"

	"gorm.io/gorm"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

const (
	buttonPushedAction := "buttonPushedAction"
	submitRestDateModal := "submitRestDateModal"
	confirmRestDateModal := "confirmRestDateModal"
	checkTransferCount := "checkTransferCount"
	studentOperation := "studentOperation"
)

func InteractiveHandler(w http.ResponseWriter, r *http.Request, api *slack.Client, db *gorm.DB)  {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("[ERROR] Failed to read Request body: %v",err)
		w.WriteHeader(200)
		return
	}
}
