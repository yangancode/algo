package search

// 二分查找（遍历的方式）
func binarySearch(nums []int, target int) int {
	min := 0
	max := len(nums) - 1 // 为什么要减一：因为索引从0开始
	// 元素找不到的话, min 会增加到和max一致, 如果这时候还是找不到, min会大于max
	for max >= min {
		middle := (min + max) / 2 // 为什么是加起来除以2: 不然范围不会一直缩小
		// 相等
		if nums[middle] == target {
			return middle
		}
		if nums[middle] > target {
			max = middle - 1 // 为什么要减一：因为middle这个值在前面已经判断它不等于目标值，所以不用再把它加进去比较；如果是max=middle，相当于重复计算
		} else {
			min = middle + 1 // 为什么要加一：同理
		}
	}
	return -1
}

// 二分查找法（递归的方式）
func binarySearchByRecursion(nums []int, target int, min, max int) int {
	middle := (min + max) / 2
	if max >= min {
		if nums[middle] == target {
			return middle
		}
		if nums[middle] > target {
			return binarySearchByRecursion(nums, target, min, middle-1)
		} else {
			return binarySearchByRecursion(nums, target, middle+1, max)
		}
	}
	return -1
}