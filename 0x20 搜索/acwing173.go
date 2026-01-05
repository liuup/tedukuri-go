package main

import (
	"bufio"
	"os"
)

func _debug() {
	const eof = 0
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	_i, _n, buf := 0, 0, make([]byte, 1<<12) // 4KB
	rc := func() byte {                      // 读一个字符
		if _i == _n {
			_n, _ = os.Stdin.Read(buf)
			if _n == 0 { // EOF
				return eof
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	ri := func() (x int) { // 读一个整数，支持负数
		neg := false
		b := rc()
		for ; '0' > b || b > '9'; b = rc() {
			// 某些多组数据的题目，不告诉有多少组数据，那么需要额外判断是否读到了 EOF
			if b == eof {
				return
			}
			if b == '-' {
				neg = true
			}
		}
		for ; '0' <= b && b <= '9'; b = rc() {
			x = x*10 + int(b&15)
		}
		if neg {
			return -x
		}
		return
	}
	rs := func() (s []byte) {
		b := rc()
		for ; '0' > b || b > '9'; b = rc() { // 'A' 'Z'
		}
		// for ; 'a' <= b && b <= 'z'; b = rc() { // 'A' 'Z'
		// 	s = append(s, b)
		// }
		for ; '0' <= b && b <= '9'; b = rc() { // 'A' 'Z'
			s = append(s, b)
		}
		return
	}
	_ = []interface{}{rc, ri, rs}

	// 手写输出，适用于有大量（~1e6）输出的场景，CF 上可以再快 60~90ms
	// 使用前 https://codeforces.com/contest/1208/submission/176961129
	// 使用后 https://codeforces.com/contest/1208/submission/176963572
	outS := make([]byte, 0, 1e6*22) // 或者创建一个全局 array _o，然后 outS := _o[:0]（效率几乎一样）
	tmpS := [20]byte{}              // 可根据绝对值的十进制长度的上限调整
	wInt := func(x int) {
		if x == 0 { // 如果保证不为零则去掉
			outS = append(outS, '0')
			return
		}
		if x < 0 { // 如果保证是非负数则去掉
			x = -x
			outS = append(outS, '-')
		}
		p := len(tmpS)
		for ; x > 0; x /= 10 {
			p--
			tmpS[p] = '0' | byte(x%10)
		}
		outS = append(outS, tmpS[p:]...)
	}

	// acwing 173
	n, m := ri(), ri()
	ans := make2dimen(n, m)

	grid := make([][]byte, n)
	for i := 0; i < n; i++ {
		grid[i] = rs()
	}

	// 多源bfs
	q := []node{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				q = append(q, node{i, j})
			}
		}
	}

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		for _, d := range directions {
			dx := cur.x + d[0]
			dy := cur.y + d[1]

			if !isok(grid, dx, dy) || grid[dx][dy] == '1' || ans[dx][dy] > 0 {
				continue
			}
			ans[dx][dy] = ans[cur.x][cur.y] + 1
			q = append(q, node{dx, dy})
		}
	}

	for i := 0; i < n; i++ {
		for _, x := range ans[i] {
			wInt(x)
			outS = append(outS, ' ')
		}
		outS = append(outS, '\n')
	}
	os.Stdout.Write(outS)
}

type node struct {
	x, y int
}

var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} // 四个遍历方向
func isok(grid [][]byte, i, j int) bool { // 判断是否在二维数组越界
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

// 创建指定维度的二维数组
// n: rows; m: cols
func make2dimen(n, m int) (ans [][]int) {
	ans = make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}
	return
}

func main() {
	_debug()
}
