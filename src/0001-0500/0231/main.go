package main

// n&(n-1) 可以消除二进制n的最后一个1
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}
