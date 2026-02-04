package main

import (
	"container/list"
	"fmt"
)

func main() {
	var t int
	x := 1
	for {
		_, err := fmt.Scan(&t)
		if err != nil || t == 0 {
			return
		}
		fmt.Printf("Scenario #%d\n", x)
		x++

		big := list.New()
		small := make([]*list.List, t+1)
		for i := 1; i <= t; i++ {
			small[i] = list.New()
		}
		roster := make(map[int]int)

		for i := 1; i <= t; i++ {
			var x int
			fmt.Scan(&x)
			for j := 0; j < x; j++ {
				var y int
				fmt.Scan(&y)
				roster[y] = i
			}
		}

		for {
			var instruction string
			fmt.Scan(&instruction)
			if instruction == "STOP" {
				fmt.Println()
				break
			}
			if instruction == "ENQUEUE" {
				var x int
				fmt.Scan(&x)
				group := roster[x]
				small[group].PushBack(x)
				if small[group].Len() == 1 {
					big.PushBack(group)
				}
			}
			if instruction == "DEQUEUE" {
				frontGroup := big.Front()
				group := frontGroup.Value.(int)
				frontSmall := small[group].Front()
				fmt.Println(frontSmall.Value.(int))
				small[group].Remove(frontSmall)
				if small[group].Len() == 0 {
					big.Remove(frontGroup)
				}
			}
		}
	}
}
