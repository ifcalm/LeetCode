package main

//存储重复计算的数据
var m = map[int]int{}

//递归+记忆法
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	if _, ok := m[n]; ok {
		return m[n]
	}
	m[n] = climbStairs(n-1) + climbStairs(n-2)
	return m[n]
}

//动态规划
func climbStairs2(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}

	return r
}
