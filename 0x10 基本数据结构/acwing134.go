package main

import (
    "fmt"
	"bufio"
	"io/ioutil"
	"os"
	"sort"
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

type Sorter struct {
    a int
    index int
    same int
}

func main() {
	bytes, _ = ioutil.ReadAll(bufio.NewReader(os.Stdin))
	max = len(bytes) - 1
	
	n := fastScan()
	a := make([]*Sorter, 0, n)
	for i := 0; i < n; i++ {
	    aa := fastScan()
	    a = append(a, &Sorter{aa, i, 0})
	}
	
	sort.Slice(a, func (i, j int) bool {
	    return a[i].a < a[j].a
	})
	
	for i := 1; i < n; i++ {
	    if a[i].a == a[i-1].a {
	        a[i].same = a[i-1].same
	    } else {
	        a[i].same = a[i-1].same+1
	    }
	   // fmt.Println(a[i].same)
	}
	
	down := true
	tail := 2000000000
	current := make([]*Sorter, 0, n)
	last := 0
	ans := 1
	for i := 0; i < n; i++ {
	   // fmt.Println(a[i].same)
	    if a[i].same == last {
	        current = append(current, a[i])
	    } else {
	        // Process segment current
	        sort.Slice(current, func(i, j int) bool {
	            return current[i].index < current[j].index
	        })
	        if down && current[len(current)-1].index > tail {
	            down = false
	            tail = current[len(current)-1].index
	        } else if !down && current[0].index < tail {
	            down = true
	            ans++
	            tail = current[0].index
	        } else if down && current[len(current)-1].index <= tail {
	            down = true
	            tail = current[0].index
	        } else {
	            down = false
	            tail = current[len(current)-1].index
	        }
	        current = []*Sorter{a[i]}
	        last = a[i].same
	    }
	}
	
    sort.Slice(current, func(i, j int) bool {
        return current[i].index < current[j].index
    })
    if down && current[len(current)-1].index > tail {
        down = false
        tail = current[len(current)-1].index
    } else if !down && current[0].index < tail {
        down = true
        ans++
        tail = current[0].index
    } else if down && current[len(current)-1].index <= tail {
        down = true
        tail = current[0].index
    } else {
        down = false
        tail = current[len(current)-1].index
    }
	fmt.Println(ans)
}
