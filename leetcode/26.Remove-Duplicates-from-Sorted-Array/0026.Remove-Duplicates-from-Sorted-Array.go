package leetcode

// 从数组中移除所有重复的元素，并返回新数组的长度
func removeDuplicates(nums []int) int {
	// 如果数组为空，返回0，因为没有元素可以处理
	if len(nums) == 0 {
		return 0
	}
	// 初始化两个指针，last 指向当前确定没有重复的最后一个元素
	// finder 用于查找下一个不重复的元素
	last, finder := 0, 0
	// 遍历数组，直到 finder 指针到达倒数第二个元素
	for last < len(nums)-1 {
		// 使用 finder 指针查找下一个不与 last 指向的元素重复的元素
		for nums[finder] == nums[last] {
			finder++
			// 如果 finder 已经到达数组末尾，说明已经找到所有重复的元素
			if finder == len(nums) {
				return last + 1
			}
		}
		// 将 finder 指向的不重复的元素复制到 last+1 的位置
		nums[last+1] = nums[finder]
		// 移动 last 指针，表示已经处理了一个不重复的元素
		last++
	}
	// 返回处理后数组的长度，即 last+1
	return last + 1
}
