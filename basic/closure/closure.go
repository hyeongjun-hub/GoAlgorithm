package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// pos, neg는 서로 다른 변수 sum을 가짐
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(i, ":", pos(i), neg(-2*i))
	}
}
