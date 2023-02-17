package Interactive_Message

import (
	"github.com/slack-go/slack"
	"github.com/kabos0809/slack_bot/go/Models"
)

const (
	addStudentLastName = "addStudentLastName"
	addStudentFirstName = "addStudentFirstName"
	addStudentSelectSchool = "addStudentSelectSchool"
	addStudentSelectGrade = "addStudentSelectGrade"
)

//生徒情報に対する操作を選択するブロックを作成する関数	
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
	updateButtonAccessory := slack.NewAccessory(updateButtonElement)
	updateButtonSectionText := slack.NewTextBlockObject("mrkdwn", "*生徒情報の編集*\n生徒情報の編集を行います", false, false)
	updateButtonSection := slack.NewSectionBlock(updateButtonSectionText, nil, updateButtonAccessory)

	deleteButtonText := slack.NewTextBlockObject("plain_text", "Do it", false, false)
	deleteButtonElement := slack.NewButtonBlockElement("actionDeleteStudent", "delete_student", deleteButtonText)
	deleteButtonAccessory := slack.NewAccessory(deleteButtonElement)
	deleteButtonSectionText := slack.NewTextBlockObject("mrkdwn", "*生徒情報の削除*\n生徒情報の削除を行います", false, false)
	deleteButtonSection := slack.NewSectionBlock(deleteButtonSectionText, nil, deleteButtonAccessory)

	blocks := slack.MsgOptionBlocks(dstTextSection, dividerBlock, addButtonSection, updateButtonSection, deleteButtonSection)

	return blocks
}

//生徒情報を記入するモーダルを作成する関数
func AddStudentModalHandle(m Models.Model) *slack.ModalViewRequest {
	barText := slack.NewTextBlockObject("plain_text", "-------------", false, false)

	dstText := slack.NewTextBlockObject("mrkdwn", "*生徒の追加*を行います\n", false, false)
	dstTextSection := slack.NewSectionBlock(dstText, nil, nil)

	dividerBlock := slack.NewDividerBlock()

	lastNameText := slack.NewTextBlockObject("plain_text", "苗字", false, false)
	lastNameInputElement := slack.NewPlainTextInputBlockElement(nil, addStudentLastName)
	lastNameInput := slack.NewInputBlock(addStudentLastName, lastNameText, nil, lastNameInputElement)

	firstNameText := slack.NewTextBlockObject("plain_text", "名前", false, false)
	firstNameInputElement := slack.NewPlainTextInputBlockElement(nil, addStudentFirstName)
	firstNameInput := slack.NewInputBlock(addStudentFirstName, firstNameText, nil, firstNameInputElement)

	schools, _ := m.GetAllSchool()
	schoolOpt := make([]*slack.OptionBlockObject, 0, len(*schools))
	for _, s := range *schools {
		optText := slack.NewTextBlockObject("plain_text", s.Name, false, false)
		schoolOpt = append(schoolOpt, slack.NewOptionBlockObject(s.ID.String(), optText, nil))
	}

	schoolLabel := slack.NewTextBlockObject("plain_text", "学校", false, false)
	schoolSelect := slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, barText, addStudentSelectSchool, schoolOpt...)

	schoolInput := slack.NewInputBlock(addStudentSelectSchool, schoolLabel, nil, schoolSelect)

	gradeOpt := make([]*slack.OptionBlockObject, 0, len(grades))

	for _, v := range grades {
		optText := slack.NewTextBlockObject("plain_text", v, false, false)
		gradeOpt = append(gradeOpt, slack.NewOptionBlockObject(v, optText, nil))
	}

	gradeLabel := slack.NewTextBlockObject("plain_text", "学年", false, false)
	gradeSelect := slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, barText, addStudentSelectGrade, gradeOpt...)

	gradeInput := slack.NewInputBlock(addStudentSelectGrade, gradeLabel, nil, gradeSelect)

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			dstTextSection,
			dividerBlock,
			lastNameInput,
			firstNameInput,
			schoolInput,
			gradeInput,
		},
	}

	modal := slack.ModalViewRequest{
		Type: slack.ViewType("modal"),
		Title: slack.NewTextBlockObject("plain_text", "生徒の追加", false, false),
		Close: slack.NewTextBlockObject("plain_text", "Cancel", false, false),
		Submit: slack.NewTextBlockObject("plain_text", "Submit", false, false),
		Blocks: blocks,
	}

	return &modal
}