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

const (
	selectSchool = "selectSchool"
	selectGrade = "selectGrade"
	selectSchoolAndGrade = "selectSchoolAndGrade"
	selectStudent = "selectStudent"
	selectAction = "checkCount_T"
	selectSchoolAndGrade4RestDate = "selectSchoolAndGrade4RestDate"
	buttonPushedAction = "buttonPushedAction"
	checkCount = "checkCount"
	addSchool = "addSchool"
	addRestDate = "addRestDate"
	addStudent = "addStudent"
)

var fallbackText slack.MsgOption = slack.MsgOptionText("This client is not supported.", false)

type privateMeta struct {
	ChannelID string `json:"channel_id"`
	StudentID string `json:"student_id"`
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
				if _, err := api.PostEphemeral(payload.Channel.ID, payload.User.ID, fallbackText, msg); err != nil {
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
			if _, err := api.PostEphemeral(payload.Channel.ID, payload.User.ID, fallbackText, msg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			return
		case "add_school":
			modal := AddSchoolModalHandle()

			modal.CallbackID = addSchool
			modal.ExternalID = payload.User.ID + strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
			
			params := privateMeta{
				ChannelID: payload.Channel.ID,
			}

			bytes, err := json.Marshal(params)
			if err != nil {
				msg := CreateErrorMsgBlock(InternalServerError)
				if _, err := api.PostEphemeral(payload.Channel.ID, payload.User.ID, fallbackText, msg); err != nil {
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
		/*
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
			if _, err := api.PostEphemeral(payload.Channel.ID, fallbackText, msg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			return
		*/
		case "restdate":
			msg := RestDateButtonPushedActionHandle()
			if _, err := api.PostEphemeral(payload.Channel.ID, payload.User.ID, fallbackText, msg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			return
		case "add_restdate":
			modal := SelectHandle(m)

			modal.CallbackID = selectSchoolAndGrade4RestDate
			modal.ExternalID = payload.User.ID + strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
			params := privateMeta{
				ChannelID: payload.Channel.ID,
			}
			bytes, err := json.Marshal(params)
			if err != nil {
				msg := CreateErrorMsgBlock(InternalServerError)
				if _, err := api.PostEphemeral(payload.Channel.ID, payload.User.ID, fallbackText, msg); err != nil {
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
		/*
		case "delete_restdate":
			msg := DeleteRestDateActionHandle()
			if _, err := api.PostEphemeral(payload.Channel.ID, fallbackText, msg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			return
		*/
		case "student":
			msg := StudentButtonPushedActionHandle()
			if _, err := api.PostEphemeral(payload.Channel.ID, payload.User.ID, fallbackText, msg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			w.WriteHeader(200)
			return
		case "add_student":
			modal := AddStudentModalHandle(m)
			
			modal.CallbackID = addStudent
			modal.ExternalID = payload.User.ID + strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

			params := privateMeta{
				ChannelID: payload.Channel.ID,
			}

			bytes, err := json.Marshal(params)
			if err != nil {
				msg := CreateErrorMsgBlock(InternalServerError)
				if _, err := api.PostEphemeral(payload.Channel.ID, payload.User.ID, fallbackText, msg); err != nil {
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
			w.WriteHeader(200)
			return
		default:
			switch payload.ActionCallback.BlockActions[0].ActionID {
			case selectStudent:
				msg, err := CheckCountActionHandle(payload, m)
				if err != nil {
					log.Printf("[ERROR] Failed to handle check count action handle: %s", err)
					w.WriteHeader(200)
					return
				}
				if _, err := api.PostEphemeral(payload.Channel.ID, payload.User.ID, fallbackText, msg); err != nil {
					log.Printf("[ERROR] Failed to send message: %s", err)
					w.WriteHeader(200)
					return
				}
			case addRestDateSelectStudent:
				modal := AddRestDateModalHandle(m)

				modal.CallbackID = addRestDate
				modal.ExternalID = payload.User.ID + strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
				
				params := privateMeta{
					ChannelID: payload.Channel.ID,
					StudentID: payload.ActionCallback.BlockActions[0].SelectedOption.Value,
				}

				log.Printf(params.StudentID)

				bytes, err := json.Marshal(params)
				if err != nil {
					msg := CreateErrorMsgBlock(InternalServerError)
					if _, err := api.PostEphemeral(payload.Channel.ID, payload.User.ID, fallbackText, msg); err != nil {
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
				w.WriteHeader(200)
				return
			}
			w.WriteHeader(200)
			return
		}
	case selectSchoolAndGrade4RestDate:
		var pMeta privateMeta
		if err := json.Unmarshal([]byte(payload.View.PrivateMetadata), &pMeta); err != nil {
			log.Printf("[ERROR] Failed to unmarshal json: %s", err)
			w.WriteHeader(200)
			return 
		}

		msg := StudentListHandle(payload, m, addRestDate)
		if _, err := api.PostEphemeral(pMeta.ChannelID, payload.User.ID, fallbackText, msg); err != nil {
			log.Printf("[ERROR] Failed to send message: %s", err)
			w.WriteHeader(200)
			return
		}
		return
	case selectSchoolAndGrade:
		var pMeta privateMeta
		if err := json.Unmarshal([]byte(payload.View.PrivateMetadata), &pMeta); err != nil {
			log.Printf("[ERROR] Failed to unmarshal json: %s", err)
			w.WriteHeader(200)
			return 
		}

		msg := StudentListHandle(payload, m, checkCount)
		if _, err := api.PostEphemeral(pMeta.ChannelID, payload.User.ID, fallbackText, msg); err != nil {
			log.Printf("[ERROR] Failed to send message: %s", err)
			w.WriteHeader(200)
			return
		}
		return
	case addSchool:
		var pMeta privateMeta
		if err := json.Unmarshal([]byte(payload.View.PrivateMetadata), &pMeta); err != nil {
			log.Printf("[ERROR] Failed to unmarshal json: %s", err)
			w.WriteHeader(200)
			return 
		}

		var school Models.School
		school.Name = payload.View.State.Values[addSchool][addSchool].Value

		err := m.AddSchool(&school)
		if err != nil {
			log.Printf("[ERROR] Failed to add school: %s", err)
			errMsg := CreateErrorMsgBlock(InternalServerError)
			if _, err := api.PostEphemeral(pMeta.ChannelID, payload.User.ID, fallbackText, errMsg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			w.WriteHeader(200)
			return
		}

		msg := slack.MsgOptionText("学校の登録が完了しました", false)
		if _, err := api.PostEphemeral(pMeta.ChannelID, payload.User.ID, fallbackText, msg); err != nil {
			log.Printf("[ERROR] Failed to send message: %s", err)
			w.WriteHeader(200)
			return
		}
		return
	case addRestDate:
		var e error
		var pMeta privateMeta
		if err := json.Unmarshal([]byte(payload.View.PrivateMetadata), &pMeta); err != nil {
			log.Printf("[ERROR] Failed to unmarshal json: %s", err)
			w.WriteHeader(200)
			return 
		}

		log.Printf(pMeta.ChannelID)
		log.Printf(pMeta.StudentID)

		var restdate Models.RestDate
		restdate.StudentID, e = uuid.Parse(pMeta.StudentID)
		restdate.Date, e = time.Parse("20060102", payload.View.State.Values[addRestDateDate][addRestDateDate].Value)
		restdate.Subject = payload.View.State.Values[addRestDateSelectSubject][addRestDateSelectSubject].SelectedOption.Value
		
		if e != nil {
			log.Printf("[ERROR] Failed to create restdate: %s", e)
			errMsg := CreateErrorMsgBlock(InternalServerError)
			if _, err := api.PostEphemeral(pMeta.ChannelID, payload.User.ID, fallbackText, errMsg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			w.WriteHeader(200)
			return
		}

		e = SucjectsVaridate(restdate.Subject)
		if e != nil {
			log.Printf("[ERROR] Invalid subject: %s", e)
			errMsg := CreateErrorMsgBlock(InternalServerError)
			if _, err := api.PostEphemeral(pMeta.ChannelID, payload.User.ID, fallbackText, errMsg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			w.WriteHeader(200)
			return
		}

		e = m.CreateRestDate(&restdate)
		if e != nil {
			log.Printf("[ERROR] Failed to create restdate: %s", e)
			errMsg := CreateErrorMsgBlock(InternalServerError)
			if _, err := api.PostEphemeral(pMeta.ChannelID, payload.User.ID, fallbackText, errMsg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			w.WriteHeader(200)
			return
		}

		e = m.AddRestDate4Student(&restdate, restdate.StudentID)
		if e != nil {
			log.Printf("[ERROR] Failed to add restdate for student: %s", e)
			errMsg := CreateErrorMsgBlock(InternalServerError)
			if _, err := api.PostEphemeral(pMeta.ChannelID, payload.User.ID, fallbackText, errMsg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			w.WriteHeader(200)
			return
		}


		msg := slack.MsgOptionText("欠席の登録が完了しました", false)

		if _, err := api.PostEphemeral(pMeta.ChannelID, payload.User.ID, fallbackText, msg); err != nil {
			log.Printf("[ERROR] Failed to send message: %s", err)
			w.WriteHeader(200)
			return
		}
		return

	case addStudent:
		var e error
		var pMeta privateMeta
		if err := json.Unmarshal([]byte(payload.View.PrivateMetadata), &pMeta); err != nil {
			log.Printf("[ERROR] Failed to unmarshal json: %s", err)
			w.WriteHeader(200)
			return 
		}
		
		var student Models.Student
		student.LastName = payload.View.State.Values[addStudentLastName][addStudentLastName].Value
		student.FirstName = payload.View.State.Values[addStudentFirstName][addStudentFirstName].Value
		student.SchoolID, e = uuid.Parse(payload.View.State.Values[addStudentSelectSchool][addStudentSelectSchool].SelectedOption.Value)
		student.Grade = payload.View.State.Values[addStudentSelectGrade][addStudentSelectGrade].SelectedOption.Value
		
		if e != nil {
			log.Printf("[ERROR] Failed to create student: %s", e)
			errMsg := CreateErrorMsgBlock(InternalServerError)
			if _, err := api.PostEphemeral(pMeta.ChannelID, payload.User.ID, fallbackText, errMsg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			w.WriteHeader(200)
			return
		}

		err := m.CreateStudent(&student)
		if err != nil {
			log.Printf("[ERROR] Failed to create student: %s", err)
			errMsg := CreateErrorMsgBlock(InternalServerError)
			if _, err := api.PostEphemeral(pMeta.ChannelID, payload.User.ID, fallbackText, errMsg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			w.WriteHeader(200)
			return
		}

		err = m.AddStudent4School(&student, student.SchoolID)
		if err != nil {
			log.Printf("[ERROR] Failed to create student: %s", err)
			errMsg := CreateErrorMsgBlock(InternalServerError)
			if _, err := api.PostEphemeral(pMeta.ChannelID, payload.User.ID, fallbackText, errMsg); err != nil {
				log.Printf("[ERROR] Failed to send message: %s", err)
				w.WriteHeader(200)
				return
			}
			w.WriteHeader(200)
			return
		}

		msg := slack.MsgOptionText("生徒の登録が完了しました", false)

		if _, err := api.PostEphemeral(pMeta.ChannelID, payload.User.ID, fallbackText, msg); err != nil {
			log.Printf("[ERROR] Failed to send message: %s", err)
			w.WriteHeader(200)
			return
		}
		return
	}	
	return
}

func CheckCallbackType(payload slack.InteractionCallback) string {
	if payload.Type == slack.InteractionTypeBlockActions && payload.View.Hash == "" {
		return buttonPushedAction
	}
	if payload.Type == slack.InteractionTypeViewSubmission && strings.Contains(payload.View.CallbackID, selectSchoolAndGrade4RestDate) {
		return selectSchoolAndGrade4RestDate
	}
	if payload.Type == slack.InteractionTypeViewSubmission && strings.Contains(payload.View.CallbackID, selectSchoolAndGrade) {
		return selectSchoolAndGrade
	}
	if payload.Type == slack.InteractionTypeViewSubmission && strings.Contains(payload.View.CallbackID, addSchool) {
		return addSchool
	}
	if payload.Type == slack.InteractionTypeViewSubmission && strings.Contains(payload.View.CallbackID, addStudent) {
		return addStudent
	}
	if payload.Type == slack.InteractionTypeViewSubmission && strings.Contains(payload.View.CallbackID, addRestDate) {
		return addRestDate
	}
	return ""
}