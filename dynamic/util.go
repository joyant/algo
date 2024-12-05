package dynamic

import "fmt"

func printMatrix(matrix [][]bool) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			t := matrix[i][j]
			if t {
				fmt.Print("T ")
			} else {
				fmt.Print("F ")
			}
		}
		fmt.Println()
	}
}
