package main
import (
    "fmt"
    "io/ioutil"
    "bufio"
    "os"
)
var l int
var bytes []byte
var s, e, d []int
func max(a int, b int) int {
    if (a > b) {
        return a
    }
    return b
}
func min(a int, b int) int {
    if (a < b) {
        return a
    }
    return b
}
func fastScan() int {
	b := bytes[l]
	for b < 48 || b > 57 {
		l++
		// Next byte
		b = bytes[l]
	}
	result := 0
	for 47 < b && b < 58 {
	    // 8 + 2 = 10
		result = (result << 3) + (result << 1) + int(b-48)
		l++
		if l > len(bytes) - 1 {
			return result
		}
		// Next byte
		b = bytes[l]
	}
	return result
}
func main() {
    bytes, _ = ioutil.ReadAll(bufio.NewReader(os.Stdin))
    t := fastScan()
    for tt := 0; tt < t; tt++ {
        n := fastScan()
        s, e, d = make([]int, n), make([]int, n), make([]int, n)
        for i := 0; i < n; i++ {
            s[i], e[i], d[i] = fastScan(), fastScan(), fastScan()
        }
        l, r := 0, (1 << 31)
        for l < r {
            mid := (l + r) >> 1
            cnt := 0
            for i := 0; i < n; i++ {
                if s[i] > mid {
                    continue
                }
                cnt += max(0, (min(mid, e[i]) - s[i])) / d[i] + 1;
            }
            // fmt.Println(mid, cnt)
            if cnt % 2 == 1 {
                r = mid
            } else {
                l = mid + 1
            }
        }
        if r == 1 << 31 {
            fmt.Println("There's no weakness.")
        } else {
            cnt1 := 0
            for i := 0; i < n; i++ {
                if s[i] > r {
                    continue
                }
                cnt1 += max(0, (min(r, e[i]) - s[i])) / d[i] + 1;
            }
            cnt2 := 0
            for i := 0; i < n; i++ {
                if s[i] > r-1 {
                    continue
                }
                cnt2 += max(0, (min(r-1, e[i]) - s[i])) / d[i] + 1;
            }
            fmt.Println(r, cnt1 - cnt2)
        }
    }
}
