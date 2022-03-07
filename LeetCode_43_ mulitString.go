package main

import (
	"fmt"
	"strconv"
)

// 48~57 Ascii 0~9
// 考虑大整数的问题
func multiply(num1 string, num2 string) string {
	var int1 int64
	var int2 int64
	var tmp, rate int64
	int1 = 0
	int2 = 0
	rate = 1
	for i := len(num1) - 1; i >= 0; i-- {
		ans := int64(num1[i]) - 48
		//    fmt.Println(ans)
		tmp = ans * rate
		int1 += tmp
		rate *= 10
	}
	rate = 1
	for i := len(num2) - 1; i >= 0; i-- {
		ans := int64(num2[i]) - 48
		//    fmt.Println(ans)
		tmp = ans * rate
		int2 += tmp
		rate *= 10
	}
	fmt.Println(int1)
	fmt.Println(int2)
	res := int1 * int2
	fmt.Println(res)
	return strconv.FormatInt(res, 10)
}

// Multi 使用模拟竖乘的方法
func Multi(num1 string, num2 string) string {
	// 有一个为0 则返回 0
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	// 返回乘积结果的长度最多不超过 m+n 位
	// 如 9*9 = 81 两个 个位数相乘，最多2位结果
	ansArr := make([]int, m+n)
	var ans string
	// 乘数
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i] - '0')
		// 被乘数
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			ansArr[i+j+1] += x * y
		}
	}
	// 处理结果数组
	for i := m + n - 1; i > 0; i-- {
		// 向上累加要进位的部分
		ansArr[i-1] += ansArr[i] / 10
		// 保留不需要进位的部分
		ansArr[i] = ansArr[i] % 10
	}
	index := 0
	// 舍弃前置的 0
	if ansArr[index] == 0 {
		index++
	}
	// 字符串拼接
	for ; index < m+n; index++ {
		ans += strconv.Itoa(ansArr[index])
	}
	return ans
}
