package leetcode

import "sort"

// sortRunes 实现 sort.Interface 用于对 []rune 进行排序
type sortRunes []rune

// Less 比较两个字符的顺序（升序排列）
func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap 交换两个字符的位置
func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Len 获取字符切片长度
func (s sortRunes) Len() int {
	return len(s)
}

// groupAnagrams 字母异位词分组函数

func groupAnagrams(strs []string) [][]string {
	// 使用哈希表记录分组，键为排序后的字符序列
	record := make(map[string][]string)

	// 结果切片（预分配内存优化）
	res := make([][]string, 0, len(strs)/3) // 经验值初始化容量

	// 遍历处理每个字符串
	for _, str := range strs {
		// 将字符串转换为可排序的 rune 切片（支持Unicode）
		sByte := []rune(str)

		// 对字符进行排序（关键步骤）
		sort.Sort(sortRunes(sByte))

		// 生成标准化的哈希键
		key := string(sByte)

		// 将原始字符串添加到对应分组
		record[key] = append(record[key], str)
	}

	// 遍历哈希表构建结果集
	for _, group := range record {
		// 对分组进行排序（可选，根据题目要求）
		// sort.Strings(group)
		res = append(res, group)
	}

	return res
}
