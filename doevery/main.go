package main

import (
	"fmt"
	"time"
)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func helloworld(t time.Time) {
	fmt.Printf("%v: Hello, World!\n", t)
}

func doStuff() {
	time.Sleep(2 * time.Second)
}

func main() {
	go doEvery(150*time.Millisecond, helloworld)
	doStuff()
}
