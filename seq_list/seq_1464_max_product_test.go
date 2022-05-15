package seq_list

import "testing"

func TestMaxProduct(t *testing.T) {
	nums := []int{1, 5, 4, 5}
	max := maxProduct(nums)
	if max != 16 {
		t.Error("failed")
		return
	}
	t.Log("success")
}
