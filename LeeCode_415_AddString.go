package main

import "strconv"

// input: 11+123
// output: 1304 want 134
// 中间多了一个0
func addStrings(num1 string, num2 string) string {
	if num1 == "0" && num2 == "0" {
		return "0"
	}
	return Sub(num1, num2)
}
func Sub(a, b string) string {
	m, n := len(a), len(b)
	i, j := m-1, n-1
	res := make([]int, m+n)
	ans := ""
	for i >= 0 && j >= 0 {
		x := int(a[i] - '0')
		y := int(b[j] - '0')
		res[i+j+1] += x + y
		i--
		j--
	}
	// 只有一个会执行
	for i >= 0 {
		x := int(a[i] - '0')
		res[i+1] += x
		i--
	}
	for j >= 0 {
		y := int(b[j] - '0')
		res[j+1] += y
		j--
	}
	for i := len(res) - 1; i > 0; i-- {
		// 向前进位
		res[i-1] += res[i] / 10
		res[i] = res[i] % 10
	}
	index := 0
	for res[index] == 0 {
		index++
	}
	// 从index 到末尾就是结果
	for i := index; i < len(res); i++ {
		ans += strconv.Itoa(res[i])
	}
	return ans
}

// Add 模拟竖式加法
func Add(num1, num2 string) string {
	// 是否进位标识符
	add := 0
	ans := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		// 若最后有进位 则 x y = 0  add =1
		res := x + y + add
		// 从 ans 的头部开始拼接
		ans = strconv.Itoa(res%10) + ans
		// 是否进位
		add = res / 10
	}
	return ans
}
