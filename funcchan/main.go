package main

import (
	"fmt"

	"github.com/indiependente/gophercises/funcchan/funcchan"
)

func main() {
	for i := range funcchan.ReadAndExecSequential(funcchan.GetPopulatedChannel(10)) {
		fmt.Println(i)
	}
}
