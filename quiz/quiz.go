package main

import (
	"goexcrs/quiz/parser"
	"goexcrs/quiz/quizzer"
)

func main() {
	path := "problems.csv"
	quizzer := quizzer.New(parser.ParseCSVFile(path))
	quizzer.Quiz()
}
