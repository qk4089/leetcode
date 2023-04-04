package misc

import "fmt"

// 八皇后算法
func findQueens(row int, result []int) {
	if row == 8 {
		printQueens(result)
		return
	}
	for i := 0; i < 8; i++ {
		if isMatch(result, row, i) {
			result[row] = i
			findQueens(row+1, result)
		}
	}
}

func isMatch(result []int, row, column int) bool {
	if row == 0 {
		return true
	}
	for i, leftup, rightup := row-1, column-1, column+1; i >= 0; i, leftup, rightup = i-1, leftup-1, rightup+1 {
		if result[i] == column || (leftup >= 0 && result[i] == leftup) || (rightup < 8 && result[i] == rightup) {
			return false
		}
	}
	return true
}

func printQueens(result []int) {
	for _, v := range result {
		for i := 0; i < len(result); i++ {
			if v == i {
				fmt.Print("* ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
