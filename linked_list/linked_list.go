package linked_list

import (
	"fmt"

)

type LinkedList[K string|int, V string|int] struct {
	Root *NodeKV[K,V]
	Length uint
}


// Constructor
func NewLinkedList[K string|int, V string|int]() LinkedList[K,V]{
	return LinkedList[K, V]{nil,0}
}

// Methods

func (ll *LinkedList[K, V]) Dump(){
	fmt.Printf("Linked list size: %d\n",ll.Length)
	if(ll.Length == 0){
		return
	}
	fmt.Printf("%v -> ",ll.Root.Key)
	var aux *NodeKV[K,V] = ll.Root
	for aux.Next!=nil{
		aux = aux.Next
		fmt.Printf("%v -> ", aux.Key)
	}
	fmt.Printf("\n")
}

func (ll *LinkedList[K,V]) Append(key K, value V){
	node := &NodeKV[K,V]{nil,key,value}

	if ll.Root == nil {
		ll.Root = node
		ll.Length ++
		return
	}

	var aux *NodeKV[K,V] = ll.Root
	for aux.Next!=nil{
		aux=aux.Next
	}
	ll.Length++
	aux.Next = node

}

func (ll *LinkedList[K, V]) Prepend(key K, value V){
	node := &NodeKV[K,V]{nil,key,value}
	if ll.Root == nil {
		ll.Root = node
		ll.Length ++
		return
	}
	node.Next = ll.Root
	ll.Root = node
	ll.Length ++
}

func (ll *LinkedList[K, V]) RemoveFirst(){
	if ll.Root == nil {
		return
	}
	if ll.Root.Next == nil {
		ll.Root = nil 
		ll.Length--
		return
	}
	ll.Root = ll.Root.Next
	ll.Length--
}

func (ll *LinkedList[K, V]) RemoveLast(){
	var aux *NodeKV[K,V] = ll.Root
	if aux == nil {
		return
	}
	if aux.Next == nil {
		ll.Length--
		ll.Root = nil
		return
	}
	var preaux *NodeKV[K,V] = aux
	for aux.Next!=nil{
		preaux = aux
		aux=aux.Next
	}
	preaux.Next = nil
	ll.Length--
}

func (ll *LinkedList[K, V]) GetValueByKey(key K) V{
	var aux *NodeKV[K,V] = ll.Root
	var nil_value V
	if aux == nil {
		return nil_value
	}
	for aux.Next!=nil{
		if(aux.Key == key){
			return aux.Value
		}
		aux=aux.Next
	}
	return nil_value
}