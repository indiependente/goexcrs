package main

import (
	"fmt"
	"strconv"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	return
}

// Walker invokes Walk and close the channel
func Walker(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

func readAndPrintChannel(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var (
		v1, v2   int
		ok1, ok2 bool
	)

	go Walker(t1, ch1)
	go Walker(t2, ch2)

	for {
		select {
		case v1, ok1 = <-ch1:
			if v2, ok2 = <-ch2; v1 != v2 {
				return false
			}
		case v2, ok2 = <-ch2:
			if v1, ok1 = <-ch1; v1 != v2 {
				return false
			}
		}
		if !ok1 || !ok2 {
			break
		}
	}
	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)

	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walker(t1, ch1)
	go Walker(t2, ch2)
	readAndPrintChannel(ch1)
	readAndPrintChannel(ch2)

	fmt.Println("Are they the same tree? " + strconv.FormatBool(Same(t1, t2)))

}
