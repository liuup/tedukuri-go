package main
import (
    "fmt"
    "sort"
)


type Point struct {
    x int
    y int
}

func collect_set(a map[int]int) []int {
    keys := make([]int, len(a))
    i := 0
    for k := range a {
        keys[i] = k
        i++
    }
    return keys
}

func main() {
    var c, n int
    fmt.Scanf("%d %d\n", &c, &n)
    pts := make([]Point, n)
    xlist := make(map[int]int)
    ylist := make(map[int]int)
    for i := 0; i < n; i++ {
         fmt.Scanf("%d %d\n", &pts[i].x, &pts[i].y)
         xlist[pts[i].x] = -1
         ylist[pts[i].y] = -1
    }
    xkeys := collect_set(xlist)
    ykeys := collect_set(ylist)
    sort.Ints(xkeys)
    sort.Ints(ykeys)
    for index, v := range xkeys {
        xlist[v] = index
    }
    for index, v := range ykeys {
        ylist[v] = index
    }
    dp := make([][]int, len(xkeys))
    for i := range dp {
        dp[i] = make([]int, len(ykeys))
    }
    for i := 0; i < n; i++ {
        dp[xlist[pts[i].x]][ylist[pts[i].y]] += 1
    }
    for i := 0; i < len(xkeys); i++ {
        for j := 0; j < len(ykeys); j++ {
            if i > 0 {
                dp[i][j] += dp[i-1][j]
            }
            if j > 0 {
                dp[i][j] += dp[i][j-1]
            }
            if i > 0 && j > 0 {
                dp[i][j] -= dp[i-1][j-1]
            }
        }
    }
    l, r := 1, 10000
    for l < r {
        mid := (l + r) >> 1
        possible := false
        for i := 0; i < len(xkeys); i++ {
            for j := 0; j < len(ykeys); j++ {
                iend := xkeys[i] + mid - 1
                jend := ykeys[j] + mid - 1
                iend_ind := sort.SearchInts(xkeys, iend)
                if iend_ind == len(xkeys) || xkeys[iend_ind] > iend {
                    iend_ind = iend_ind - 1
                }
                jend_ind := sort.SearchInts(ykeys, jend)
                if jend_ind == len(ykeys) || ykeys[jend_ind] > jend {
                    jend_ind = jend_ind - 1
                }
                sum := dp[iend_ind][jend_ind]
                if i > 0 {
                    sum -= dp[i-1][jend_ind]
                }
                if j > 0 {
                    sum -= dp[iend_ind][j-1]
                }
                if i > 0 && j > 0 {
                    sum += dp[i-1][j-1]
                }
                if sum >= c {
                    possible = true
                }
            }
        }
        if possible {
            r = mid
        } else {
            l = mid + 1
        }
    }
    fmt.Println(l)
}

