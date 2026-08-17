package main
import (
	"bufio"
	"io/ioutil"
	"os"
	"fmt"
)

var bytes []byte
var l, max int

/**
 * Perform Quick Scan of 1 data(with no spaces in-between) on the buffered input
 */
func fastScan() int {
	b := bytes[l]

	// Check if the bytes is not a number
	for b < 48 || b > 57 {
		l++
		b = bytes[l]
	}

	result := 0
	for 47 < b && b < 58 {
		result = (result << 3) + (result << 1) + int(b-48)

		l++
		if l > max {
			return result
		}
		b = bytes[l]
	}
	return result
}
var P int
func reverse(numbers []int) {
    for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
        numbers[i], numbers[j] = numbers[j], numbers[i]
    }
}
func hash(a []int) int {
    hasha := 0
    for i := 0; i < 6; i++ {
        hasha += a[i]
        hasha %= P
    }
    hashb := 1
    for i := 0; i < 6; i++ {
        hashb *= a[i]
        hashb %= P
    }
    return (hasha + hashb) % P
}
func sameFlake(a, b []int) bool {
    aa := append(a, a...)
    bb := append(b, b...)
    for offset := 0; offset < 6; offset++ {
        ok := true
        for i := 0; i < 6; i++ {
            if aa[i] != bb[i + offset] {
                ok = false
            } 
        }
        if ok {
            return true
        }
    }
    reverse(aa)
    for offset := 0; offset < 6; offset++ {
        ok := true
        for i := 0; i < 6; i++ {
            if aa[i] != bb[i + offset] {
                ok = false
            } 
        }
        if ok {
            return true
        }
    }
    return false
}
func main() {
	bytes, _ = ioutil.ReadAll(bufio.NewReader(os.Stdin))
	max = len(bytes) - 1
    P = 100003
    grouped := make([][][]int, P)
    for i := 0; i < P; i++ {
        grouped[i] = make([][]int, 0)
    }
    var n int
    n = fastScan()
    snows := make([][]int, n)
    for i := 0; i < n; i++ {
        snows[i] = make([]int, 6)
        for j := 0; j < 6; j++ {
            snows[i][j] = fastScan()
        }
        H := hash(snows[i])
        grouped[H] = append(grouped[H], snows[i])
    }
    ans := 0
    for i := 0; i < n; i++ {
        H := hash(snows[i])
        for j := 0; j < len(grouped[H]); j++ {
            if sameFlake(snows[i], grouped[H][j]) {
                ans += 1
            }
        }
        ans -= 1
    }
    if ans > 0 {
        fmt.Println("Twin snowflakes found.")
    } else {
        fmt.Println("No two snowflakes are alike.")
    }
}
