package main

import (
	"fmt"
)

var n int
var path []int

func main() {
	fmt.Scan(&n)

	path = []int{}

	back(1)
}

func back(x int) {
	for _, x := range path {
		fmt.Print(x)
		fmt.Print(" ")
	}
	fmt.Println()


	for j := x; j <= n; j++ {
		path = append(path, j)
		back(j + 1)
		path = path[:len(path)-1]
	}
}
