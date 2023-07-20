package coro_test

import (
	"fmt"

	"github.com/eihigh/bonsai/coro"
)

func Example() {
	corun := func(yield func()) {
		fmt.Println("start")
		for i := 0; i < 3; i++ {
			fmt.Println(i)
			yield()
		}
	}
	resume, _ := coro.New0(corun)
	for resume() {
		fmt.Println("---")
	}

	// Output:
	// start
	// 0
	// ---
	// 1
	// ---
	// 2
	// ---
}
