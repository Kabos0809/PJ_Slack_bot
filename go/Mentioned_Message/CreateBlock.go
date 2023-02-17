package Mentioned_Message

import (
	"github.com/slack-go/slack"
	"github.com/kabos0809/slack_bot/go/Models"
)

//学年
var grades = []string{"小1", "小2", "小3", "小4", "小5", "小6", "中1", "中2", "中3", "高校生"}

const (
	NotFound = 404
	InternalServerError = 500
	BadRequest = 400
	UnAuthorized = 401
	BadGateWay = 502
	SelectSchool = "select_school"
	SchoolandGradeSelect = "school_and_grade_select"
)

//講師向けブロックを作成する関数
func createSelectBlock4Teachers(m Models.Model) slack.MsgOption {
	descText := slack.NewTextBlockObject("mrkdwn", "*学年*と*学校*を選択してください\n", false, false)
	descTextSection := slack.NewSectionBlock(descText, nil, nil)

	dividerBlock := slack.NewDividerBlock()
	
	checkCountButtonText := slack.NewTextBlockObject("plain_text", "Do it", false, false)
	checkCountButtonElement := slack.NewButtonBlockElement("actioncheckCount_T", "checkCount_T", checkCountButtonText)
	checkCountAccessory := slack.NewAccessory(checkCountButtonElement)
	checkCountSectionText := slack.NewTextBlockObject("mrkdwn", "*振替回数確認*\n生徒の残り振替回数の確認ができます", false, false)
	checkCountSection := slack.NewSectionBlock(checkCountSectionText, nil, checkCountAccessory)

	blocks := slack.MsgOptionBlocks(descTextSection, dividerBlock, checkCountSection)

	return blocks
}

//社員向けブロックを作成する関数
func createSelectBlock4Employee() slack.MsgOption {
	descText := slack.NewTextBlockObject("mrkdwn", "*何を行いますか？*\n", false, false)
	descTextSection := slack.NewSectionBlock(descText, nil, nil)

	dividerBlock := slack.NewDividerBlock()
	
	RestDateButtonText := slack.NewTextBlockObject("plain_text", "Do it", false, false)
	RestDateButtonElement := slack.NewButtonBlockElement("actionRestDate", "restdate", RestDateButtonText)
	RestDateAccessory := slack.NewAccessory(RestDateButtonElement)
	RestDateSectionText := slack.NewTextBlockObject("mrkdwn", "*欠席の登録*\n欠席登録ができます", false, false)
	RestDateSection := slack.NewSectionBlock(RestDateSectionText, nil, RestDateAccessory)

	checkTransferCountButtonText := slack.NewTextBlockObject("plain_text", "Do it", false, false)
	checkTransferCountButtonElement := slack.NewButtonBlockElement("actionCheckTransferCount", "checkCount_T", checkTransferCountButtonText)
	checkTransferCountAccessory := slack.NewAccessory(checkTransferCountButtonElement)
	checkTransferCountSectionText := slack.NewTextBlockObject("mrkdwn", "*残り振替回数確認*\n特定の生徒の残り振替回数の確認ができます", false, false)
	checkTransferCountSection := slack.NewSectionBlock(checkTransferCountSectionText, nil, checkTransferCountAccessory)

	studentButtonText := slack.NewTextBlockObject("plain_text", "Do it", false, false)
	studentButtonElement := slack.NewButtonBlockElement("actionStudentOperation", "student", studentButtonText)
	studentAccessory := slack.NewAccessory(studentButtonElement)
	studentSectionText := slack.NewTextBlockObject("mrkdwn", "*生徒情報の追加・編集*\n生徒情報の追加、編集、削除ができます", false, false)
	studentSection := slack.NewSectionBlock(studentSectionText, nil, studentAccessory)

	schoolButtonText := slack.NewTextBlockObject("plain_text", "Do it", false, false)
	schoolButtonElement := slack.NewButtonBlockElement("actionSchoolOperation", "school", schoolButtonText)
	schoolAccessory := slack.NewAccessory(schoolButtonElement)
	schoolSectionText := slack.NewTextBlockObject("mrkdwn", "*学校の追加・削除*\n学校の追加・削除ができます。", false, false)
	schoolSection := slack.NewSectionBlock(schoolSectionText, nil, schoolAccessory)

	blocks := slack.MsgOptionBlocks(descTextSection, dividerBlock, checkTransferCountSection, RestDateSection, studentSection, schoolSection)

	return blocks
}

//内部でエラーが発生したとき用のエラーメッセージ
func CreateErrorMsgBlock(status int, v string) slack.MsgOption {
	var errText slack.MsgOption
	switch status{
	case NotFound:
		switch v{
		case SelectSchool:
			errText = slack.MsgOptionText("*学校を取得できませんでした*\n学校が一つも登録されていません。", false)
		default:
			errText = slack.MsgOptionText("*データを取得できませんでした*\n数学科 鶴賀までご連絡ください。:pray:", false)
		}
	case InternalServerError:
		errText = slack.MsgOptionText("*サーバー内部でエラーが発生しました*\n数学科 鶴賀までご連絡ください。 :pray:", false)
	}

	return errText
}