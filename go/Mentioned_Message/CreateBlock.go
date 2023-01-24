package Mentioned_Message

import (
	"github.com/slack-go/slack"
	"github.com/kabos0809/slack_bot/go/Models"
)

//学年
var grades = []string{"小1", "小2", "小3", "小4", "小5", "小6", "中1", "中2", "中3", "高校生"}

//一応エラー発生時に使う予定だが、そもそもエラーが発生した際にエラーメッセージを送信することができないかもしれないので使わない可能性大
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
	descText := slack.NewTextBlockObject("mrkdwn", "*学年*と*学校*を選択してください.", true, true)
	descTextSection := slack.NewSectionBlock(descText, nil, nil)

	dividerBlock := slack.NewDividerBlock()

	schools, err := m.GetAllSchool()
	if err != nil {
		errBlock := CreateErrorMsgBlock(InternalServerError, SelectSchool)
		return errBlock
	}

	school_opt := make([]*slack.OptionBlockObject, 0, len(*schools))

	if len(*schools) != 0 {
		for _, v := range *schools {
			optText := slack.NewTextBlockObject("plain_text", v.Name, true, true)
			school_opt = append(school_opt, slack.NewOptionBlockObject(string(v.ID), nil, optText))
		}
	} else {
		errBlock := CreateErrorMsgBlock(NotFound, SelectSchool)
		return errBlock
	}

	s_placeholder := slack.NewTextBlockObject("plain_text", "学校を選択してください", true, true)
	school_select := slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, s_placeholder, "", school_opt...)

	grades_opt := make([]*slack.OptionBlockObject, 0, len(grades))
	for _, v := range grades {
		optText := slack.NewTextBlockObject("plain_text", v, true, true)
		grades_opt = append(grades_opt, slack.NewOptionBlockObject(v, nil, optText))
	}

	g_placeholder := slack.NewTextBlockObject("plain_text", "学年を選択してください", true, true)
	grade_select := slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, g_placeholder, "", grades_opt...)

	actionBlock := slack.NewActionBlock(SchoolandGradeSelect, school_select, grade_select)

	blocks := slack.MsgOptionBlocks(descTextSection, dividerBlock, actionBlock)

	return blocks
}

//社員向けブロックを作成する関数
func createSelectBlock4Employee() slack.MsgOption {
	descText := slack.NewTextBlockObject("mrkdwn", "何を行いますか？", true, true)
	descTextSection := slack.NewSectionBlock(descText, nil, nil)

	dividerBlock := slack.NewDividerBlock()
	
	RestDateButtonText := slack.NewTextBlockObject("plain_text", "Do it", true, true)
	RestDateButtonElement := slack.NewButtonBlockElement("actionRestDate", "RestDate", RestDateButtonText)
	RestDateAccessory := slack.NewAccessory(RestDateButtonElement)
	RestDateSectionText := slack.NewTextBlockObject("mrkdwn", "*欠席の登録*\n欠席登録ができます", true, true)
	RestDateSection := slack.NewSectionBlock(RestDateSectionText, nil, RestDateAccessory)

	checkTransferCountButtonText := slack.NewTextBlockObject("plain_text", "Do it", true, true)
	checkTransferCountButtonElement := slack.NewButtonBlockElement("actionCheckTransferCount", "checkTransferCount", checkTransferCountButtonText)
	checkTransferCountAccessory := slack.NewAccessory(checkTransferCountButtonElement)
	checkTransferCountSectionText := slack.NewTextBlockObject("mrkdwn", "*残り振替回数確認*\n特定の生徒の残り振替回数の確認ができます", true, true)
	checkTransferCountSection := slack.NewSectionBlock(checkTransferCountSectionText, nil, checkTransferCountAccessory)

	studentButtonText := slack.NewTextBlockObject("plain_text", "Do it", true, true)
	studentButtonElement := slack.NewButtonBlockElement("actionStudentOperation", "student", studentButtonText)
	studentAccessory := slack.NewAccessory(studentButtonElement)
	studentSectionText := slack.NewTextBlockObject("mrkdwn", "*生徒情報の追加・編集*\n生徒情報の追加、編集、削除ができます", true, true)
	studentSection := slack.NewSectionBlock(studentSectionText, nil, studentAccessory)

	schoolButtonText := slack.NewTextBlockObject("plain_text", "Do it", true, true)
	schoolButtonElement := slack.NewButtonBlockElement("actionSchoolOperation", "school", schoolButtonText)
	schoolAccessory := slack.NewAccessory(schoolButtonElement)
	schoolSectionText := slack.NewTextBlockObject("mrkdwn", "*学校の追加・削除*\n学校の追加・削除ができます。", true, true)
	schoolSection := slack.NewSectionBlock(schoolSectionText, nil, schoolAccessory)

	blocks := slack.MsgOptionBlocks(descTextSection, dividerBlock, checkTransferCountSection, RestDateSection, studentSection, schoolSection)

	return blocks
}

//内部でエラーが発生したとき用のエラーメッセージ
func CreateErrorMsgBlock(status int, v string) slack.MsgOption {
	var block slack.MsgOption
	switch status{
	case NotFound:
		switch v{
		case SelectSchool:
			errText := slack.NewTextBlockObject("mrkdwn", "*学校を取得できませんでした*\n学校が一つも登録されていない可能性があります。", true, true)
			errTextSection := slack.NewSectionBlock(errText, nil, nil)
			block = slack.MsgOptionBlocks(errTextSection)
		default:
			errText := slack.NewTextBlockObject("mrkdwn", "*データを取得できませんでした*\n数学科 鶴賀までご連絡ください。:pray:", true, true)
			errTextSection := slack.NewSectionBlock(errText, nil, nil)
			block = slack.MsgOptionBlocks(errTextSection)
		}
	case InternalServerError:
		errText := slack.NewTextBlockObject("mrkdwn", "*サーバー内部でエラーが発生しました*\n数学科 鶴賀までご連絡ください。 :pray:", true, true)
		errTextSection := slack.NewSectionBlock(errText, nil, nil)
		block = slack.MsgOptionBlocks(errTextSection)
	}

	return block
}