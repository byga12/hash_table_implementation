package hashtable

import (
	"errors"
	ll "hash_table_implementation/linked_list"
)

type HashTable[K string|int, V string|int] struct {
	occupied_entries uint
	table []ll.LinkedList[K,V] 
}

// Constructor
func NewHashTable[K string|int, V string|int]() HashTable[K,V]{
	// Inicializo hash table con un slice de capacidad 8
	table := make([]ll.LinkedList[K, V], 8)
	ht := HashTable[K,V]{0,table}
	return ht
}


// Methods
func (ht HashTable[K, V]) Dump(){

}

func (ht HashTable[K, V]) Set(key K, value V){

}

func Hash_function(key any, htSize int) int{
	switch k := key.(type) {
		case string:
			return len(k) % htSize
		case int:
			return k % htSize
		default:
			panic(errors.New("Tipo de key no soportado"))
	}
}