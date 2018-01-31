package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(iterativeFib(n))
	fmt.Println(recursiveFib(n))
	fmt.Println(tailRecursiveFib(n))
}

func recursiveFib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}

	return recursiveFib(n-1) + recursiveFib(n-2)
}

func tailRecursiveFib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}

	return _tailRecursiveFib(n, 0, 1)
}
func _tailRecursiveFib(n, a, b int) int {
	if n == 0 {
		return a
	}
	return _tailRecursiveFib(n-1, b, a+b)
}

func iterativeFib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	fib1, fib2 := 0, 1
	for i := 0; i < n-1; i++ {
		fib2, fib1 = fib2+fib1, fib2

	}
	return fib2
}
