package extreme_val_

import "testing"

func TestThirdMax(t *testing.T)  {
	nums := []int{2, 2, 3, 1}
	val := thirdMax(nums)
	if val != 1 {
		t.Error("failed")
		return
	}
	t.Log("success")

	nums = []int{5, 2, 2}
	val = thirdMax(nums)
	if val != 5 {
		t.Error("failed")
		return
	}
	t.Log("success")

	nums = []int{1, 2, 2, 5, 3, 5}
	val = thirdMax(nums)
	if val != 2 {
		t.Error("failed")
		return
	}
	t.Log("success")
}
