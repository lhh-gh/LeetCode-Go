package structures

import "fmt"

type ListNode struct {
	Value    int       //节点值
	Previous *ListNode //前指针
	Next     *ListNode //后指针
}

var ListHead *ListNode

// 添加节点
func addListNode(t *ListNode, v int) int {
	if ListHead == nil {
		t = &ListNode{v, nil, nil}
		ListHead = t
		return 0
	}

	if v == t.Value {
		fmt.Println("节点已存在:", v)
		return -1
	}

	// 如果当前节点下一个节点为空
	if t.Next == nil {
		// 与单链表不同的是每个节点还要维护前驱节点指针
		temp := t
		t.Next = &ListNode{v, temp, nil}
		return -2
	}

	// 如果当前节点下一个节点不为空
	return addListNode(t.Next, v)
}

// 遍历链表
func traverseListNode(t *ListNode) {
	if t == nil {
		fmt.Println("-> 空链表!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

// 查找节点
func lookupListNode(t *ListNode, v int) bool {
	if t == nil {
		fmt.Println("-> 空链表!")
		return false
	}

	for t != nil {
		if t.Value == v {
			return true
		}
		t = t.Next
	}
	return false
}

// 获取链表长度
func sizeListNode(t *ListNode) int {
	if t == nil {
		fmt.Println("-> 空链表!")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}

	return i

}

// 删除节点
func deleteListNode(t *ListNode, v int) bool {
	if t == nil {
		fmt.Println("-> 空链表!")
		return false
	}

	if t.Value == v {
		t = t.Next
		return true
	}

	return deleteListNode(t.Next, v)

}
