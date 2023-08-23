package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
    result := make([][]int, 0)
    for i := 0; i < rows; i++ {
        row := []int{}
        for j := 0; j < cols; j++ {
            row = append(row, i*j)
        }
        result = append(result, row)
    }
    return result
}

// dont edit below this line

func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(3, 3)
	test(5, 5)
	test(10, 10)
	test(15, 15)
}