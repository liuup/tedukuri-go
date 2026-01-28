package main
import (
    "fmt"
    "sort"
)
var n int
var popstates [][]int
func dfs(next int, curstack []int, popstate []int) {
    if next == n && len(curstack) == 0 {
        popstates = append(popstates, popstate)
    } else if next == n {
        ps := append([]int(nil), popstate...)
        dfs(next, curstack[:len(curstack)-1], append(ps, curstack[len(curstack)-1]))
    } else if len(curstack) == 0 {
        cs := append([]int(nil), curstack...)
        dfs(next+1, append(cs, next), popstate)
    } else {
        ps := append([]int(nil), popstate...)
        cs := append([]int(nil), curstack...)
        dfs(next, curstack[:len(curstack)-1], append(ps, curstack[len(curstack)-1]))
        dfs(next+1, append(cs, next), popstate)
    }
}
func lexCompare(a, b []int) int {
    n := len(a)
    if len(b) < n {
        n = len(b)
    }

    for i := 0; i < n; i++ {
        if a[i] < b[i] {
            return -1
        }
        if a[i] > b[i] {
            return 1
        }
    }

    if len(a) < len(b) {
        return -1
    }
    if len(a) > len(b) {
        return 1
    }
    return 0
}
func main() {
    fmt.Scanf("%d", &n)
    popstates = make([][]int, 0, 1048576)
    dfs(0, make([]int, 0, 20), make([]int, 0, 20))
    sort.Slice(popstates, func (i, j int) bool {
        return lexCompare(popstates[i], popstates[j]) < 0
    })
    for i := 0; i < 20 && i < len(popstates); i++ {
        for j := 0; j < len(popstates[i]); j++ {
            fmt.Printf("%d", popstates[i][j]+1)
        }
        fmt.Println()
    }
}
