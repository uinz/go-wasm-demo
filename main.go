package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fib(30)
	elapsed := time.Since(start)
	println(fmt.Sprintf("go wasm fib(30) took: %fms", elapsed.Seconds()*1e3))
}

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
