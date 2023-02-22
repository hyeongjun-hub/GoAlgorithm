package main

import "fmt"

var num1 int

var num2, num3 int

var num4, num5, str1 = 4, 5, "example"

var (
	i int
	b bool
	s string
)

func main() {
	fmt.Println("1", num1)
	fmt.Println(2, num2, num3)
	fmt.Println("3", num4, num5, str1)

	num6 := 6

	fmt.Println(4, num6)

	fmt.Println(5, i, b, s)

}
