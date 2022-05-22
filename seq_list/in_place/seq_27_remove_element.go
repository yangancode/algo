package in_place

// https://leetcode.cn/problems/remove-element/

// 和26题移除重复元素有点像

func removeElement(nums []int, val int) int {
	count := 0
	for idx, num := range nums {
		if num == val {
			continue
		}
		// 数字不相等时，判断当前索引是否大过(最小重复的索引)，是的话移动元素
		// count就是最小重复集合的最后索引，因为相等时跳过了
		if idx >= count {
			nums[count] = num
		}
		// 不包含目标数字的总数
		count += 1
	}
	return count
}
