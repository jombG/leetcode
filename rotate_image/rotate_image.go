package main

func rotate(matrix [][]int) {
	weidth := len(matrix[0])
	height := len(matrix)

	for i := 0; i < height; i++ {
		for j := i; j < weidth; j++ {
			n := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = n
		}
	}

	for r := 0; r < height; r++ {
		for j := 0; j < weidth/2; j++ {
			n := matrix[r][j]
			matrix[r][j] = matrix[r][weidth-j-1]
			matrix[r][weidth-j-1] = n
		}
	}
}
