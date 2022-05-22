package other

// https://leetcode.cn/problems/group-the-people-given-the-group-size-they-belong-to/submissions/

func groupThePeople(groupSizes []int) [][]int {
	sizeMap := make(map[int][]int, 0)
	// groupSizes就是分组号组成的数组，所以一开始先将一个map存储分组号对应的数字（索引）
	for idx, group := range groupSizes {
		// group 是分组号
		if _, ok := sizeMap[group]; ok {
			sizeMap[group] = append(sizeMap[group], idx)
		} else {
			sizeMap[group] = []int{idx}
		}
	}
	// 分好组后，再根据分组号切割
	var userGroup [][]int
	for key, groups := range sizeMap {
		// key可以理解为步长
		if key == len(groups) {
			userGroup = append(userGroup, groups)
		} else {
			// 分组号内的元素大于分组号（步长），就需要分割
			for i := 0; i < len(groups); i += key {
				userGroup = append(userGroup, groups[i:i+key])
			}
		}
	}
	return userGroup
}

// 优化版: 减少遍历
func groupThePeopleOptimize(groupSizes []int) [][]int {
	var userGroup [][]int
	sizeMap := make(map[int][]int, 0)

	for idx, group := range groupSizes {
		// group 是分组号
		if _, ok := sizeMap[group]; ok {
			sizeMap[group] = append(sizeMap[group], idx)
		} else {
			sizeMap[group] = []int{idx}
		}
		if len(sizeMap[group]) == group {
			userGroup = append(userGroup, sizeMap[group])
			delete(sizeMap, group)
		}
	}
	return userGroup
}
