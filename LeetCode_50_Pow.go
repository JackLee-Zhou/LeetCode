package main

//实现pow(x, n)，即计算 x 的 n 次幂函数（即，xn ）。
//示例 1：
//
//输入：x = 2.00000, n = 10
//输出：1024.00000
//示例 2：
//
//输入：x = 2.10000, n = 3
//输出：9.26100
//示例 3：
//
//输入：x = 2.00000, n = -2
//输出：0.25000
//解释：2-2 = 1/22 = 1/4 = 0.25
//
//提示：
//-100.0 < x < 100.0
//-231 <= n <= 231-1
//-104 <= xn <= 104

// 法一 递归
func myPow(x float64, n int) float64 {
	if n > 0 {
		return div(x, n)
	}
	// 负数次方，取倒数
	return 1.0 / div(x, n)
}
func div(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	p := div(x, n/2)
	// 是偶数
	if n%2 == 0 {
		return p * p
	}
	return p * p * x
}

// 法二
