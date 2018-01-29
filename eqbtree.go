package main

import (
	"fmt"
	"strconv"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) { // in order visit
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	walkAndClose := func(t *tree.Tree, c chan int) {
		Walk(t, c)
		close(c)
	}
	var (
		v1, v2   int
		ok1, ok2 bool
	)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go walkAndClose(t1, ch1)
	go walkAndClose(t2, ch2)

	for {
		v1, ok1 = <-ch1
		v2, ok2 = <-ch2
		fmt.Println("v1: " + strconv.Itoa(v1) + " ok1: " + strconv.FormatBool(ok1))
		fmt.Println("v2: " + strconv.Itoa(v2) + " ok2: " + strconv.FormatBool(ok2))
		if !ok1 && !ok2 && v1 == v2 { // final case: read from both closed channels (0 value) and both return false flag
			return true
		}
		if ok1 != ok2 || v1 != v2 {
			return false
		}

	}
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)
	same := Same(t1, t2)
	fmt.Println(same)
}
