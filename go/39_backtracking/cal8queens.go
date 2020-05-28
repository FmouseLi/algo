package main

import "fmt"

var (
	result = make([]int, 8)
	count  = 0
)

func Cal8Queens(row int) {
	if row == 8 {
		PrintQueens(result)
		return
	}

	for c := 0; c < 8; c++ {
		if IsOk(row, c) {
			result[row] = c
			Cal8Queens(row + 1)
		}
	}
}

func PrintQueens(result []int) {
	count += 1

	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			if result[r] == c {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func IsOk(row, column int) bool {
	leftup := column - 1
	rightup := column + 1
	for i := row - 1; i >= 0; i-- {
		if result[i] == column {
			return false
		}
		if leftup >= 0 {
			if result[i] == leftup {
				return false
			}
		}
		if rightup < 8 {
			if result[i] == rightup {
				return false
			}
		}
		leftup--
		rightup++
	}
	return true
}

func main() {
	Cal8Queens(0)
	fmt.Println(count)
}
