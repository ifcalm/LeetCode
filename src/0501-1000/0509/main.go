package main

//递归法
func fib(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

//递归+记忆法
var m = map[int]int{}

func fib2(n int) int {
	if n == 1 || n == 0 {
		return n
	}

	//判断是否已经计算过了
	if v, ok := m[n]; ok {
		return v
	}
	//把计算过的值存储在map中，避免重复计算
	m[n] = fib(n-1) + fib(n-2)
	return m[n]
}

//动态规划，自低向上求值
func fib3(n int) int {
	if n == 1 || n == 0 {
		return n
	}

	p, q, result := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = result
		//状态转移
		result = p + q
	}
	return result
}
