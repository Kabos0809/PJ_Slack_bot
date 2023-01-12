package Mentioned_Message

import (
	"github.com/slack-go/slack"
	"github.com/kabos0809/slack_bot/go/Models"
)

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

//講師向けブロック
func createSelectBlock4Teachers(m Models.Model) slack.MsgOption {
	descText := slack.NewTextBlockObject("mrkdwn", "*学年*と*学校*を選択してください.", false, false)
	descTextSection := slack.NewSectionBlock(descText, nil, nil)

	dividerBlock := slack.NewDividerBlock()

	schools, err := m.GetAllSchool()
	if err != nil {
		errBlock := CreateErrorMsgBlock(NotFound, SelectSchool)
		return errBlock
	}

	school_opt := make([]*slack.OptionBlockObject, 0, len(*schools))

	if len(*schools) != 0 {
		for _, v := range *schools {
			optText := slack.NewTextBlockObject("plain_text", v.Name, false, false)
			school_opt = append(school_opt, slack.NewOptionBlockObject(string(v.ID), nil, optText))
		}
	} else {
		errBlock := CreateErrorMsgBlock(NotFound, SelectSchool)
		return errBlock
	}

	s_placeholder := slack.NewTextBlockObject("plain_text", "学校を選択してください", false, false)
	school_select := slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, s_placeholder, "", school_opt...)

	grades_opt := make([]*slack.OptionBlockObject, 0, len(grades))
	for _, v := range grades {
		optText := slack.NewTextBlockObject("plain_text", v, false, false)
		grades_opt = append(grades_opt, slack.NewOptionBlockObject(v, nil, optText))
	}

	g_placeholder := slack.NewTextBlockObject("plain_text", "学年を選択してください", false, false)
	grade_select := slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, g_placeholder, "", grades_opt...)

	actionBlock := slack.NewActionBlock(SchoolandGradeSelect, school_select, grade_select)

	blocks := slack.MsgOptionBlocks(descTextSection, dividerBlock, actionBlock)

	return blocks
}

//社員向けブロック
func createSelectBlock4Employee() slack.MsgOption {
	descText := slack.NewTextBlockObject("mrkdwn", "何を行いますか？", false, false)
	descTextSection := slack.NewSectionBlock(descText, nil, nil)

	dividerBlock := slack.NewDividerBlock()
	
	RestDateButtonText := slack.NewTextBlockObject("plain_text", "Do it", true, false)
	RestDateButtonElement := slack.NewButtonBlockElement("actionRestDate", "RestDate", RestDateButtonText)
	RestDateAccessory := slack.NewAccessory(RestDateButtonElement)
	RestDateSectionText := slack.NewTextBlockObject("mrkdwn", "*欠席の登録*\n欠席登録ができます", false, false)
	RestDateSection := slack.NewSectionBlock(RestDateSectionText, nil, RestDateAccessory)

	checkTransferCountButtonText := slack.NewTextBlockObject("plain_text", "Do it", true, false)
	checkTransferCountButtonElement := slack.NewButtonBlockElement("actionCheckTransferCount", "checkTransferCount", checkTransferCountButtonText)
	checkTransferCountAccessory := slack.NewAccessory(checkTransferCountButtonElement)
	checkTransferCountSectionText := slack.NewTextBlockObject("mrkdwn", "*残り振替回数確認*\n特定の生徒の残り振替回数の確認ができます", false, false)
	checkTransferCountSection := slack.NewSectionBlock(checkTransferCountSectionText, nil, checkTransferCountAccessory)

	studentButtonText := slack.NewTextBlockObject("plain_text", "Do it", true, false)
	studentButtonElement := slack.NewButtonBlockElement("actionStudentOperation", "student", studentButtonText)
	studentAccessory := slack.NewAccessory(studentButtonElement)
	studentSectionText := slack.NewTextBlockObject("mrkdwn", "*生徒情報の追加・編集*\n生徒情報の追加、編集、削除ができます", false, false)
	studentSection := slack.NewSectionBlock(studentSectionText, nil, studentAccessory)

	schoolButtonText := slack.NewTextBlockObject("plain_text", "Do it", true, false)
	schoolButtonElement := slack.NewButtonBlockElement("actionSchoolOperation", "school", schoolButtonText)
	schoolAccessory := slack.NewAccessory(schoolButtonElement)
	schoolSectionText := slack.NewTextBlockObject("mrkdwn", "*学校の追加・削除*\n学校の追加・削除ができます。", false, false)
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
			errText := slack.NewTextBlockObject("mrkdwn", "*学校を取得できませんでした*\n学校が一つも登録されていない可能性があります。", true, false)
			errTextSection := slack.NewSectionBlock(errText, nil, nil)
			block = slack.MsgOptionBlocks(errTextSection)
		default:
			errText := slack.NewTextBlockObject("mrkdwn", "*データを取得できませんでした*\n数学科 鶴賀までご連絡ください。:pray:", true, false)
			errTextSection := slack.NewSectionBlock(errText, nil, nil)
			block = slack.MsgOptionBlocks(errTextSection)
		}
	}

	return block
}