package Interactive_Message

import (
	"github.com/kabos0809/slack-go/slack"
	"github.com/kabos0809/slack_bot/Models"
)

//学校に関する操作を選択するブロックを表示する関数
func SchoolButtonPushedActionHandle() slack.MsgOption {
	dstText := slack.NewTextBlockObject("mrkdwn", "*学校*に関する操作を行います", false, false)
	dstTextSection := slack.NewSectionBlock(dstText, nil, nil)

	dividerBlock := slack.NewDividerBlock()

	addButtonText := slack.NewTextBlockObject("plain_text", "Do it", false, false)
	addButtonElement := slack.NewButtonBlockElement("actionAddSchool", "add_school", addButtonText)
	addButtonAccessory := slack.NewAccessory(addButtonElement)
	addButtonSectionText := slack.NewTextBlockObject("mrkdwn", "*学校追加*\n学校の追加を行います", false, false)
	addButtonSection := slack.NewSectionBlock(addButtonSectionText, nil, addButtonAccessory)

	deleteButtonText := slack.NewTextBlockObject("plain_text", "Do it", false, false)
	deleteButtonElement := slack.NewButtonBlockElement("actionDeleteSchool", "delete_school", deleteButtonText)
	deleteButtonAccessory := slack.NewAccessory(deleteButtonElement)
	deleteButtonSectionText := slack.NewTextBlockObject("mrkdwn", "*学校削除*\n学校の削除を行います", false, false)
	deleteButtonSection := slack.NewSectionBlock(deleteButtonSectionText, nil, deleteButtonAccessory)

	blocks := slack.MsgOptionBlocks(dstTextSection, dividerBlock, addButtonSection, deleteButtonSection)

	return blocks
}