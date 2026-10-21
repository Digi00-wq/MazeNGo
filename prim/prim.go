package prim

func CreateMaze(size int, a chan [][]int) {
	field := make([][]int, size)

	for i := range size {
		field[i] = make([]int, size)
	}

	for i := range size {
		for j := range size {
			field[i][j] = 0
		}
	}
	a <- field
}
