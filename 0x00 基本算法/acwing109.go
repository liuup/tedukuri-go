package main

import (
    "bufio"
    "os"
    "strconv"
    "sort"
    // "fmt"
)

var tmp []int
var curray []int

func iMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}


func reverseSlice(input []int) []int {
    reversed := make([]int, len(input))
    for i, v := range input {
        reversed[len(input)-1-i] = v
    }
    return reversed
}
func validate(st []int, m, t int) bool {
    m = iMin(len(st)/2, m)
    l1, l2 := reverseSlice(st[:m]), st[len(st)-m:]
    rmse := 0
    for i := 0; i < len(l1); i++ {
        rmse += (l1[i] - l2[i]) * (l1[i] - l2[i])
    }
    return rmse <= t
}

// Merge b into curray and return a new array.
func merge(b []int) []int {
    a := curray
    var tomerge = tmp[:0]
    i, j := 0, 0
    n := len(a) + len(b)
    for k := 0; k < n; k++ {
        if j == len(b) || (i < len(a) && a[i] < b[j]) {
            tomerge = append(tomerge, a[i])
            i++
        } else {
            tomerge = append(tomerge, b[j])
            j++
        }
    }
    return tomerge
}

// nextInt reads the next integer from the buffered reader.
func nextInt(reader *bufio.Reader) int {
    var b []byte
    for {
        ch, err := reader.ReadByte()
        if err != nil {
            break
        }
        if ch >= '0' && ch <= '9' {
            b = append(b, ch)
        } else if ch == '-' && len(b) == 0 {
            b = append(b, ch)
        } else if len(b) > 0 {
            break
        }
    }
    n, _ := strconv.Atoi(string(b))
    return n
}

func sortedCopy(a []int) []int {
    b := make([]int, len(a))  // allocate new slice
    copy(b, a)                 // copy contents
    sort.Ints(b)               // sort the copy
    return b
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    k := nextInt(reader)

    for i := 0; i < k; i++ {
        n := nextInt(reader)
        m := nextInt(reader)
        t := nextInt(reader)

        a := make([]int, n)
        tmp = make([]int, n)
        for j := 0; j < n; j++ {
            a[j] = nextInt(reader)
        }

        start := 0
        slcs := 0
        for start < len(a) {
            curray = curray[:0]
            curray = append(curray, a[start])
            incr, end := 1, start
            stmerge := start+1
            for incr > 0 {
                if stmerge >= n {
                    incr = 0
                    break
                }
                if end+incr+1 > n {
                    incr /= 2
                    continue
                }
                attempt := merge(sortedCopy(a[stmerge : end+incr+1]))
                // fmt.Printf("%v + %v to %v\n", curray, a[stmerge : end+incr+1], attempt)
                if validate(attempt, m, t) {
                    end += incr
                    incr *= 2
                    stmerge = end + 1
                    curray = curray[:0]
                    curray = append(curray, attempt...)
                    // fmt.Printf("Curray to %v\n", curray)
                } else {
                    incr /= 2
                }
            }
            slcs += 1
            start = end + 1
        }

        writer.WriteString(strconv.Itoa(slcs) + "\n")
    }
}
