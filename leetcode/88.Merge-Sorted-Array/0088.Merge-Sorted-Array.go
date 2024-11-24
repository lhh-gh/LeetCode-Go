package leetcode

// 合并两个有序数组的函数
// nums1 是目标数组，m 是 nums1 中已排序部分的长度
// nums2 是源数组，n 是 nums2 的长度
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 初始化 p 为 nums1 的末尾位置，因为我们要逆序填充 nums1
	p := m + n

	// 循环直到 m 或 n 为 0，即一个数组被完全遍历
	for m > 0 && n > 0 {
		// 如果 nums1 的当前元素小于等于 nums2 的当前元素，则将 nums2 的元素放入 nums1
		if nums1[m-1] <= nums2[n-1] {
			nums1[p-1] = nums2[n-1] // 将 nums2 的元素放入 nums1 的末尾
			n--                     // 移动 nums2 的指针
		} else {
			// 否则，将 nums1 的元素放入 nums1 的末尾
			nums1[p-1] = nums1[m-1] // 将 nums1 的元素放入 nums1 的末尾
			m--                     // 移动 nums1 的指针
		}
		p-- // 移动 nums1 的末尾指针
	}

	// 如果 nums2 还有剩余元素，直接将它们复制到 nums1 的前面
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1] // 将 nums2 的剩余元素复制到 nums1
	}
}
