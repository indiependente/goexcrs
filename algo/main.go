package main

import (
	"fmt"
	"goexcrs/algo/sorting"
)

func main() {
	a := []int{3, 4, 2, 1, 0}
	ims := &sorting.IntMergeSorter{}
	ims.Sort(a)
	fmt.Println(a)
}
