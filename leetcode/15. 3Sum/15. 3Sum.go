package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	length := len(nums)

	for i := 0; i < length-2; i++ {
		// 跳过重复的第一个元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, length-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 跳过重复的左指针元素
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				// 跳过重复的右指针元素
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
