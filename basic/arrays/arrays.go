package main

import "fmt"

func main() {
	// 1. 배열 선언과 초기화를 따로
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println("a[0], a[1]:", a[0]+",", a[1])
	fmt.Println("a:", a) // 그냥 배열만 출력 가능!

	// 2. 배열선언과 초기화를 동시에
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("primes:", primes)

}
