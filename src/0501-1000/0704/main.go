package main

import "fmt"

//暴力解法，对数组中的值一一判断
func search(nums []int, target int) int {
	//判断特殊情况
	if len(nums) == 0 {
		return -1
	}

	//线性探测，一一对比
	for k, v := range nums {
		//找到目标值，返回索引
		if v == target {
			return k
		}
	}
	//没找到目标值
	return -1
}

//有序数组，可以有限考虑二分法
func search2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	//进行二分的两个边界值
	start, end := 0, len(nums)-1
	//保存二分后的中间值
	mid := 0

	//此处需要 <= , 不能够只有 < , 当需要判断的列表长度 <= 2 时, 随着 mid 的变化, start 和 end 会相等
	for start <= end {
		//每次判断都把区间缩短一半
		mid = (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			//此时 mid 索引值已经判断过了，需要 +1 赋值给 start
			start = mid + 1
		} else if nums[mid] > target {
			//与上面逻辑类似
			end = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(5 / 4)
}
