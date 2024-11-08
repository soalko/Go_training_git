package main

import (
	"fmt"
)

func generatePascalsTriangle(n int) [][]int {
	triangle := make([][]int, n)

	for i := 0; i < n; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}

		triangle[i][i] = 1
	}

	return triangle
}

func printPascalsTriangle(triangle [][]int) {
	for _, row := range triangle {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Print("Введите количество строк треугольника Паскаля: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Количество строк должно быть положительным.")
		return
	}

	triangle := generatePascalsTriangle(n)
	printPascalsTriangle(triangle)
}
