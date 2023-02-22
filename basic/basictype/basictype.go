package main

import (
	"fmt"
	"math/cmplx"
)

var (
	i      int = 1<<32 - 1
	f      float64
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const format = "%T(%v)\n"
	fmt.Printf(format, i, i)
	fmt.Printf(format, MaxInt, MaxInt)
	fmt.Printf(format, z, z)

	f = float64(i)
}
