package main

import (
	"math"
	"math/rand"
)

func findPeakElement(nums []int) int {
	// 统计长度
	l := len(nums)
	if l == 1 {
		return 0
	}
	// 前后双指针的移动,必然会在中间相遇
	// TODO 是否可以 i j 移位到相遇就停止
	for i, j := 1, l-2; i <= l-1 && j >= 0; i, j = i+1, j-1 {
		// 单独判断左右是否超过边界
		if i == l-1 && j == 0 {
			if nums[i] > nums[j] {
				return i
			} else {
				return j
			}
		}
		// 左侧峰值判断
		if nums[i] > nums[i+1] && nums[i] > nums[i-1] {
			return i
		}
		// 右侧峰值判断
		if nums[j] > nums[j+1] && nums[j] > nums[j-1] {
			return j
		}
	}
	return 0
}

// 法二 爬坡
func find(nums []int) int {
	n := len(nums)
	idx := rand.Intn(n)
	// 获取指定的值，并考虑边界处理的问题
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}
	// 若不是封顶
	for !(get(idx-1) < get(idx) && get(idx) > get(idx+1)) {
		if get(idx-1) > get(idx) {
			idx--
		} else {
			idx++
		}
	}

	return idx
}
