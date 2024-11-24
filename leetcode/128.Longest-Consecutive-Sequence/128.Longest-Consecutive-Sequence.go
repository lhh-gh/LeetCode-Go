package leetcode

// 哈希表法，时间复杂度 O(n)
// 核心思想：通过哈希表记录每个数字所在的连续区间长度
func longestConsecutive(nums []int) int {
	res := 0                // 记录最长连续序列长度
	numMap := map[int]int{} // key: 数值，value: 该数值所在连续区间的长度

	for _, num := range nums {
		if numMap[num] == 0 { // 避免重复处理相同数字
			left := numMap[num-1]   // 获取左邻接区间的长度
			right := numMap[num+1]  // 获取右邻接区间的长度
			sum := left + right + 1 // 计算当前数值连接后的总长度

			numMap[num] = sum   // 更新当前数值的区间长度
			res = max(res, sum) // 更新全局最大值

			// 关键步骤：更新区间边界的长度（实现O(1)时间连接相邻区间）
			numMap[num-left] = sum  // 更新左边界
			numMap[num+right] = sum // 更新右边界
		}
	}
	return res
}

// 辅助函数：返回两个整数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
