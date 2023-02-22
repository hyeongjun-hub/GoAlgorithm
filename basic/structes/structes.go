package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// 구조체 인스턴스 선언 방법
var (
	// 1. 일반적인 선언 방식
	v1 = Vertex{1, 2}
	// 2. X 값 만을 지정해주고 Y 는 int에 zero value로 설정
	v2 = Vertex{X: 1}
	// 3. X, Y 모두 int에 zero value
	v3 = Vertex{}
)

func main() {
	fmt.Println("v1.X값:", v1.X)
	v1.X = 4
	fmt.Println("v1.X = 4로 바꾼 v1.X값:", v1.X)

	// 4. 구조체 포인터로도 구조체의 값을 바꿀 수 있음
	var p = &v1
	p.X = 10
	fmt.Println("포인터로 바꾼 v1.X값:", v1.X)
}
