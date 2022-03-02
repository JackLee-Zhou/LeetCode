package main

var res [][]int

func Permute(nums []int) [][]int {
	res = [][]int{}
	var path []int
	var st []bool
	for i := 0; i < len(nums); i++ {
		path = append(path, 0)
		st = append(st, false)
	}
	dfs(nums, 0, path, st)
	return res
}

// n 表示层数
func dfs(nums []int, n int, path []int, st []bool) {
	// 表示到底了
	if n == len(nums) {
		// 每次都是一个新的切片
		p := make([]int, len(path))
		// 复制
		copy(p, path)
		// 添加到二维数组中
		res = append(res, p)
		return
	}
	for i := 0; i < len(nums); i++ {
		// 表示没有访问过
		if st[i] == false {
			path[n] = nums[i]
			st[i] = true
			// 向下递归
			dfs(nums, n+1, path, st)
			//	恢复现场
			path[n] = 0
			st[i] = false
		}

	}
}
