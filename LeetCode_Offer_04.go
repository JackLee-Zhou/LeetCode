package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 特殊条件判断
	if len(matrix) == 0 {
		return false
	}
	// 行列数
	m, n := len(matrix), len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == target {
				return true
			}
			// 表示该列可以继续向下面的行访问
			if matrix[i][j] < target {
				continue
			} else {
				// 右侧的列都会比target元素的值大，所以不访问
				// 重新赋值可访问的列
				n = j
				// 进入下一行判断
				break
			}
		}
	}
	return false
}
