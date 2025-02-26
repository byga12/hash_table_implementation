package linked_list

type NodeKV[K string|int, V string|int] struct{
	Next *NodeKV[K,V]
	Key K
	Value V
}