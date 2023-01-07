package Mentioned_Message

import (
	"os"
	"io"
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

var slack_signing_secret string = os.Getenv("SLACK_SIGNING_SECRET")

/*
例外が発生した場合もSlack App側でリクエストを受け取れた場合は200を返すように推奨されているため、例外が発生した際は
ログにエラーメッセージを表示するが、200を返している。
*/

func MentionedHandler(w http.ResponseWriter, r *http.Request, api *slack.Client, db *gorm.DB) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("[ERROR] Failed to read Request body: %v", err)
	}
	//OptionNoVerifyTokenが推奨
	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
	if err != nil {
		log.Printf("[ERROR] Failed to parse request body: %v", err)
		w.WriteHeader(200)
		return
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var res *slackevents.ChallengeResponse
		if err := json.Unmarshal([]byte(body), &res); err != nil {
			log.Printf("[ERROR] Failed to unmarshal json: %v", err)
			w.WriteHeader(200)
			return
		}
		w.Header().Set("Content-Type", "text")
		if _, err := w.Write([]byte(res.Challenge)); err != nil {
			log.Printf("[ERROR] Failed to write Challenge to response: %v", err)
			w.WriteHeader(200)
			return
		}
		w.WriteHeader(200)
		return
	}

	if eventsAPIEvent.Type != slackevents.CallbackEvent {
		log.Printf("[ERROR] Unexpected event type: expect = CallbackEvent, actual = %v", eventsAPIEvent.Type)
		w.WriteHeader(200)
		return
	}
	switch ev := eventsAPIEvent.InnerEvent.Data.(type) {
	case *slackevents.AppMentionEvent:
		msg := createSelectBlock4Teachers(db)

		if _, _, err := api.PostMessage(ev.Channel, msg); err != nil {
			errMsg := CreateErrorMsgBlock(InternalServerError, "")
			api.SendMessage(ev.Channel, errMsg)
			log.Printf("[ERROR] Failed to send a message to Slack: %v", err)
			w.WriteHeader(200)
			return
		}
	default:
		w.WriteHeader(200)
		return
	}
}

/*
func verify(r *http.Request, sc string) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	header := http.Header{}
	for k, v := range r.Header {
		for _, s := range v	{
			header.Set(k, s)
		}
	}
	sv, err := slack.NewSecretsVerifier(header, sc)
	if err != nil {
		return err
	}

	sv.Write(body)
	return sv.Ensure()
}
*/