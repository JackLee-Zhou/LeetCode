package main

// 在不考虑进位的情况下，其无进位加法结果为 a⊕b。
//而所有需要进位的位为a&b，进位后的进位结果为 (a & b) << 1

func getSum(a int, b int) int {
	// 进位为0 则加法过程完成
	for b != 0 {
		//    进位过后的结果
		carry := (a & b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}
