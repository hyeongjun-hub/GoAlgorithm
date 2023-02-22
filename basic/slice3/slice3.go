package main

import "fmt"

func main() {
	// make()로 가변 길이 배열 만들기
	a := make([]int, 5)
	fmt.Printf("a := make([]int, 5)의\t")
	a[0] = 1
	a[1] = 2
	printSlice(a)

	b := make([]int, 0, 5)
	b = append(b, 1)
	fmt.Printf("b := make([]int, 0, 5)의\t")
	printSlice(b)

	c := b[:2]
	fmt.Printf("c := b[:2]의\t\t")
	printSlice(c)

	d := c[2:5]
	fmt.Printf("d := c[2:5]의\t\t")
	printSlice(d)

	// 한번에 여러개 원소를 추가 가능
	d = append(d, 1, 2)
	fmt.Printf("d = append(d, 1,2,3)후\t")
	printSlice(d)
	// append()는 cap과 len이 같아지면 기존 cap만큼 더 크기를 늘림
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
