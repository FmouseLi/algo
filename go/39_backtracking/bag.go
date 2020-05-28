package main

import "fmt"

var (
	itemsW = []int{5, 4, 6, 3}
	itemsV = []int{10, 40, 30, 50}
	maxW   = 10
	maxV   = 0
	n      = 4
	r      = []int{}
)

func f(i, cw, cv int, cr []int) {

	if i >= n {
		if cv > maxV {
			maxV = cv
			r = cr
		}
		return
	}

	f(i+1, cw, cv, cr)
	if cw+itemsW[i] <= maxW {
		cr = append(cr, i)
		f(i+1, cw+itemsW[i], cv+itemsV[i], cr)
	}
}

func main() {
	f(0, 0, 0, []int{})
	fmt.Println("最大价值：", maxV)
	fmt.Println("放置的物品下标：", r)
}
