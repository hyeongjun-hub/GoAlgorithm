package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	// 1. map 사용
	// map[string] 타입 변수 선언
	var mymap map[string]Vertex
	// make()로 맵 생성
	mymap = make(map[string]Vertex)
	mymap["Bell Labs"] = Vertex{
		40.68433, -74.39667,
	}
	fmt.Println("1. mymap[\"Bell Labs\"]: ", mymap["Bell Labs"])
	// 2. map literal 사용
	var mymap_literal = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39667,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println("2. mymap_literal[\"Bell Labs\"]", mymap_literal["Bell Labs"])

	// 3. map
	var mymap_mine = map[int]int{
		1: 1,
	}
	fmt.Println("3. ", mymap_mine[1])

}
