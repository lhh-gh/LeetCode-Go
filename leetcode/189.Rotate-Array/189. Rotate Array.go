package leetcode

// rotate 函数接受一个整数数组 nums 和一个整数 k，
// 将数组中的元素向右旋转 k 个位置。
func rotate(nums []int, k int) {
	// 首先，对 k 进行取模操作，以防止 k 大于数组长度的情况，
	// 因为旋转数组 len(nums) 次与旋转 0 次的效果是相同的。
	k %= len(nums)

	// reverse 函数用于反转数组中的元素。
	// 首先反转整个数组。
	reverse(nums)

	// 然后反转数组的前 k 个元素。
	reverse(nums[:k])

	// 最后反转数组剩余的部分，即从索引 k 到数组末尾的元素。
	reverse(nums[k:])
}

// reverse 函数用于反转数组 a 中的元素。
func reverse(a []int) {
	// 循环从数组的开始到中间位置，每次交换对应位置的元素。
	for i, n := 0, len(a); i < n/2; i++ {
		// 交换索引 i 和 n-1-i 处的元素。
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
