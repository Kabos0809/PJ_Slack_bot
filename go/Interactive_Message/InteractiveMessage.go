package Interactive_Message

import (
	"os"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/kabos0809/slack_bot/go/Models"
)

var (
	buttonPushedAction = "buttonPushedAction"
	submitRestDateModal = "submitRestDateModal"
	confirmRestDateModal = "confirmRestDateModal"
	checkTransferCount = "checkTransferCount"
	studentOperation = "studentOperation"
	submitStudentModal = "submitStudentModal"
	confirmStudentModal = "confirmStudentModal"
	deleteStudent = "deleteStudent"
	updateStudent = "updateStudent"
	schoolOperation = "schoolOperation"
	submitSchoolModal = "submitSchoolModal"
	confirmSchoolModal = "confirmSchoolModal"
)

var fallbackText slack.MsgOption = slack.MsgOptionText("This client is not supported.", false)

//InteractiveMessageの送信
func InteractiveHandler(w http.ResponseWriter, r *http.Request, api *slack.Client, m Models.Model)  {
	var payload slack.InteractionCallback
	if err := json.Unmarshal([]byte(r.FormValue("payload")), &payload); err != nil {
		log.Printf("[ERROR]: Internal Server Error has occured: %s", err)
		w.WriteHeader(200)
		return
	}

	switch CheckCallbackType(payload) {
	case buttonPushedAction:
		if len(payload.ActionCallback.BlockActions) > 1 {
			school := payload.ActionCallback.BlockActions[0].Value
			grade := payload.ActionCallback.BlockActions[1].Value
			msg, err := StudentListHandle(payload, school, grade)
			if _, _, _, err := api.SendMessage("", payload.ResponseURL, fallbackText, msg); err != nil {
				log.Printf("[ERROR] Failed to handle school button push action: %s", err)
				w,WriteHeader(200)
				return
			}
		}
		switch payload.ActionCallback.BlockActions[0].Value {
		case "school":
			msg, err := SchoolButtonPushedActionHandle(payload)
			if err != nil {
				log.Printf("[ERROR] Failed to handle school button push action: %s", err)
				w.WriteHeader(200)
				return
			}
			if _, _, _, err := api.SendMessage("", payload.ResponseURL, fallbackText, msg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			return
		case "restdate":
			msg, err := RestDateButtonPushedActionHandle(payload)
			if err != nil {
				log.Printf("[ERROR] Failed to handle restdate button push action: %s", err)
				w.WriteHeader(200)
				return
			}
			if _, _, _, err := api.SendMessage("", payload.ResponseURL, fallbackText, msg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			return
		case "student":
			msg, err := StudentButtonPushedActionHandle(payload)
			if err != nil {
				log.Printf("[ERROR] Failed to handle student button push action: %s", err)
				w.WriteHeader(200)
				return
			}
			if _, _, _, err := api.SendMessage("", payload.ResponseURL, fallbackText, msg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
		case "TransferCount":
			msg, err := TransferCountPushedActionHandle(payload)
			if err != nil {
				log.Printf("[ERROR] Failed to handle student button push action: %s", err)
				w.WriteHeader(200)
				return
			}
			if _, _, _, err := api.SendMessage("", payload.ResponseURL, fallbackText, msg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
		}
	case submitRestDateModal:
		modal, err := SubmitRestDateHandle(payload)
		if err != nil {
			log.Printf("[ERROR] Failed to handle submit restdate: %s", err)
			w.WriteHeader(200)
			return
		}
		if _, err := api.OpenView(payload.TriggerID, *modal); err != nil {
			log.Printf("[ERROR] Failed to open modal: %s", err)
			w.WriteHeader(200)
			return
		}
		return
	case confirmRestDateModal:
		err := ConfirmRestDateHandle(payload, m)
		if err != nil {
			log.Printf("[ERROR] Failed to handle confirm rest date: %s", err)
			w.WriteHeader(200)
			return
		}
		return
	case submitSchoolModal:
		modal, err := SubmitSchoolHandle(payload)
		if err != nil {
			log.Printf("[ERROR] Failed to handle submit school handle: %s", err)
			w.WriteHeader(200)
			return
		}
		if _, err := api.OpenView(payload.TriggerID, *modal); err != nil {
			log.Printf("[ERROR] Failed to open modal: %s", err)
			w.WriteHeader(200)
			return
		}
		return
	case confirmSchoolModal:
		err := ConfirmSchoolHandle(payload, m)
		if err != nil {
			log.Printf("[ERROR] Failed to handle confirm school handle: %s", err)
			w.WriteHeader(200)
			return
		}
		return
	case submitStudentModal:
		modal, err := SubmitStudentModal(payload)
		if err != nil {
			log.Printf("[ERROR] Failed to handle submit student: %s", err)
			w.WriteHeader(200)
			return
		}
		if _, err := api.OpenView(payload,TriggerID, *modal); err != nil {
			log.Printf("[ERROR] Failed to open modal: %s", err)
			w.WriteHeader(200)
			return
		}
		return
	case confirmStudentModal:
		err := ConfirmStudentModal(payload, m)
		if err != nil {
			log.Printf("[ERROR] Failed to handle confirm student: %s", err)
			w.WriteHeader(200)
			return
		}
		return
	}	
}

func CheckCallbackType(payoad slack.InteractionCallback) string {
	if payload.Type == slack.InteractionTypeBlockActions && payload.View.Hash == "" {
		return buttonPushedAction
	}

	if payload.Type == slack.InteractionTypeViewSubmission && strings.Contains(payload.View.CallbackID, submitRestDateModal) {
		return submitRestDateModal
	}

	if payload.Type == slack.InteractionTypeViewSubmission && strings.Contains(payload.View.CallbackID, confirmRestDateModal) {
		return confirmRestDateModal
	}

	if payload.Type == slack.InteractionTypeViewSubmission && strings.Contains(payload.View.CallbackID, submitSchoolModal) {
		return submitSchoolModal
	}

	if payload.Type == slack.InteractionTypeViewSubmission && strings.Contains(payload.View.CallbackID, confirmSchoolModal) {
		return confirmSchoolModal
	}

	if payload.Type == slack.InteractionTypeViewSubmission && strings.Contains(payload.View.CallbackID, submitStudentModal) {
		return submitStudentModal
	}

	if payload.Type == slack.InteractionTypeViewSubmission && strings.Contains(payload.View.CallbackID, confirmStudentModal) {
		return confirmStudentModal
	}
}