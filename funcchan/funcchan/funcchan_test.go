package funcchan

import (
	"testing"
)

func BenchmarkReadAndExecSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _ = range ReadAndExecSequential(GetPopulatedChannel(i)) {
		}
	}
}

func BenchmarkReadAndExecConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _ = range ReadAndExecConcurrent(GetPopulatedChannel(i)) {
		}
	}
}
