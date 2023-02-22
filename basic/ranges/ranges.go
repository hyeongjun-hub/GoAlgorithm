package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	fmt.Println("1. 일반적인 range")
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	fmt.Println("2. 인덱스가 필요없는 경우")
	for _, v := range pow {
		fmt.Println(v)
	}

	fmt.Println("3. 하나만 해보자")
	for v := range pow {
		fmt.Println(v)
	}
}
