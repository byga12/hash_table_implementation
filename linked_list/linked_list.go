package linked_list

type LinkedList[T string|int] struct {
	root *NodeKV[T]
	length uint
}

// Methods
