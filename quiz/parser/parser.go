package parser

import (
	"bufio"
	"os"
	"strings"
)

// ReadLine is a "generator" that reads one line at a time and returns an array split on the comma character
func ReadLine(path string) <-chan []string {
	ch := make(chan []string)
	go func() {
		inFile, _ := os.Open(path)
		defer inFile.Close()
		scanner := bufio.NewScanner(inFile)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			ch <- strings.Split(scanner.Text(), ",")
		}
		close(ch)
	}()
	return ch
}

func ParseCSVFile(path string) (qas [][]string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		qas = append(qas, strings.Split(scanner.Text(), ","))
	}
	return qas
}
