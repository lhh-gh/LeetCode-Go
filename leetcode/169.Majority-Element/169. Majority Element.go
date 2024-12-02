package leetcode

// majorityElement 函数用于找出数组 nums 中的众数。
// 众数是指在数组中出现次数超过数组长度一半的元素。
func majorityElement(nums []int) int {
	// 初始化 res 为数组的第一个元素，count 用于记录当前元素的出现次数。
	res, count := nums[0], 0
	// 遍历数组 nums。
	for i := 0; i < len(nums); i++ {
		// 如果 count 为 0，则更新 res 为当前遍历到的元素，并将 count 设置为 1。
		if count == 0 {
			res, count = nums[i], 1
		} else {
			// 如果当前元素与 res 相同，则增加 count。
			if nums[i] == res {
				count++
			} else {
				// 如果当前元素与 res 不同，则减少 count。
				count--
			}
		}
	}
	// 返回找到的众数。
	return res
}
