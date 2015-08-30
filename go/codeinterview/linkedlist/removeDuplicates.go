package linkedlist

import "fmt"

func RemoveDuplicates(start *ForwardNode) {

	var last *ForwardNode
	var iter *ForwardNode = start
	vals := map[int]bool{}
	for ; iter != nil; iter = iter.next {
		fmt.Printf("Value: %d\n", iter.val)
		if vals[iter.val] {
			last.next = iter.next
		} else {
			vals[iter.val] = true
			last = iter
		}
	}
}
