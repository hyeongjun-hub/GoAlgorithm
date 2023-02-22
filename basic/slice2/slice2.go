package main

import "fmt"

func main() {
	fmt.Println("1. 슬라이스 리터럴 선언")
	// 1. 기본형 슬라이스 리터럴
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("기본형 슬라이스 리터럴:", q)
	// 2. 구조체 슬라이스 리터럴
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false,},
		{5, true,},
		{7, true,},
		{11, false,},
		{13, true,},
	}
	fmt.Println("구조체 슬라이스 리터럴:", s)

	fmt.Println("2. 슬라이스를 슬라이스")
	q = q[:2]
	fmt.Println("q[:2]:", q)

	q = q[1:]
	fmt.Println("q[1:]:", q)
}