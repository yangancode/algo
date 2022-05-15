package seq_list

import "sort"

// https://leetcode.cn/problems/maximum-product-of-two-elements-in-an-array/submissions/
func maxProduct(nums []int) int {
	max := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			vi := nums[i] - 1
			vj := nums[j] - 1
			if vi*vj > max {
				max = vi * vj
			}
		}
	}
	return max
}

func maxProductBySort(nums []int) int {
	sort.Ints(nums)
	count := len(nums)
	return (nums[count-2] - 1) * (nums[count-1] - 1)
}

// https://blog.csdn.net/WhereIsHeroFrom/article/details/119909053
// 可以引申出，如何求TopK
func maxProductByTwoMax(nums []int) int {
	max := 0     // 第一大
	nextMax := 0 // 第二大
	for _, num := range nums {
		// 比最大的数大
		if num > max {
			nextMax = max
			max = num
		} else {
			// 比最大的数小，但是比第二的数大
			if num > nextMax {
				nextMax = num
			}
		}
	}
	return (nextMax - 1) * (max - 1)
}
