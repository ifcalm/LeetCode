package main

//暴力解法，按照版本号从小到大一一判断
func firstBadVersion(n int) int {
	if n == 1 {
		return 1
	}

	for i := 0; i <= n; i++ {
		if isBadVersion(i) {
			return i
		}
	}
	return -1
}

//有序列表，优先考虑二分法
func firstBadVersion2(n int) int {
	if n == 1 {
		return 1
	}

	//二分查找的初始边界值
	start, end := 1, n
	//存储二分的中间值
	mid := 0

	//这里需要使用 <= , 保证区间缩小为 1-2 时, 也能够判断余下的数据
	for start <= end {
		mid = (start + end) / 2
		if isBadVersion(mid) {
			end = mid - 1
			//这里的逻辑：当 mid 判断为 true 时，并且 前一个值 mid-1 判断为 false，那么 mid 就为首次出错的版本
			if !isBadVersion(end) {
				return mid
			}
		} else {
			//缩小搜索空间为原本的一半，注意缩小方向，是左半截还是右半截
			start = mid + 1
		}
	}

	return -1
}

func isBadVersion(version int) bool
