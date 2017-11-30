package quizzer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Quizzer struct {
	qas [][]string
}

func New(qas [][]string) *Quizzer {
	return &Quizzer{qas}
}

func (q *Quizzer) Quiz() {
	totQuestions := len(q.qas)
	correct := 0
	ansCh := make(chan string)
LL:
	for i, qa := range q.qas {
		timer := time.NewTimer(time.Second * 2)
		askQuestion(i, qa[0])
		go readAnswer(ansCh)
		select {
		case <-timer.C:
			break LL
		case ans := <-ansCh:
			if checkAnswer(ans, qa[1]) {
				correct++
			}
		}
	}
	fmt.Printf("\nYou scored %d out of %d.\n", correct, totQuestions)
}

func askQuestion(i int, q string) {
	fmt.Printf("Problem #%d: %s = ", i+1, q)
}

func readAnswer(ch chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		ch <- scanner.Text()
	}
}

func checkAnswer(a string, realAns string) bool {
	return strings.EqualFold(a, realAns)
}
