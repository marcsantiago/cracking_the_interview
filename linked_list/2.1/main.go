package main

import (
	"strconv"
	"strings"
)

// Node ...
type Node struct {
	Next *Node
	Data int
}

func (n *Node) String() string {
	arr := []string{strconv.Itoa(n.Data)}
	for n.Next != nil {
		n = n.Next
		arr = append(arr, strconv.Itoa(n.Data))
	}
	return strings.Join(arr, ",")
}

// New ...
func New(d int) *Node {
	return &Node{Data: d}
}

// AppendToTail ...
func (n *Node) AppendToTail(d int) {
	end := &Node{Data: d}

	for n.Next != nil {
		n = n.Next
	}
	n.Next = end
	return
}

func deleteNode(head *Node, d int) *Node {
	n := head
	if n.Data == d {
		return head.Next
	}

	for n.Next != nil {
		if n.Next.Data == d {
			n.Next = n.Next.Next
			return head
		}
		n = n.Next
	}
	return head
}

// RemoveDups1 uses a map as a set
func (n *Node) RemoveDups1() {
	// i'm pretty sure this doesn't count as a temp buffer i'm not holding
	// data to move from one stop to another
	m := make(map[int]struct{})
	for n.Next != nil {
		if _, ok := m[n.Data]; ok {
			// pointers is needed to make sure data is overwritten
			*n = *deleteNode(n, n.Data)
		} else {
			m[n.Data] = struct{}{}
			n = n.Next
		}
	}
}

// RemoveDups2 uses a runner
func (n *Node) RemoveDups2() {
	for n.Next != nil {
		runner := n
		for runner.Next != nil {
			if n.Data == runner.Next.Data {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}
		n = n.Next
	}
}

func main() {}
