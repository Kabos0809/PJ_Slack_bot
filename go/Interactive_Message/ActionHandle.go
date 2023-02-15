package Interactive_Message

import (
	"github.com/google/uuid"
	"github.com/slack-go/slack"
	"github.com/kabos0809/slack_bot/go/Models"
)

//学年
var grades = []string{"小1", "小2", "小3", "小4", "小5", "小6", "中1", "中2", "中3", "高校生"}

func SelectHandle(m Models.Model) *slack.ModalViewRequest {
	barText := slack.NewTextBlockObject("plain_text", "-------------", false, false)

	dstText := slack.NewTextBlockObject("mrkdwn", "*学校*・*学年*を選択してください", false, false)
	dstTextSection := slack.NewSectionBlock(dstText, nil, nil)

	dividerBlock := slack.NewDividerBlock()

	schools, _ := m.GetAllSchool()

	school_opt := make([]*slack.OptionBlockObject, 0, len(*schools))

	optText := slack.NewTextBlockObject("plain_text", "選択しない", false, false)
	school_opt = append(school_opt, slack.NewOptionBlockObject("選択しない", optText, nil))

	for _, v := range *schools {
		optText = slack.NewTextBlockObject("plain_text", v.Name, false, false)
		school_opt = append(school_opt, slack.NewOptionBlockObject(v.ID.String(), optText, nil))
	}

	school_label := slack.NewTextBlockObject("plain_text", "学校を選択してください", false, false)
	school_select := slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, barText, selectSchool, school_opt...)

	schoolInput := slack.NewInputBlock(selectSchool, school_label, nil, school_select)

	grades_opt := make([]*slack.OptionBlockObject, 0, len(grades))
	for _, v := range grades {
		optText := slack.NewTextBlockObject("plain_text", v, false, false)
		grades_opt = append(grades_opt, slack.NewOptionBlockObject(v, optText, nil))
	}

	grade_label := slack.NewTextBlockObject("plain_text", "学年を選択してください", false, false)
	grade_select := slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, barText, selectGrade, grades_opt...)

	gradeInput := slack.NewInputBlock(selectGrade, grade_label, nil, grade_select)

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			dstTextSection,
			dividerBlock,
			schoolInput,
			gradeInput,
		},
	}

	modal := slack.ModalViewRequest{
		Type: slack.ViewType("modal"),
		Title: slack.NewTextBlockObject("plain_text", "学校・学年選択", false, false),
		Close: slack.NewTextBlockObject("plain_text", "Cancel", false, false),
		Submit: slack.NewTextBlockObject("plain_text", "Submit", false, false),
		Blocks: blocks,
	}

	return &modal
}

func StudentListHandle(payload slack.InteractionCallback, m Models.Model) slack.MsgOption {
	barText := slack.NewTextBlockObject("plain_text", "-------------", false, false)
	
	dstText := slack.NewTextBlockObject("mrkdwn", "*生徒*を選択してください", false, false)
	dstTextSection := slack.NewSectionBlock(dstText, nil, nil)

	dividerBlock := slack.NewDividerBlock()

	school := payload.View.State.Values[selectSchool][selectSchool].SelectedOption.Value
	grade := payload.View.State.Values[selectGrade][selectGrade].SelectedOption.Value

	if (school == "選択しない") {
		students, err := m.GetStudentbyGrade(grade)
		if err != nil {
			errBlock := CreateErrorMsgBlock(InternalServerError)
			return errBlock
		}
		students_opt := make([]*slack.OptionBlockObject, 0, len(students))

		for _, s := range students {
			optText := slack.NewTextBlockObject("plain_text", s.Name, false, false)
			students_opt = append(students_opt, slack.NewOptionBlockObject(s.ID.String(), optText, nil))
		}

		student_select := slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, barText, selectStudent, students_opt...)

		actionBlock := slack.NewActionBlock(selectStudent, student_select)

		blocks := slack.MsgOptionBlocks(dstTextSection, dividerBlock, actionBlock)

		return blocks
	}

	school_id, err := uuid.Parse(school)
	if err != nil {
		errBlock := CreateErrorMsgBlock(InternalServerError)
		return errBlock
	}

	students, err := m.GetStudentbySchoolAndGrade(school_id, grade)
	if err != nil {
		errBlock := CreateErrorMsgBlock(InternalServerError)
		return errBlock
	}
	students_opt := make([]*slack.OptionBlockObject, 0, len(students))

	for _, s := range students {
		optText := slack.NewTextBlockObject("plain_text", s.Name, false, false)
		students_opt = append(students_opt, slack.NewOptionBlockObject(s.ID.String(), optText, nil))
	}

	student_select := slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, barText, selectStudent, students_opt...)

	actionBlock := slack.NewActionBlock(selectStudent, student_select)

	blocks := slack.MsgOptionBlocks(dstTextSection, dividerBlock, actionBlock)

	return blocks
}