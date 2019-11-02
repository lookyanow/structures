package list

type node struct {
	val  int
	next *node
}

// List structure
type List struct {
	head *node //Starting of the list
	tail *node //Last element of the list
	len  int   //Number of elements in the list
}

// Returns previous node to this position
func (l List) previousNode(pos int) *node {
	ptr := l.head

	// Out of left bounds
	if pos <= 1 {
		return l.head
	}

	// Out of right bounds
	if pos > l.len {
		return l.tail
	}

	// Iterate through pos-2
	for i := 1; i < pos-1; i++ {
		ptr = ptr.next
	}

	return ptr
}
