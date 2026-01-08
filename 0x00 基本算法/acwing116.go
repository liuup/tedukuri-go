package main
import "fmt"
func main() {
    var state int
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            var c rune
            fmt.Scanf("%c", &c)
            if c == '\n' {
                fmt.Scanf("%c", &c)
            }
            if c == '+' {
                state |= (1 << (4 * i + j))
            }
        }
    }
    minflip := -1
    popmin := 2000000000
    for flip := 0; flip < 65536; flip++ {
        state2 := state
        popcnt := 0
        for flipp := 0; flipp < 16; flipp++ {
            if flip & (1 << flipp) > 0 {
                popcnt++
                for i := 0; i < 4; i++ {
                    state2 ^= (1 << (4 * i + flipp % 4))
                }
                for i := 0; i < 4; i++ {
                    state2 ^= (1 << (flipp / 4 * 4 + i))
                }
                state2 ^= (1 << flipp)
            }
        }
        if (state2 == 0 && popcnt < popmin) {
            popmin = popcnt
            minflip = flip
        }
    }
    fmt.Println(popmin)
    for flipp := 0; flipp < 16; flipp++ {
        if minflip & (1 << flipp) > 0 {
            fmt.Println(flipp/4 + 1, flipp%4 + 1)
        }
    }
}
