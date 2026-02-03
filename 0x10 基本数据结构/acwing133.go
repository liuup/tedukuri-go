package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

var bytes []byte
var l, max int

func fastScan() int {
	b := bytes[l]
	for b < '0' || b > '9' {
		l++
		b = bytes[l]
	}
	res := 0
	for b >= '0' && b <= '9' {
		res = res*10 + int(b-'0')
		l++
		if l > max {
			return res
		}
		b = bytes[l]
	}
	return res
}

func main() {
	bytes, _ = ioutil.ReadAll(bufio.NewReader(os.Stdin))
	max = len(bytes) - 1
	ws := bufio.NewWriter(os.Stdout)

	var n, m, q, u, v, t int
	n, m, q, u, v, t =
		fastScan(), fastScan(), fastScan(),
		fastScan(), fastScan(), fastScan()

	var a, b, c []int
	ha, hb, hc := 0, 0, 0

	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = fastScan()
	}
	sort.Ints(A)

	for i := n - 1; i >= 0; i-- {
		a = append(a, A[i])
	}

	offset := 0

	for i := 1; i <= m; i++ {
		current := -1 << 60
		from := 0

		if ha < len(a) && a[ha] > current {
			current = a[ha]
			from = 1
		}
		if hb < len(b) && b[hb] > current {
			current = b[hb]
			from = 2
		}
		if hc < len(c) && c[hc] > current {
			current = c[hc]
			from = 3
		}

		if from == 1 {
			ha++
		} else if from == 2 {
			hb++
		} else {
			hc++
		}

		current += offset

		if i%t == 0 {
			ws.WriteString(strconv.Itoa(current) + " ")
		}
		b = append(b, current*u/v-offset-q)
		c = append(c, current-current*u/v-offset-q)
		offset += q
	}

	B := make([]int, 0, n+m)
	B = append(B, a[ha:]...)
	B = append(B, b[hb:]...)
	B = append(B, c[hc:]...)

	sort.Sort(sort.Reverse(sort.IntSlice(B)))

	ws.WriteString("\n")
	for i := 1; i <= len(B); i++ {
		if i%t == 0 {
			ws.WriteString(strconv.Itoa(B[i-1]+offset) + " ")
		}
	}

	ws.Flush()
}
