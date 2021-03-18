package questions

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func ReadAndParseQuestions(filePath string, shuffleQuestions bool) []QuestionAnswer {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		log.Panic("There was an error reading the questions file", err)
	}
	question_answers := parseQuestionsFromFile(file)
	if shuffleQuestions {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(question_answers), func(i, j int) {
			question_answers[i], question_answers[j] = question_answers[j], question_answers[i]
		})
	}
	return question_answers
}

func parseQuestionsFromFile(file *os.File) []QuestionAnswer {
	reader := csv.NewReader(file)
	line, err := reader.Read()
	var question_answers []QuestionAnswer
	for line != nil {
		if err != nil {
			log.Panic("Failed to parse questions file", err)
		}
		question_answers = append(question_answers, QuestionAnswer{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		})
		line, err = reader.Read()
	}
	return question_answers
}
