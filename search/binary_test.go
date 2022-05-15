package search

import "testing"

func TestBinary(t *testing.T) {
	nums := []int{1}
	target := 1
	search := binarySearch(nums, target)
	if search != 0 {
		t.Log("failed")
		return
	}
	t.Log("success")

	nums = []int{1}
	target = 2
	search = binarySearch(nums, target)
	if search != -1 {
		t.Log("failed")
		return
	}
	t.Log("success")

	nums = []int{1, 2, 3, 4}
	target = 3
	search = binarySearch(nums, target)
	if search != 2 {
		t.Log("failed")
		return
	}
	t.Log("success")

	nums = []int{1, 2, 3, 4}
	target = 5
	search = binarySearch(nums, target)
	if search != -1 {
		t.Log("failed")
		return
	}
	t.Log("success")

	nums = []int{1, 2, 3, 4, 5}
	target = 1
	search = binarySearch(nums, target)
	if search != 0 {
		t.Log("failed")
		return
	}
	t.Log("success")
}

func TestBinaryByRecursion(t *testing.T) {
	nums := []int{1}
	target := 1
	search := binarySearchByRecursion(nums, target, 0, len(nums)-1)
	if search != 0 {
		t.Log("failed")
		return
	}
	t.Log("success")

	nums = []int{1}
	target = 2
	search = binarySearchByRecursion(nums, target, 0, len(nums)-1)
	if search != -1 {
		t.Log("failed")
		return
	}
	t.Log("success")

	nums = []int{1, 2, 3, 4}
	target = 3
	search = binarySearchByRecursion(nums, target, 0, len(nums)-1)
	if search != 2 {
		t.Log("failed")
		return
	}
	t.Log("success")

	nums = []int{1, 2, 3, 4}
	target = 5
	search = binarySearchByRecursion(nums, target, 0, len(nums)-1)
	if search != -1 {
		t.Log("failed")
		return
	}
	t.Log("success")

	nums = []int{1, 2, 3, 4, 5}
	target = 1
	search = binarySearchByRecursion(nums, target, 0, len(nums)-1)
	if search != 0 {
		t.Log("failed")
		return
	}
	t.Log("success")
}
