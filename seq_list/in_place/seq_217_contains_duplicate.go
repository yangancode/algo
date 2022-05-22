package in_place

import "sort"

// https://leetcode.cn/problems/contains-duplicate/submissions/

// 使用map统计出现的次数
func containsDuplicateByMap(nums []int) bool {
	numStat := make(map[int]int, 0)
	for _, num := range nums {
		if _, ok := numStat[num]; ok {
			return true
		}
		numStat[num] = 1
	}
	return false
}

// 排序后两两比对
func containsDuplicateBySort(nums []int) bool {
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

// 暴力遍历
func containsDuplicateByForeach(nums []int) bool {
	count := len(nums)
	if count == 1 {
		return false
	}

	for i := 0; i < count-1; i++ {
		for j := i + 1; j < count; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
