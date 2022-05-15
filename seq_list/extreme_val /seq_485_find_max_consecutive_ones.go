package extreme_val_

// https://leetcode.cn/problems/max-consecutive-ones/
func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	max := 0
	for _, num := range nums {
		if num == 1 {
			count += 1
		} else {
			count = 0
		}
		if count > max {
			max = count
		}
	}
	return max
}
