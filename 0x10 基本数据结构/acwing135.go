package main
import (
    "fmt"
	"bufio"
	"io/ioutil"
	"os"
)
var bytes []byte
var l, max int

func fastScan() int {
    b := bytes[l]


    for (b < 48 || b > 57) && b != 45 {
        l++
        b = bytes[l]
    }

    mult := 1
    if b == 45 {
        mult = -1
        l++
        b = bytes[l]
    }

    result := 0
    for 47 < b && b < 58 {
        result = (result << 3) + (result << 1) + int(b-48)

        l++
        if l > max {
            return result * mult
        }
        b = bytes[l]
    }
    return result * mult
}

type Node struct {
    s int
    i int
}
type Queue struct {
    nodes []*Node
    begin int
}
func (q *Queue) push(a *Node) {
    q.nodes = append(q.nodes, a)
}
func (q *Queue) pop() {
    q.begin++
}
func (q *Queue) front() *Node {
    return q.nodes[q.begin]
}
func (q *Queue) popback() {
    q.nodes = q.nodes[:len(q.nodes)-1]
}
func (q *Queue) back() *Node {
    return q.nodes[len(q.nodes)-1]
}
func (q *Queue) size() int {
    return len(q.nodes) - q.begin
}

func main() {
	bytes, _ = ioutil.ReadAll(bufio.NewReader(os.Stdin))
	max = len(bytes) - 1
    n, m := fastScan(), fastScan()
    a := make([]int, 0, n)
    for i := 0; i < n; i++ {
        a = append(a, fastScan())
        if i > 0 {
            a[len(a)-1] += a[len(a)-2]
        }
    }
    var q Queue
    q.push(&Node{0, -1})
    ans := -2147483649
    for i := 0; i < n; i++ {
        for q.size() > 0 && i - m > q.front().i {
            q.pop()
        }
        if a[i] - q.front().s > ans {
            ans = a[i] - q.front().s
        }
        for q.size() > 0 && q.back().s >= a[i] {
            q.popback()
        }
        q.push(&Node{a[i], i})
    }
    fmt.Println(ans)
}
