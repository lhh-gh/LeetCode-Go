# [189. Rotate Array](https://leetcode.com/problems/rotate-array/)

## 题目

给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

## 题目大意

给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

## 解题思路

- 解法二，使用一个额外的数组，先将原数组下标为 i 的元素移动到 `(i+k) mod n` 的位置，再将剩下的元素拷贝回来即可。
- 解法一，由于题目要求不能使用额外的空间，所以本题最佳解法不是解法二。翻转最终态是，末尾 `k mod n` 个元素移动至了数组头部，剩下的元素右移 `k mod n` 个位置至最尾部。确定了最终态以后再变换就很容易。先将数组中所有元素从头到尾翻转一次，尾部的所有元素都到了头部，然后再将 `[0,(k mod n) − 1]` 区间内的元素翻转一次，最后再将 `[k mod n, n − 1]` 区间内的元素翻转一次，即可满足题目要求。

## 代码

```go
package leetcode

// 解法一 时间复杂度 O(n)，空间复杂度 O(1)
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// 解法二 时间复杂度 O(n)，空间复杂度 O(n)
func rotate1(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}
```
## 时间复杂度分析
rotate 函数中，reverse 函数被调用了三次。每次调用reverse函数的时间复杂度是O(N/2)，其中N是数组的长度。因此，总的时间复杂度是O(N)。
由于reverse函数中的循环是固定的，即每次循环都会处理数组的一半元素，所以每次调用reverse的时间复杂度是线性的。
## 空间复杂度分析
这个算法是原地操作的，没有使用额外的数组来存储数据，因此空间复杂度是O(1)