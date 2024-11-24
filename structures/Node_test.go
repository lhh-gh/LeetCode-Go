package structures

import (
	"fmt"
	"testing"
)

// 循环链表和单链表的区别是尾节点指向了头节点，从而首尾相连，有点像贪吃蛇，可用于解决「约瑟夫环」
func TestAddNode(t *testing.T) {
	//head := new(Node)
	//var head = new(Node)
	fmt.Println(head)
	head = nil
	// 遍历链表
	traverseNode(head)
	// 添加节点
	addNode(head, 1)
	addNode(head, -1)
	// 再次遍历
	traverseNode(head)
	// 添加更多节点
	addNode(head, 10)
	addNode(head, 5)
	addNode(head, 45)
	// 添加已存在节点
	addNode(head, 5)
	// 再次遍历
	traverseNode(head)

	// 查找已存在节点
	if lookupNode(head, 5) {
		fmt.Println("该节点已存在!")
	} else {
		fmt.Println("该节点不存在!")
	}

	// 查找不存在节点
	if lookupNode(head, -100) {
		fmt.Println("该节点已存在!")
	} else {
		fmt.Println("该节点不存在!")
	}

}
