package extreme_val_

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/third-maximum-number/solution/di-san-da-de-shu-by-leetcode-solution-h3sp/

// 先排序 + 找到第三个不同的数
func thirdMax(nums []int) int {
	result := sort.Reverse(sort.IntSlice(nums))
	sort.Sort(result)

	diff := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			diff += 1
		}
		if diff == 3 {
			return nums[i+1]
		}
	}
	return nums[0]
}

// 去重 + 排序
func thirdMaxBySort(nums []int) int {
	var sortSlice []int
	sortMap := make(map[int]string)

	for _, num := range nums {
		if _, ok := sortMap[num]; !ok {
			sortMap[num] = ""
			sortSlice = append(sortSlice, num)
		}
	}
	sort.Ints(sortSlice)
	if len(sortSlice) == 0 {
		return -1
	}
	if len(sortSlice) >= 3 {
		return sortSlice[len(sortSlice)-3]
	}
	return sortSlice[len(sortSlice)-1]
}

func thirdMaxByVal(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num > a {
			c = b
			b = a
			a = num
		// num == a 这种情况要去除，否则不准，相当于一个有序集合有了重复的数字，但没去除
		} else if a > num && num > b {
			c = b
			b = num
		// 同理 num == b 这种情况也要去除
		} else if b > num && num > c {
			c = num
		}
	}
	if c == math.MinInt64 {
		return a
	}
	return c
}
