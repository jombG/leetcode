package main

func setZeroes(matrix [][]int) {
	columns := len(matrix[0])
	rows := len(matrix)

	firstRowHasZero := false
	firstColumnHasZero := false

	for i := 0; i < columns; i++ {
		if matrix[0][i] == 0 {
			firstRowHasZero = true
			break
		}
	}

	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			firstColumnHasZero = true
			break
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < rows; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < columns; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for i := 1; i < columns; i++ {
		if matrix[0][i] == 0 {
			for j := 1; j < rows; j++ {
				matrix[j][i] = 0
			}
		}
	}

	if firstRowHasZero {
		for i := 0; i < columns; i++ {
			matrix[0][i] = 0
		}
	}

	if firstColumnHasZero {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}
}
