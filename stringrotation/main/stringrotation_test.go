package main

import (
	sr "goexcrs/stringrotation"
	"testing"
)

func BenchmarkIsRotationFF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sr.IsRotationFF("waterbottle", "erbottlewat")
	}
}

func BenchmarkIsRotation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sr.IsRotation("waterbottle", "erbottlewat")
	}
}
