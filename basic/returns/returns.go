package main

import "fmt"

func divide1(dividend, divisor int) (int, int) {
	var quotient = dividend / divisor
	var reminder = dividend % divisor
	return quotient, reminder
}

func divide2(dividend, divisor int) (quotient, remainder int) {
	quotient = dividend / divisor
	remainder = dividend % divisor
	return
}

func main() {
	var quotient, remainder int
	quotient, remainder = divide1(10, 3)
	fmt.Println("1의 결과:", quotient, remainder)

	quotient, remainder = divide2(10, 3)
	fmt.Println("2의 결과", quotient, remainder)

}
