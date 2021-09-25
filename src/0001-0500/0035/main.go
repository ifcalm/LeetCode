package main

//线性判断法
func searchInsert(nums []int, target int) int {
	//如果数组为空，直接返回插入索引值
	if len(nums) == 0 {
		return 0
	}
	//如果 target 小于第一个值，直接返回索引 0
	if target < nums[0] {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		}
		//如果 nums[i] < target < nums[i+1], 并且 i+1 为有效索引
		if i < len(nums)-1 && target > nums[i] && target < nums[i+1] {
			return i + 1
		}
	}
	//target 大于 nums 中所有值，直接插入最后，返回一个新的索引值
	return len(nums)
}

//有序数组，优先考虑二分法
func searchInsert2(nums []int, target int) int {
	//边界值判断
	if len(nums) == 0 || target < nums[0] {
		return 0
	}
	//最大边界值判断
	if target > nums[len(nums)-1] {
		return len(nums)
	}

	start, end := 0, len(nums)-1
	mid := 0

	//需要使用 <=, 不能使用 <, 要保证区间缩小到 1-2 时，也能够正常判断余下的数据
	for start <= end {
		mid = (start + end) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			end = mid - 1
			//如果 nums[mid-1] < target < nums[mid], 返回 mid
			if target > nums[end] {
				return mid
			}
		} else if target > nums[mid] {
			start = mid + 1
		}
	}
	return -1
}
