package extreme_val_

// https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/

// 遍历法
func minArray(numbers []int) int {
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

// 取巧法：由于是旋转数组，所以出现比他小的数就一定是最小的数
func minArrayByBinary(numbers []int) int {
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
			break
		}
	}
	return min
}
