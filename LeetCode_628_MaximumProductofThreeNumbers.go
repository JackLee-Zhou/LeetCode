package main

import "sort"

func maximumProduct(nums []int) int {
	// 排序过后，数组必定以0为分界线
	sort.Ints(nums)
	m, n := 0, 0
	// 统计正负元素的个数
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			m++
		}
		if nums[i] < 0 {
			n++
		}
	}
	a, b := 0, 0
	// 最小的两个负数相乘为正数，在*一个最右侧的正数
	if n >= 2 {
		a = nums[0] * nums[1] * nums[len(nums)-1]
	}
	// × 3个最右侧的正数
	b = nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
	if a > b {
		return a
	}
	return b
}
