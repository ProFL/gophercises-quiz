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
	question_answers := questions.ReadAndParseQuestions(questionsFilePath, shuffleQuestions)

	correct_answer_count := 0
	question_count := len(question_answers)

	timer := time.AfterFunc(
		time.Duration(answerTimeoutSeconds)*time.Second,
		func() {
			exit(question_count, correct_answer_count, 1)
		},
	)

	for i := 0; i < question_count; i++ {
		fmt.Printf("%d) ", i+1)
		if question_answers[i].AskQuestion() {
			correct_answer_count += 1
		}
	}

	timer.Stop()
	exit(question_count, correct_answer_count, 0)
}

func exit(questions_count int, correct_count int, exitCode int) {
	fmt.Println()
	fmt.Println("You got", correct_count, "out of", questions_count, "right")
	os.Exit(exitCode)
}
