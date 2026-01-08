package main

import (
	"bufio"
	"os"
	"strconv"
)

type T rune

type Grid struct {
	data   []T
	width  int
	height int
	stride int
	offset int
}

func NewGrid(w, h int) Grid {
	data := make([]T, w*h)
	return Grid{
		data:   data,
		width:  w,
		height: h,
		stride: w,
		offset: 0,
	}
}

func (g Grid) At(x, y int) *T {
	return &g.data[g.offset+y*g.stride+x]
}

func (g Grid) Sub(x, y, w, h int) Grid {
	return Grid{
		data:   g.data,
		width:  w,
		height: h,
		stride: g.stride,
		offset: g.offset + y*g.stride + x,
	}
}

func iPow(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}

func draw(grid Grid) {
	if grid.width == 1 {
		*grid.At(0, 0) = 'X'
	} else {
		newwidth := grid.width / 3
		draw(grid.Sub(0, 0, newwidth, newwidth))
		draw(grid.Sub(0, 2*newwidth, newwidth, newwidth))
		draw(grid.Sub(2*newwidth, 0, newwidth, newwidth))
		draw(grid.Sub(2*newwidth, 2*newwidth, newwidth, newwidth))
		draw(grid.Sub(newwidth, newwidth, newwidth, newwidth))
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	ans := make(map[int]Grid)

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == -1 {
			return
		}

		size := iPow(3, n-1)

		grid, ok := ans[n]
		if !ok {
			grid = NewGrid(size, size)
			for i := 0; i < size; i++ {
				for j := 0; j < size; j++ {
					*grid.At(i, j) = ' '
				}
			}
			draw(grid)
			ans[n] = grid
		}

		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				writer.WriteByte(byte(*grid.At(i, j)))
			}
			writer.WriteByte('\n')
		}
		writer.WriteByte('-')
		writer.WriteByte('\n')
	}
}

