package leetcode

// removeDuplicates 函数用于移除数组 nums 中的所有重复元素，并返回移除重复元素后数组的长度。
func removeDuplicates(nums []int) int {
	// 初始化 slow 指针，指向数组的开始位置，用于跟踪不包含重复元素的数组的长度。
	slow := 0
	// 使用 for 循环和一个 fast 指针遍历数组 nums。
	for fast, v := range nums {
		// 如果 fast 指针的位置小于 2，或者当前元素与 slow 指针前两个位置的元素不相等，
		// 则说明当前元素不是重复的，可以将其添加到结果数组中。
		if fast < 2 || nums[slow-2] != v {
			// 将当前元素 v 赋值给 nums[slow]，并将 slow 指针向前移动一位。
			nums[slow] = v
			slow++
		}
	}
	// 返回 slow 指针的位置，即移除重复元素后数组的长度。
	return slow
}
