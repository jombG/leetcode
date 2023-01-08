package nearest_exit_from_entrance_in_maze

func NearestExit(maze [][]byte, entrance []int) int {
	dirs := [][]int{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}}
	queue := make([][]int, 0)
	queue = append(queue, []int{entrance[0], entrance[1], 0})
	for len(queue) > 0 {
		row, col, dis := queue[0][0], queue[0][1], queue[0][2]
		for _, d := range dirs {
			nRow, nCol, nDis := row+d[0], col+d[1], dis+1
			if nRow >= 0 && nRow < len(maze) && nCol >= 0 && nCol < len(maze[0]) && maze[nRow][nCol] == '.' {
				if nRow == 0 || nRow == len(maze)-1 || nCol == 0 || nCol == len(maze[0])-1 {
					return nDis
				}
				maze[row][col] = '+'
				queue = append(queue, []int{nRow, nCol, nDis})
			}
		}
		queue = queue[1:]
	}
	return -1
}

//[".","+","+","+",".",".",".","+","+"],
//[".",".","+",".","+",".","+","+","."],
//[".",".","+",".",".",".",".",".","."],
//[".","+",".",".","+","+",".","+","."],
//[".",".",".",".",".",".",".","+","."],
//[".",".",".",".",".",".",".",".","."],
//[".",".",".","+",".",".",".",".","."],
//[".",".",".",".",".",".",".",".","+"],
//["+",".",".",".","+",".",".",".","."]
//[5,6]
