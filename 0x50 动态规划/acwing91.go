package main

import (
    "fmt"
)

const N int = 20
const M int = 1 << N

var f [M][N]int
var weight [N][N]int

var n int

// acwing 91
func main() {
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            fmt.Scan(&weight[i][j])
        }
    }

    for i := 0; i < M; i++ {
        for j := 0; j < N; j++ {
            f[i][j] = 1<<32 - 1
        }
    }

    f[1][0] = 0

    for i := 0; i < 1<<n; i++ { // i表示所有的情况
        for j := 0; j < n; j++ { // j表示走到哪一个点
            if i>>j&1 == 1 {
                for k := 0; k < n; k++ { // k表示走到j这个点之前，以k为终点的最短距离
                    if i>>k&1 == 1 {
                        f[i][j] = min_i(f[i][j], f[i-(1<<j)][k]+weight[k][j])
                    }
                }
            }
        }
    }
    fmt.Println(f[(1<<n)-1][n-1])
}

func min_i(a, b int) int {
    if a < b {
        return a
    }
    return b
}
