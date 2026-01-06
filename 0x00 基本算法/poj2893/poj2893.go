package main

import (
    "bufio"
    "os"
    "strconv"
)

var tmp []int
var reader = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func nextInt() int {
    reader.Scan()
    n, _ := strconv.Atoi(reader.Text())
    return n
}

func merge(ls []int) int {
    if len(ls) <= 1 {
        return 0
    }
    mid := len(ls) / 2
    leftPairs := merge(ls[:mid])
    rightPairs := merge(ls[mid:])

    merged := tmp[:0]
    i, j := 0, mid
    cnt := 0
    for k := 0; k < len(ls); k++ {
        if j >= len(ls) || (i <= mid-1 && ls[i] < ls[j]) {
            merged = append(merged, ls[i])
            i++
        } else {
            merged = append(merged, ls[j])
            j++
            cnt += mid - i
        }
    }
    copy(ls, merged)
    return cnt + leftPairs + rightPairs
}

func main() {
    defer writer.Flush()
    reader.Split(bufio.ScanWords)
    tmp = make([]int, 300000)

    for {
        n, m := nextInt(), nextInt()
	    if n == 0 && m == 0 {
            break
        }

        size := n * m
        a := make([]int, 0, size)
        row := 0

        for i := 0; i < size; i++ {
            x := nextInt()
            if x != 0 {
                a = append(a, x)
            } else {
                row = i / m
            }
        }

        inv := merge(a)

        if m % 2 == 1 {
            if inv % 2 == 0 {
                writer.WriteString("YES\n")
            } else {
                writer.WriteString("NO\n")
            }
        } else {
            if (inv + n - row) % 2 == 1 {
                writer.WriteString("YES\n")
            } else {
                writer.WriteString("NO\n")
            }
        }
    }
}

