package main
import (
    "fmt"
    "sort"
)
type Cow struct {
    min int
    max int
}
type Screen struct {
    spf int
    cnt int
}
func main() {
    var c, l int
    fmt.Scanf("%d %d", &c, &l)
    cows := make([]Cow, c)
    screens := make([]Screen, l)
    for i := 0; i < c; i++ {
        fmt.Scanf("%d %d", &cows[i].min, &cows[i].max)
    }
    for i := 0; i < l; i++ {
        fmt.Scanf("%d %d", &screens[i].spf, &screens[i].cnt)
    }
    sort.Slice(cows, func(i, j int) bool {
		return cows[i].min > cows[j].min
	})
	sort.Slice(screens, func(i, j int) bool {
	    return screens[i].spf > screens[j].spf
	})
	ans := 0
	for _, cow := range(cows) {
	    done := false
	    for ind, screen := range screens {
	        if cow.min <= screen.spf && screen.spf <= cow.max && screen.cnt > 0 {
	            done = true
	            screens[ind].cnt -= 1
	            break
	        }
	    }
	    if done {
	        ans += 1
	    }
	}
	fmt.Println(ans)
}
