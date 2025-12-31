package main
import "fmt"
type relat struct {
    start int
    end int
}
func main() {
    var d [5005]int
    mp := make(map[relat]bool)
    var n, p, h, m int
    fmt.Scanf("%d %d %d %d", &n, &p, &h, &m)
    for i:=1; i<=m; i++ {
        var a, b int
        fmt.Scanf("%d %d", &a, &b)
        if a > b {
            a, b = b, a
        }
        if mp[relat{a, b}] {
            continue
        }
        mp[relat{a, b}] = true
        d[a+1]--
        d[b]++
    }
    for i:=1; i<=n; i++ {
        d[i] += d[i-1]
        fmt.Println(h + d[i])
    }
}
