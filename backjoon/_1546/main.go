package main

import "fmt"

func main() {
	var num int
	fmt.Scanf("%d", &num)
	var arr = make([]int, num)
	max := 0
	sum := 0

	for i := 0; i < num; i++ {
		fmt.Scanf("%d", &arr[i])
		sum += arr[i]
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println(float64(sum) / float64(num) / float64(max) * 100.0)
}
