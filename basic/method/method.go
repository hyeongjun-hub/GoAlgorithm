package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f MyFloat) power10() {
	f = f * MyFloat(10)
}

func (f *MyFloat) power100() {
	*f = *f * MyFloat(100)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("v.Abs(): ", v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println("f.Abs(): ", f.Abs())

	fmt.Println("기존의 f\t\t\t", f)
	f.power10()
	fmt.Println(f)

	f.power100()
	fmt.Println(f)
}
