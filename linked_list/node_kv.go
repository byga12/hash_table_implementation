package linked_list

type NodeKV[T string|int] struct{
	next *NodeKV[T]
	key T
	value T
}