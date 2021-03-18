package questions

import (
	"fmt"
	"log"
	"strings"
)

type QuestionAnswer struct {
	question string
	answer   string
}

func (m *QuestionAnswer) AskQuestion() bool {
	var userAnswer string
	fmt.Printf("%s\n: ", m.question)
	_, err := fmt.Scanln(&userAnswer)
	if err != nil {
		log.Println(err.Error())
		log.Printf("Failed to read the answer for the question \"%s\"\n", m.question)
	}
	return strings.EqualFold(m.answer, strings.TrimSpace(userAnswer))
}
