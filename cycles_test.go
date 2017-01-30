package jsonapi

import (
	"bytes"
	"fmt"
	"testing"
)

type ListNode struct {
	ID     int       `jsonapi:"primary,graph-node"`
	Parent *ListNode `jsonapi:"relation,parent,omitempty"`
	Child  *ListNode `jsonapi:"relation,child,omitempty"`
}

func (n *ListNode) SetChild(child *ListNode) {
	n.Child = child
	child.Parent = n
}

func (n *ListNode) String() string {
	if n.Child == nil {
		return fmt.Sprintf("%d <end>", n.ID)
	}
	return fmt.Sprintf("%d <=> %s", n.ID, n.Child.String())
}

func testList(length int) (root *ListNode) {
	root = &ListNode{ID: 1}

	curNode := root
	for i := 2; i <= length; i++ {
		newNode := &ListNode{ID: i}
		curNode.SetChild(newNode)
		curNode = newNode
	}

	return root
}

func TestMarshal_withCycles(t *testing.T) {
	list := testList(3)
	fmt.Println(list.String())

	out := bytes.NewBuffer(nil)
	if err := MarshalOnePayload(out, list); err != nil {
		t.Fatal(err)
	}

	fmt.Println(out)
}
