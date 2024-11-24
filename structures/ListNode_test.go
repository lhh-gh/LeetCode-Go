package structures

import (
	"fmt"
	"testing"
)

func TestListNode(t *testing.T) {

	fmt.Println(ListHead)
	ListHead = nil
	addListNode(ListHead, 1)
	addListNode(ListHead, 2)
	addListNode(ListHead, 3)
	lookupListNode(ListHead, 3)

}
