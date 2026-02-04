package main
import (
    "fmt"
    "bufio"
    "os"
)
type Node struct {
    weight float64
    total float64
    count int
    origWeight int
    fa int
    child []int
    seq []int
}
func main() {
    reader := bufio.NewReader(os.Stdin)
    var n, R int
    fmt.Fscan(reader, &n, &R)
    R--
    tr := make([]Node, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(reader, &tr[i].weight)
        // fmt.Println(tr[i].weight)
        tr[i].origWeight = int(tr[i].weight)
        tr[i].total = tr[i].weight
        tr[i].seq = append(make([]int, 0, n), i)
        tr[i].child = append(make([]int, 0, n), i)
        tr[i].fa = -1
        tr[i].count = 1
    }
    for i := 0; i < n - 1; i++ {
        var a, b int
        fmt.Fscan(reader, &a, &b)
        a--
        b--
        // fmt.Println("%d", a)
        tr[b].fa = a
        tr[a].child = append(tr[a].child, b)
    }
    for i := 0; i < n - 1; i++ {
        maxn := 0.0
        maxi := -1
        for j := 0; j < n; j++ {
            if tr[j].fa == -1 {
                continue
            }
            if tr[j].weight > maxn {
                maxn = tr[j].weight
                maxi = j
            }
        }
        // fmt.Println(tr[maxi].fa, maxi)
        tr[tr[maxi].fa].total += tr[maxi].total
        tr[tr[maxi].fa].count += tr[maxi].count
        tr[tr[maxi].fa].weight = tr[tr[maxi].fa].total / float64(tr[tr[maxi].fa].count)
        tr[tr[maxi].fa].seq = append(tr[tr[maxi].fa].seq, tr[maxi].seq...)
        for j := 0; j < len(tr[maxi].child); j++ {
            if tr[tr[maxi].child[j]].fa != -1 {
                tr[tr[maxi].child[j]].fa = tr[maxi].fa
                // fmt.Println("parent of", tr[maxi].child[j], "is", tr[maxi].fa)
                tr[tr[tr[maxi].child[j]].fa].child = append(tr[tr[tr[maxi].child[j]].fa].child, tr[maxi].child[j])
            }
        }
        tr[maxi].fa = -1
    }
    // for j := 0; j < n; j++ {
    //     for i := 0; i < len(tr[j].seq); i++ {
    //         fmt.Printf("%d ", tr[j].seq[i])
    //     }
    //     fmt.Println()
    // }
    ans := 0
    for i := 0; i < n; i++ {
        ans += tr[tr[R].seq[i]].origWeight * (i+1)
    }
    fmt.Println(ans)
}
