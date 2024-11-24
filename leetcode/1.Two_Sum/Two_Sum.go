package leetcode

func twoSum(nums []int, target int) []int {
	//创建一个map m(哈希表) 来存储已经遍历过的数字及其索引
	m := make(map[int]int)
	for i, v := range nums {
		//如果当前元素 v 没有找到配对，则将其及其索引 i 存入哈希表 m 中
		if j, ok := m[target-v]; ok {

			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}
