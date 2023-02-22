package main

import "fmt"

// Pi1 상수도 선언과 동시에 초기화하면 타입을 지정하지 않아도 됨
const Pi1 float32 = 3.14
const Pi2 = 3.14

// 괄호로 묶으면 상수 키워드를 한 번만 명시
const (
	BigConst   = 1 << 100
	SmallConst = BigConst >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("needInt(SmallConst):", needInt(SmallConst))
	fmt.Println("needFloat(SmallConst):", needFloat(SmallConst))
	fmt.Println("needFloat(BigConst)", needFloat(BigConst))
}
