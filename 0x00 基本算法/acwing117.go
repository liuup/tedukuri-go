package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
	    // 1-based index, so +1 size
	    // 6 per stack, as 4+1 may overflow to 5
	    // before removal
		a    [14][6]int
		open [13]int
		ans  int
	)

	readin := func() {
		scanner := bufio.NewScanner(os.Stdin)
		for i := 1; i <= 13; i++ {
			scanner.Scan()
			line := scanner.Text()
			j := 0
			for _, ch := range line {
				if ch != ' ' {
    				j++
    				switch ch {
    				case 'A':
    					a[i][j] = 1
    				case '0':
    					a[i][j] = 10
    				case 'J':
    					a[i][j] = 11
    				case 'Q':
    					a[i][j] = 12
    				case 'K':
    					a[i][j] = 13
    				default:
    					a[i][j] = int(ch - '0')
    				}
    				if j == 4 {
    					break
    				}
				}
			}
		}
	}

	readin()

	for i := 1; i <= 4; i++ {
		now := a[13][i]
		for now != 13 {
			open[now]++
			for j := 5; j >= 2; j-- {
				a[now][j] = a[now][j-1]
			}
			a[now][1] = now
			now = a[now][5]
		}
	}

	for i := 1; i <= 12; i++ {
		if open[i] == 4 {
			ans++
		}
	}

	fmt.Println(ans)
}
