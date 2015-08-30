package linkedlist

import (
	"fmt"
	"strconv"
	"strings"
)

type ForwardNode struct {
	val  int
	next *ForwardNode
}

func (f ForwardNode) String() string {
	result := make([]string, 0)
	var iter *ForwardNode = &f
	for ; iter != nil; iter = iter.next {
		result = append(result, strconv.Itoa(iter.val))
	}
	return fmt.Sprintf("{ %s }", strings.Join(result, ", "))
}

type DoubleNode struct {
	val  int
	next *DoubleNode
	last *DoubleNode
}
