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

// Private
func (ll *LinkedList[K, V]) IsEmpty()bool{
	if(ll.Root==nil){
		return true
	} else {
		return false
	}
}

// Public
func (ll *LinkedList[K, V]) Dump()string{
	ret := fmt.Sprintf("Linked list size: %d [ ",ll.Length)
	
	if(ll.IsEmpty()){
		ret += "]"
		return ret
	}
	ret += fmt.Sprintf("{%v:%v} -> ",ll.Root.Key, ll.Root.Value)
	var aux *NodeKV[K,V] = ll.Root
	for aux.Next!=nil{
		aux = aux.Next
		ret += fmt.Sprintf("{%v:%v} -> ", aux.Key,aux.Value)
	}
	ret += fmt.Sprintf("]")
	return ret
}

func (ll *LinkedList[K,V]) Append(key K, value V){
	node := &NodeKV[K,V]{nil,key,value}

	if ll.IsEmpty() {
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
	if ll.IsEmpty() {
		ll.Root = node
		ll.Length ++
		return
	}
	node.Next = ll.Root
	ll.Root = node
	ll.Length ++
}

func (ll *LinkedList[K, V]) RemoveFirst(){
	if ll.IsEmpty() {
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
	if ll.IsEmpty() {
		return
	}
	var aux *NodeKV[K,V] = ll.Root
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

// Si no encuentra la key, devuelve el nil_value de tipo V y un false
func (ll *LinkedList[K, V]) GetValueByKey(key K) (V, bool){
	var nil_value V
	if ll.IsEmpty() {
		return nil_value, false
	}
	// var aux *NodeKV[K,V] = ll.Root
	// if(aux.Key == key){
	// 	return aux.Value, true
	// }
	// for aux.Next!=nil{
	// 	aux=aux.Next
	// 	if(aux.Key == key){
	// 		return aux.Value, true
	// 	}
	// }
	for aux:=ll.Root;aux!=nil;aux=aux.Next{
		if(aux.Key == key){
			return aux.Value, true
		}
	}
	return nil_value, false
}

func (ll *LinkedList[K, V]) SetValueByKey(key K, value V){
	if ll.IsEmpty() {
		return
	}
	for aux:=ll.Root;aux!=nil;aux=aux.Next{
		if(aux.Key == key){
			aux.Value=value
		}
	}
}

func (ll *LinkedList[K, V]) DeleteByKey(key K){
	if ll.IsEmpty() {
		return
	}
	if ll.Root.Key == key{
		ll.RemoveFirst()
		return
	}

	preaux := ll.Root
	for aux := preaux.Next ;aux!=nil;{
		if(aux.Key == key){
			preaux.Next = aux.Next
			aux.Next = nil
			ll.Length--
			return
		}
		preaux = aux
		aux = aux.Next
	}
}