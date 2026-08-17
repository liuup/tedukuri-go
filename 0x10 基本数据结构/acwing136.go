package main
import (
    "fmt"
    "sort"
	"bufio"
	"io/ioutil"
	"os"
)
var bytes []byte
var l, max int

/**
 * Perform Quick Scan of 1 data(with no spaces in-between) on the buffered input
 */
func fastScan() int {
	b := bytes[l]

	// Check if the bytes is not a number
	for (b < 48 || b > 57) && b != 45 {
		l++
		b = bytes[l]
	}
	
	flg := 1
	
	if b == 45 {
	    l++
	    flg = -1
	    b = bytes[l]
	}

	result := 0
	for 47 < b && b < 58 {
		result = (result << 3) + (result << 1) + int(b-48)

		l++
		if l > max {
			return result * flg
		}
		b = bytes[l]
	}
	return result * flg
}
type SortNode struct {
    dat int
    ind int
}
type Node struct {
    dat int
    ind int
    prev *Node
    next *Node
}
type Ans struct {
    min int
    argmin int
}
func abs(x int) int {
    if (x > 0) {
        return x
    }
    return -x
}
func main() {
    bytes, _ = ioutil.ReadAll(bufio.NewReader(os.Stdin))
	max = len(bytes) - 1
    var n int
    n = fastScan()
    aa := make([]*SortNode, n)
    for i := 0; i < n; i++ {
        aa[i] = &SortNode{0, i}
        aa[i].dat = fastScan()
    }
    sort.Slice(aa, func (i, j int) bool {
        return aa[i].dat < aa[j].dat
    })
    a := make([]*Node, n)
    b := make([]*Node, n)
    a[0] = &Node{aa[0].dat, aa[0].ind, nil, nil}
    b[aa[0].ind] = a[0]
    for i := 1; i < n; i++ {
        a[i] = &Node{aa[i].dat, aa[i].ind, a[i-1], nil}
        a[i-1].next = a[i]
        b[aa[i].ind] = a[i]
    }
    ans := make([]*Ans, n)
    for i := n - 1; i > 0; i-- {
        argmin := -1
        sec := -1
        min := 9000000000
        if b[i].prev != nil && abs(b[i].dat - b[i].prev.dat) < min {
            argmin = b[i].prev.ind
            min = abs(b[i].dat - b[i].prev.dat)
            sec = b[i].dat
        }
        if b[i].next != nil && (abs(b[i].dat - b[i].next.dat) < min || abs(b[i].dat - b[i].next.dat) == min && b[i].next.dat < sec) {
            argmin = b[i].next.ind
            min = abs(b[i].dat - b[i].next.dat)
        }
        if b[i].prev != nil {
            b[i].prev.next = b[i].next
        }
        if b[i].next != nil {
            b[i].next.prev = b[i].prev
        }
        ans[i] = &Ans{min, argmin+1}
    }
    for i := 1; i < n; i++ {
        fmt.Println(ans[i].min, ans[i].argmin)
    }
}
