package linkedlist

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	root := ForwardNode{val: 0}
	e1 := ForwardNode{val: 1}
	e2 := ForwardNode{val: 0}
	e3 := ForwardNode{val: 0}
	e4 := ForwardNode{val: 2}
	e5 := ForwardNode{val: 3}
	e6 := ForwardNode{val: 1}
	e7 := ForwardNode{val: 0}
	root.next = &e1
	e1.next = &e2
	e2.next = &e3
	e3.next = &e4
	e4.next = &e5
	e5.next = &e6
	e6.next = &e7

	fmt.Println("Before: %s", root)
	RemoveDuplicates(&root)
	fmt.Println("After: %s", root)
}
