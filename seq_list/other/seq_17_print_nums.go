package other

import "math"

// https://leetcode.cn/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/submissions/
func printNumbers(n int) []int {
	nums := []int{}

	max := math.Pow10(n)
	for i := 1; i < int(max); i++ {
		nums = append(nums, i)
	}
	return nums
}
