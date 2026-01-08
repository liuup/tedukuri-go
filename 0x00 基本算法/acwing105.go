package main

import (
	"fmt"
	"sort"
	"math"
)

func main() {
	var n, m, t int
	fmt.Scan(&n, &m, &t)

	r := make([]int, n)
	c := make([]int, m)
	rs := make([]int, n)
	cs := make([]int, m)

	for i := 0; i < t; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		x-- // convert to 0-indexed
		y--
		r[x]++
		c[y]++
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			rs[i] = t/n - r[i]
		} else {
			rs[i] = rs[i-1] + t/n - r[i]
		}
	}

	for i := 0; i < m; i++ {
		if i == 0 {
			cs[i] = t/m - c[i]
		} else {
			cs[i] = cs[i-1] + t/m - c[i]
		}
	}

	row, column := false, false
	ans := 0

	if t%n == 0 {
		rsCopy := make([]int, n)
		copy(rsCopy, rs)
		sort.Ints(rsCopy)
		mm := rsCopy[n/2]
		for i := 0; i < n; i++ {
			ans += int(math.Abs(float64(mm - rs[i])))
		}
		row = true
	}

	if t%m == 0 {
		csCopy := make([]int, m)
		copy(csCopy, cs)
		sort.Ints(csCopy)
		mm := csCopy[m/2]
		for i := 0; i < m; i++ {
			ans += int(math.Abs(float64(mm - cs[i])))
		}
		column = true
	}

	if row && column {
		fmt.Print("both ")
	} else if row {
		fmt.Print("row ")
	} else if column {
		fmt.Print("column ")
	} else {
		fmt.Println("impossible")
		return
	}

	fmt.Println(ans)
}

