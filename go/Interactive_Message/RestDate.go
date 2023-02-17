package Interactive_Message

import (
	"strings"
	"errors"
	"github.com/slack-go/slack"
	"github.com/kabos0809/slack_bot/go/Models"
)

var subjects = []string{"国語", "数学", "英語"} 

const (
	addRestDateSelectStudent = "addRestDateSelectStudent"
	addRestDateSelectSubject = "addRestDateSelectSubject"
	addRestDateDate = "addRestDateDate"
)
//欠席情報に関する操作を選択するブロックを表示する関数
func RestDateButtonPushedActionHandle() slack.MsgOption {
	dstText := slack.NewTextBlockObject("mrkdwn", "*欠席情報*に関する操作を行います\n", false, false)
	dstTextSection := slack.NewSectionBlock(dstText, nil, nil)

	dividerBlock := slack.NewDividerBlock()

	addButtonText := slack.NewTextBlockObject("plain_text", "Do it", false, false)
	addButtonElement := slack.NewButtonBlockElement("addRestDateAction", "add_restdate", addButtonText)
	addButtonAccessory := slack.NewAccessory(addButtonElement)
	addButtonSectionText := slack.NewTextBlockObject("mrkdwn", "*欠席情報の追加*\n欠席情報の追加を行います", false, false)
	addButtonSection := slack.NewSectionBlock(addButtonSectionText, nil, addButtonAccessory)

	deleteButtonText := slack.NewTextBlockObject("plain_text", "Do it", false, false)
	deleteButtonElement := slack.NewButtonBlockElement("deleteRestDateAction", "delete_restdate", deleteButtonText)
	deleteButtonAccessory := slack.NewAccessory(deleteButtonElement)
	deleteButtonSectionText := slack.NewTextBlockObject("mrkdwn", "*欠席情報の削除*\n欠席情報の削除を行います", false, false)
	deleteButtonSection := slack.NewSectionBlock(deleteButtonSectionText, nil, deleteButtonAccessory)

	blocks := slack.MsgOptionBlocks(dstTextSection, dividerBlock, addButtonSection, deleteButtonSection)

	return blocks
}

//欠席情報を記入するモーダルを作成する関数
func AddRestDateModalHandle(m Models.Model) *slack.ModalViewRequest {
	barText := slack.NewTextBlockObject("plain_text", "-------------", false, false)

	dstText := slack.NewTextBlockObject("mrkdwn", "*欠席情報を登録します*\n", false, false)
	dstTextSection := slack.NewSectionBlock(dstText, nil, nil)

	dividerBlock := slack.NewDividerBlock()

	subjectOpt := make([]*slack.OptionBlockObject, 0, len(subjects))
	for _, s := range subjects {
		optText := slack.NewTextBlockObject("plain_text", s, false, false)
		subjectOpt = append(subjectOpt, slack.NewOptionBlockObject(s, optText, nil))
	}
	subjectLabel := slack.NewTextBlockObject("plain_text", "教科", false, false)
	subjectSelect := slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, barText, addRestDateSelectSubject, subjectOpt...)
	subjectInput := slack.NewInputBlock(addRestDateSelectSubject, subjectLabel, nil, subjectSelect)

	dateText := slack.NewTextBlockObject("plain_text", "日にち", false, false)
	dateHint := slack.NewTextBlockObject("plain_text", "YYYYMMDD形式で入力してください(例:20230301)", false, false)
	dateInputElement := slack.NewPlainTextInputBlockElement(nil, addRestDateDate) 
	dateInput := slack.NewInputBlock(addRestDateDate, dateText, dateHint, dateInputElement)

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			dstTextSection,
			dividerBlock,
			subjectInput,
			dateInput,
		},
	}

	modal := slack.ModalViewRequest{
		Type: slack.ViewType("modal"),
		Title: slack.NewTextBlockObject("plain_text", "欠席情報入力", false, false),
		Close: slack.NewTextBlockObject("plain_text", "Cancel", false, false),
		Submit: slack.NewTextBlockObject("plain_text", "Submit", false, false),
		Blocks: blocks,
	}

	return &modal
}

func SucjectsVaridate(subject string) error {
	if (strings.Contains(subject, subjects[0]) || strings.Contains(subject, subjects[1]) || strings.Contains(subject, subjects[2])) {
		return nil
	}
	err := errors.New("Subjects must be Japanese, Mathmetics or English")
	return err
}