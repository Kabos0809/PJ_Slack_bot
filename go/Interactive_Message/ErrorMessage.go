package Interactive_Message

import "github.com/slack-go/slack"

const (
	NotFound = 404
	InternalServerError = 500
	BadRequest = 400
	UnAuthorized = 401
	BadGateWay = 502
)

func CreateErrorMsgBlock(status int) slack.MsgOption {
	var errText slack.MsgOption
	switch status{
	case NotFound:
		errText = slack.MsgOptionText("*データを取得できませんでした*\n数学科 鶴賀までご連絡ください。:pray:", false)
	case InternalServerError:
		errText = slack.MsgOptionText("*サーバー内部でエラーが発生しました*\n数学科 鶴賀までご連絡ください。 :pray:", false)
	case BadRequest:
		errText = slack.MsgOptionText("*不正なリクエストです*", false)

	return errText
}