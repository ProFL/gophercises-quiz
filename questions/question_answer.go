package questions

import (
	"fmt"
	"strings"
)

type QuestionAnswer struct {
	question string
	answer   string
}

func (m *QuestionAnswer) AskQuestion() bool {
	var user_answer string
	fmt.Printf("%s\n: ", m.question)
	fmt.Scanln(&user_answer)
	return strings.EqualFold(m.answer, strings.TrimSpace(user_answer))
}
