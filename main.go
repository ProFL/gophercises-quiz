package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ProFL/gophercises-quiz/questions"
)

var questionsFilePath string
var answerTimeoutSeconds int
var shuffleQuestions bool

func init() {
	flag.StringVar(&questionsFilePath, "path", "problems.csv",
		"Path to the problems CSV file, absolute or relative to the CWD")
	flag.IntVar(&answerTimeoutSeconds, "timeout", 30, "the response timeout for a question")
	flag.BoolVar(&shuffleQuestions, "shuffle", false, "wether or not to shuffle the questions")
	flag.Parse()
}

func main() {
	questionAnswers := questions.ReadAndParseQuestions(questionsFilePath, shuffleQuestions)

	correctAnswerCount := 0
	questionCount := len(questionAnswers)

	timer := time.AfterFunc(
		time.Duration(answerTimeoutSeconds)*time.Second,
		func() {
			exit(questionCount, correctAnswerCount, 1)
		},
	)

	for i := 0; i < questionCount; i++ {
		fmt.Printf("%d) ", i+1)
		if questionAnswers[i].AskQuestion() {
			correctAnswerCount += 1
		}
	}

	timer.Stop()
	exit(questionCount, correctAnswerCount, 0)
}

func exit(questionsCount int, correctCount int, exitCode int) {
	fmt.Println()
	fmt.Println("You got", correctCount, "out of", questionsCount, "right")
	os.Exit(exitCode)
}
