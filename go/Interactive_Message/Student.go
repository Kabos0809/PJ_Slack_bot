package Interactive_Message

import (
	"github.com/slack-go/slack"
	"github.com/kasbo0809/slack_bot/go/Models"
)

func StudentButtonPushedActionHandle() slack.MsgOption {
	dstText := slack.NewTextBlockObject("mrkdwn", "*生徒情報*に関する操作を行います", false, false)
	dstTextSection := slack.NewSectionBlock(dstText, nil, nil)

	dividerBlock := slack.NewDividerBlock()

	addButtonText := slack.NewTextBlockObject("plain_text", "Do it", false, false)
	addButtonElement := slack.NewButtonBlockElement("actionAddStudent", "add_student", addButtonText)
	addButtonAccessory := slack.NewAccessory(addButtonElement)
	addButtonSectionText := slack.NewTextBlockObject("mrkdwn", "*生徒情報の追加*\n生徒情報の追加を行います", false, false)
	addButtonSection := slack.NewSectionBlock(addButtonSectionText, nil, addButtonAccessory)

	updateButtonText := slack.NewTextBlockObject("plain_text", "Do it", false, false)
	updateButtonElement := slack.NewButtonBlockElement("actionUpdateStudent", "update_student", updateButtonText)
	updateButtonAccesory := slack.NewAccessory(updateButtonElement)
	updateButtonSectionText := slack.NewTextBlockObject("mrkdwn", "*生徒情報の編集\n生徒情報の編集を行います", false, false)
	updateButtonSection := slack.NewSectionBlock(updateButtonSectionText, nil, updateButtonAccesory)

	deleteButtonText := slack.NewTextBlockObject("plain_text", "Do it", false, false)
	deleteButtonElement := slack.NewButtonBlockElement("actionDeleteStudent", "delete_student", deleteButtonText)
	deleteButtonAccessory := slack.NewAccessory(deleteButtonElement)
	deleteButtonSectionText := slack.NewTextBlockObject("mrkdwn", "*生徒情報の削除*\n生徒情報の削除を行います", false, false)
	deleteButtonSection := slack.NewSectionBlock(deleteButtonSectionText, nil, deleteButtonAccessory)

	blocks := slack.MsgOptionBlocks(dstText, dividerBlock, addButtonSection, updateButtonSection, deleteButtonSection)

	return blocks
}