package in_place

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/

// https://blog.csdn.net/WhereIsHeroFrom/article/details/118887446
func removeDuplicates(nums []int) int {
	count := len(nums)
	// 处理边界条件
	if count == 0 {
		return 0
	}
	if count == 1 {
		return 1
	}

	distinct := 1 // 去重后的总数
	for i := 0; i < count-1; i++ {
		// 重复的跳过
		if nums[i] == nums[i+1] {
			continue
		}
		// 不重复的总数加一
		distinct += 1
		// 判断索引的大小，从而移动位置，实现原地移动
		if distinct-1 <= i {
			nums[distinct-1] = nums[i+1]
		}
	}
	return distinct
}
