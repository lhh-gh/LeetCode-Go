# [169. Majority Element](https://leetcode.com/problems/majority-element/)


## 题目

给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

## 题目大意

 
给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。你可以假设数组是非空的，并且给定的数组总是存在众数。


## 解题思路

- 题目要求找出数组中出现次数大于 `⌊ n/2 ⌋` 次的数。要求空间复杂度为 O(1)。简单题。
- 这一题利用的算法是 Boyer-Moore Majority Vote Algorithm。[https://www.zhihu.com/question/49973163/answer/235921864](https://www.zhihu.com/question/49973163/answer/235921864)


摩尔投票算法是基于这个事实：每次从序列里选择两个不相同的数字删除掉（或称为“抵消”），最后剩下一个数字或几个相同的数字，就是出现次数大于总数一半的那个

时间复杂度分析
遍历数组：函数中有一个 for 循环，它遍历整个数组 nums。因此，这个循环的时间复杂度是 O(n)，其中 n 是数组 nums 的长度。

条件判断和计数：在循环内部，有一个条件判断和计数操作，这些操作的时间复杂度都是 O(1)。

由于循环是主要的时间消耗操作，因此整个函数的时间复杂度主要由循环决定，即 O(n)。

空间复杂度分析
输入数组：函数的输入是一个数组 nums，这个数组本身就是占用空间的，但是它不是由函数创建的，所以不算作函数的空间复杂度。

额外空间：函数中只使用了两个变量 res 和 count 来存储信息。这两个变量占用的空间是常数级别的，因此空间复杂度是 O(1)。