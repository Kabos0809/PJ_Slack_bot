package Interactive_Message

import (
	"github.com/slack-go/slack"
	"github.com/kabos0809/slack_bot/go/Models"
)

//振替回数を表示する関数
func CheckCountActionHandle(payload slack.InteractionCallback, m Models.Model) slack.MsgOption, error {
	stu_id := payload.ActionCallback.BlockActions.Value[0].Parse()
	student, err := m.GetStudentbyID(stu_id)
	
	if err != nil {
		errBlock := CreateErrorMsgBlock(NotFound)
		return errBlock, err
	}

	titleText := slack.NewTextBlockObject("mrkdwn", "*" + student.Name + "*の各教科残り振替回数.", false, false)
	titleTextSection := slack.NewSectionBlock(titleText, nil, nil)

	mathCount := slack.NewTextBlockObject("mrkdwn", "*数学*:" + student.MathCounts + "回", false, false)
	mathCountSection := slack.NewSectionBlock(mathCount, nil, nil)

	jpnCount := slack.NewTextBlockObject("mrkdwn", "*国語*:" + student.JpnCounts + "回", false, false)
	jpnCountSection := slack.NewSectionBlock(jpnCount, nil, nil)

	engCount := slack.NewTextBlockObject("mrkdwn", "*英語*:" + student.EngCounts + "回", false, false)
	engCountSection := slack.NewSectionBlock(engCount, nil, nil)

	blocks := slack.MsgOptionBlocks(titleTextSection, mathCountSection, jpnCountSection, engCountSection)

	return blocks, nil
}