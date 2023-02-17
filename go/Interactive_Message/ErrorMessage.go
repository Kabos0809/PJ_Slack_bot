package Interactive_Message

import "github.com/slack-go/slack"

const (
	NotFound = 404
	InternalServerError = 500
	BadRequest = 400
	UnAuthorized = 401
	BadGateWay = 502
	schoolNotExist = 1404
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
	case schoolNotExist:
		errText = slack.MsgOptionText("*学校が存在しないため生徒の追加ができません*\n先に学校の追加をしてください", false)
	}
	
	return errText
}