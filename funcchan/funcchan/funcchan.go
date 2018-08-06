package funcchan

import (
	"sync"
)

type fn func() int

func FibonacciFactory(n int) fn {
	return func() int {
		a, b := 1, 1
		for i := 0; i < n; i++ {
			a, b = b, a+b
		}
		return a
	}
}

func GetPopulatedChannel(n int) <-chan fn {
	c := make(chan fn, 10)

	for i := 0; i < cap(c); i++ {
		c <- FibonacciFactory(i)
	}

	close(c)
	return c
}

func ReadAndExecSequential(c <-chan fn) <-chan int {
	fibchan := make(chan int)
	go func() {
		for f := range c {
			fibchan <- f()
		}
		close(fibchan)
	}()
	return fibchan
}

func ReadAndExecConcurrent(c <-chan fn) <-chan int {
	var wg sync.WaitGroup
	fibchan := make(chan int)
	go func() {
		for f := range c {
			fibchan <- f()
		}
		close(fibchan)
	}()
	wg.Wait()
	return fibchan
}
