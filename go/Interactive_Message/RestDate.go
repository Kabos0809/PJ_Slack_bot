package Interactive_Message

import (
	"github.com/slack-go/slack"
	//"github.com/kabos0809/slack_bot/go/Models"
)

//欠席情報に関する操作を選択するブロックを表示する関数
func RestDateButtonPushedActionHandle() slack.MsgOption {
	dstText := slack.NewTextBlockObject("mrkdwn", "*欠席情報*に関する操作を行います", false, false)
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