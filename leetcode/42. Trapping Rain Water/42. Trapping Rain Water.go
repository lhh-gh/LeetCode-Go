package leetcode

func trap(height []int) int {
	res := 0                        // 存储最终结果（总雨水量）
	left, right := 0, len(height)-1 // 初始化双指针
	maxLeft, maxRight := 0, 0       // 记录左右遍历过程中的最大高度

	// 双指针向中间移动，直到相遇
	for left <= right {
		// 选择较小高度的一侧进行处理（关键策略）
		if height[left] <= height[right] {
			// 处理左指针侧
			if height[left] > maxLeft {
				// 更新左侧最大高度
				maxLeft = height[left]
			} else {
				// 计算当前列的积水量并累加
				res += maxLeft - height[left]
			}
			left++ // 左指针右移
		} else {
			// 处理右指针侧
			if height[right] > maxRight {
				// 更新右侧最大高度
				maxRight = height[right]
			} else {
				// 计算当前列的积水量并累加
				res += maxRight - height[right]
			}
			right-- // 右指针左移
		}
	}
	return res
}
