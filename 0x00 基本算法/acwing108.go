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
        if !reader.Scan() {
            break
        }
        n, _ := strconv.Atoi(reader.Text())
        if n == 0 {
            break
        }

        size := n * n
        a := make([]int, 0, size)
        b := make([]int, 0, size)

        for i := 0; i < size; i++ {
            x := nextInt()
            if x != 0 {
                a = append(a, x)
            }
        }
        for i := 0; i < size; i++ {
            x := nextInt()
            if x != 0 {
                b = append(b, x)
            }
        }

        if merge(a)%2 == merge(b)%2 {
            writer.WriteString("TAK\n")
        } else {
            writer.WriteString("NIE\n")
        }
    }
}
