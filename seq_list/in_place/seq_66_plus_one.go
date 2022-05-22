package in_place

// https://leetcode.cn/problems/plus-one/

// 从后面往前面看，逐一进位
func plusOne(digits []int) []int {
	cnt := len(digits)
	carry := false  // 是否要进位
	for i := cnt - 1; i >= 0; i-- {
		if digits[i]+1 >= 10 {
			digits[i] = 0
			carry = true
		} else {
			digits[i] += 1
			carry = false
			break
		}
	}
	// 这种情况说明进到最后一位，所以在前面还要+1
	if carry {
		nums := []int{1}
		nums = append(nums, digits...)
		return nums
	}
	return digits
}
