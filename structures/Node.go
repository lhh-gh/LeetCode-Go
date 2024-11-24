package structures

import "fmt"

// 定义节点结构体    单链表
type Node struct {
	Value int   //
	Next  *Node //
}

// 初始化头节点
var head = new(Node)

// 添加节点
func addNode(t *Node, v int) int {

	if head == nil {
		t = &Node{v, nil}
		head = t
		return 0
	}

	if v == t.Value {
		fmt.Println("节点已存在")
		return 1
	}
	// 如果当前节点下一个节点为空
	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}

	// 如果当前节点下一个节点不为空
	return addNode(t.Next, v)
}

// 遍历链表
func traverseNode(t *Node) {
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
func lookupNode(t *Node, v int) bool {
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
func size(t *Node) int {
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
func deleteNode(t *Node, v int) bool {
	if t == nil {
		fmt.Println("-> 空链表!")
		return false
	}

	if t.Value == v {
		t = t.Next
		return true
	}

	for t.Next != nil {
		if t.Next.Value == v {
			t.Next = t.Next.Next
			return true
		}
		t = t.Next
	}
	return false
}
