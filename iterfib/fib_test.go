package main

import "testing"

func BenchmarkIterativeFib20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		iterativeFib(20) // run the iterFib function b.N times
	}
}

func BenchmarkRecursiveFib20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		recursiveFib(20) // run the iterFib function b.N times
	}
}

func BenchmarkTailRecursiveFib20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tailRecursiveFib(20) // run the iterFib function b.N times
	}
}
