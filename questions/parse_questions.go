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
	if err != nil {
		log.Panic("There was an error openning the questions file", err.Error())
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			log.Println("Failed to close the questions file", closeErr.Error())
		}
	}()
	questionAnswers := parseQuestionsFromFile(file)
	if shuffleQuestions {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(questionAnswers), func(i, j int) {
			questionAnswers[i], questionAnswers[j] = questionAnswers[j], questionAnswers[i]
		})
	}
	return questionAnswers
}

func parseQuestionsFromFile(file *os.File) []QuestionAnswer {
	reader := csv.NewReader(file)
	line, err := reader.Read()
	var questionAnswers []QuestionAnswer
	for line != nil {
		if err != nil {
			log.Panic("Failed to parse the questions file", err.Error())
		}
		questionAnswers = append(questionAnswers, QuestionAnswer{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		})
		line, err = reader.Read()
	}
	return questionAnswers
}
