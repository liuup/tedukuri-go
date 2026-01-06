package main
import (
    "container/heap"
    "sort"
    "fmt"
)

type Item struct {
    val      int
	priority int
	index    int
}
type PriorityQueue []*Item
func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	// For GC
	(*pq)[n-1] = nil
	*pq = (*pq)[0 : n-1]
	return item
}

type Cow struct {
    start   int
    end     int
    room    int
    id      int
}

func main() {
    var n int
    fmt.Scanf("%d", &n)
    cows := make([]*Cow, n)
    for i := 0; i < n; i++ {
        cows[i] = &Cow{id: i}
        fmt.Scanf("%d %d", &cows[i].start, &cows[i].end)
    }
    sort.Slice(cows, func(i, j int) bool {
        return cows[i].start < cows[j].start
    })
	pq := make(PriorityQueue, 1, n)
	pq[0] = &Item{val: 1, priority: 0}
	heap.Init(&pq)
	cnt := 1
	for i := 0; i < n; i++ {
	    top := heap.Pop(&pq).(*Item)
	    if top.priority < cows[i].start {
	        top.priority = cows[i].end
	    } else {
	        heap.Push(&pq, top)
	        top = &Item {val: cnt+1, priority: cows[i].end}
	        cnt++
	    }
        cows[i].room = top.val
        heap.Push(&pq, top)
	}
	fmt.Println(len(pq))
    sort.Slice(cows, func(i, j int) bool {
        return cows[i].id < cows[j].id
    })
	for i := 0; i < n; i++ {
        fmt.Println(cows[i].room)
	}
}
