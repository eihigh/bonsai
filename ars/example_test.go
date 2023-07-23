package ars_test

import (
	"fmt"

	"github.com/eihigh/bonsai/ars"
)

func Example() {
	v := ars.Vec2{1.2, 3.4}
	fmt.Println(v.ExtendZ(5))

	// Output:
	// {1.2 3.4 5}
}
