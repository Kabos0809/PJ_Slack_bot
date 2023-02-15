package Interactive_Message

import (
	"log"
	"time"
	"strconv"
	"net/http"
	"strings"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/slack-go/slack"
	"github.com/kabos0809/slack_bot/go/Models"
)

var (
	selectSchool = "selectSchool"
	selectGrade = "selectGrade"
	selectSchoolAndGrade = "selectSchoolAndGrade"
	selectStudent = "selectStudent"
	selectAction = "checkCount_T"
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

type privateMeta struct {
	ChannelID string `json:"channel_id"`
	school_data
	restdate_data
	student_data
}

type school_data struct {
	Name string `json:"name"`
}

type restdate_data struct {
	studentID uuid.UUID `json:"student_id"`
	date time.Time `json:"date"`
	subject string `json:"subject"`
}

type student_data struct {
	firstName string `json:"first_name"`
	lastName string `json:"last_name"`
	grade string `json:"grade"`
	schoolID uuid.UUID `json:"school_id"`
}

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
		if (payload.ActionCallback.BlockActions[0].Value == selectAction) {
			modal := SelectHandle(m)

			modal.CallbackID = selectSchoolAndGrade
			modal.ExternalID = payload.User.ID + strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
			params := privateMeta{
				ChannelID: payload.Channel.ID,
			}
			bytes, err := json.Marshal(params)
			if err != nil {
				msg := CreateErrorMsgBlock(InternalServerError)
				if _, _, _, err := api.SendMessage(payload.Channel.ID, fallbackText, msg); err != nil {
					log.Printf("[ERROR] Failed to send message: %s", err)
					w.WriteHeader(200)
					return
				}
			}
			modal.PrivateMetadata = string(bytes)

			if _, err := api.OpenView(payload.TriggerID, *modal); err != nil {
				log.Printf("[ERROR] Failed to open modal: %s", err)
				w.WriteHeader(200)
				return
			}
			return
		}
		switch payload.ActionCallback.BlockActions[0].Value {
		case "school":
			msg := SchoolButtonPushedActionHandle()
			if _, _, _, err := api.SendMessage(payload.Channel.ID, fallbackText, msg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			return
		/*
		case "add_school":
			modal, err := AddSchoolModalHandle()
			if err != nil {
				log.Printf("[ERROR] Failed to handle to create add school modal")
				w.WriteHeader(200)
				return
			}
			if _, err := api.OpenView(payload.TriggerID, *modal); err != nil {
				log.Printf("[ERROR] Failed to open modal: %s", err)
				w.WriteHeader(200)
				return
			}
			return
		case "update_school":
			modal, err := UpdateSchoolModalHandle()
			if err != nil {
				log.Printf("[ERROR] Failed to handle to create update school modal: %s", err)
				w.WriteHeader(200)
				return
			}
			if _, err := api.OpenView(payload.TriggerID, *modal); err != nil {
				log.Printf("[ERROR] Failed to open modal: %s", err)
				w.WriteHeader(200)
				return
			}
			return
		case "delete_school":
			msg := DeleteSchoolActionHandle()
			if _, _, _, err := api.SendMessage(payload.Channel.ID, fallbackText, msg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			return
		*/
		case "restdate":
			msg := RestDateButtonPushedActionHandle()
			if _, _, _, err := api.SendMessage(payload.Channel.ID, fallbackText, msg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			return
		/*
		case "add_restdate":
			modal, err := AddRestDateModalHandle()
			if err != nil {
				log.Printf("[ERROR] Failed to handle to create add restdate modal: %s", err)
				w.WriteHeader(200)
				return
			}
			return
		case "delete_restdate":
			msg := DeleteRestDateActionHandle()
			if _, _, _, err := api.SendMessage(payload.Channel.ID, fallbackText, msg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			return
		*/
		case "student":
			msg := StudentButtonPushedActionHandle()
			if _, _, _, err := api.SendMessage(payload.Channel.ID, fallbackText, msg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
		/*
		case "add_student":
			modal, err := AddStudentModalHandle()
			if _, err := api.OpenView(payload.TriggerID, *modal); err != nil {
				log.Printf("[ERROR] Failed to open modal: %s", err)
				w.WriteHeader(200)
				return
			}
			return
		*/
		default:
			msg, err := CheckCountActionHandle(payload, m)
			if err != nil {
				log.Printf("[ERROR] Failed to handle check count action handle: %s", err)
				w.WriteHeader(200)
				return
			}
			if _, _, _, err := api.SendMessage(payload.Channel.ID, fallbackText, msg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
		}
	case selectSchoolAndGrade:
		var pMeta privateMeta
		if err := json.Unmarshal([]byte(payload.View.PrivateMetadata), &pMeta); err != nil {
			log.Printf("[ERROR] Failed to unmarshal json: %s", err)
			w.WriteHeader(200)
			return 
		}

		msg := StudentListHandle(payload, m)
		if _, _, _, err := api.SendMessage(pMeta.ChannelID, fallbackText, msg); err != nil {
			log.Printf("[ERROR] Failed to send message: %s", err)
			w.WriteHeader(200)
			return
		}
	/*
	case confirmRestDateModal:
		err := ConfirmRestDateHandle(payload, m)
		if err != nil {
			log.Printf("[ERROR] Failed to handle confirm rest date: %s", err)
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
	case confirmStudentModal:
		err := ConfirmStudentModal(payload, m)
		if err != nil {
			log.Printf("[ERROR] Failed to handle confirm student: %s", err)
			w.WriteHeader(200)
			return
		}
		return
	*/
	}	
	return
}

func CheckCallbackType(payload slack.InteractionCallback) string {
	if payload.Type == slack.InteractionTypeBlockActions && payload.View.Hash == "" {
		return buttonPushedAction
	}
	if payload.Type == slack.InteractionTypeViewSubmission && strings.Contains(payload.View.CallbackID, confirmRestDateModal) {
		return confirmRestDateModal
	}
	if payload.Type == slack.InteractionTypeViewSubmission && strings.Contains(payload.View.CallbackID, confirmSchoolModal) {
		return confirmSchoolModal
	}
	if payload.Type == slack.InteractionTypeViewSubmission && strings.Contains(payload.View.CallbackID, confirmStudentModal) {
		return confirmStudentModal
	}
	if payload.Type == slack.InteractionTypeViewSubmission && strings.Contains(payload.View.CallbackID, selectSchoolAndGrade) {
		return selectSchoolAndGrade
	}
	return ""
}