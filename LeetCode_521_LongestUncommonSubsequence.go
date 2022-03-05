package main

import "math"

func findLUSlength(a string, b string) int {
	// 若两个字符相同的话，则没有最长的非公共字符串
	if a == b {
		return -1
	}
	// 否则返回最长的那个就是最长的非公共字符串
	return int(math.Max(float64(len(a)), float64(len(b))))
}
