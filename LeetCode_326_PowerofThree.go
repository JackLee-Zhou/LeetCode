package main

import "math"

func IsPowerOfThree(n int) bool {
	if n == 0 && n <= 3 {
		return false
	}
	return divPower(n)
}
func divPower(n int) bool {
	if n == 1 {
		return true
	}
	temp := float64(n)
	if _, v := math.Modf(temp / 3); v != 0 {
		//	表示不是整除
		return false
	}
	divPower(n / 3)
	return false
}

/*
	func isPowerOfThree(n int) bool {
    for n > 0 && n%3 == 0 {
        n /= 3
    }
    return n == 1
}
*/
