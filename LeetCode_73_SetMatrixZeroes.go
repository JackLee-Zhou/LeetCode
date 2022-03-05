package main

func setZeroes(matrix [][]int) {
	// 行
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		//	列
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix = doZero(matrix, i, j)
				m--
				n--

			}
		}
	}
}
func doZero(matrix [][]int, m, n int) [][]int {
	//	将 m 行 n列的元素都置为0
	//	 n 行置为0
	for j := 0; j < len(matrix[0]); j++ {
		matrix[m][j] = 0
	}
	// 列置为0
	for i := 0; i < len(matrix); i++ {
		matrix[i][n] = 0
	}
	return matrix
}
