package main

import (
    "bufio"
    "container/heap"
    "fmt"
    "os"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    var t int
    fmt.Fscan(reader, &t)
    for tt := 0; tt < t; tt++ {
        var ign int
        fmt.Fscan(reader, &ign)
        fmt.Printf("%d ", ign)
        var m int
        fmt.Fscan(reader, &m)
        fmt.Println((m + 1) / 2)
        p := make(MaxHeap, 0)
        q := make(MaxHeap, 0)
        heap.Init(&p)
        heap.Init(&q)
        heap.Push(&p, -2000000000)
        heap.Push(&q, -2000000000)
        for j := 1; j <= m; j++ {
            var x int
            fmt.Fscan(reader, &x)
            top := p[0]
            if top > x {
                heap.Push(&p, x)
            } else {
                heap.Push(&q, -x)
            }
            for len(p) > len(q)+1 {
                heap.Push(&q, -heap.Pop(&p).(int))
            }
            for len(q) > len(p) {
                heap.Push(&p, -heap.Pop(&q).(int))
            }
            if j%2 == 1 {
                fmt.Printf("%d ", p[0])
            }
            if j%20 == 0 || j == m {
                fmt.Println()
            }
        }
    }
}

