package main

import (
	"fmt"
	"os"
	"strconv"
)

func emojiToBBCode(e string) string {
	b := make([]byte, 0, 40)
	b = append(b, "[code]"...)
	for _, char := range e {
		//fmt.Printf("character %#U (%d) starts at byte position %d\n", char, char, pos)
		b = append(b, "&#"...)
		b = append(b, strconv.Itoa(int(char))...)
		b = append(b, ";"...)
	}
	b = append(b, "[/code]"...)
	return string(b)
}

func main() {

	emoji := os.Args[1]

	fmt.Println(emojiToBBCode(emoji))

}
