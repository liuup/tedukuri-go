package main

import (
    "fmt"
    "os"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	var n int
	fmt.Fscan(os.Stdin, &n)
	a := make([][]int, n)
	s := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		s[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(os.Stdin, &a[i][j])
			if i == 0 {
				s[i][j] = a[i][j]
			} else {
				s[i][j] = s[i-1][j] + a[i][j]
			}
		}
	}
	ans := int(-2e9)
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			last := int(-2e9)
			for j := 0; j < n; j++ {
				cur := s[i2][j]
				if i1 > 0 {
					cur -= s[i1-1][j]
				}
				dp := max(last+cur, cur)
				ans = max(ans, dp)
				last = dp
			}
		}
	}
	fmt.Println(ans)
}

