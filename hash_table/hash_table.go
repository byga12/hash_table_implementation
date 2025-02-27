package hashtable

import (
	"errors"
	"fmt"
	ll "hash_table_implementation/linked_list"
)

type HashTable[K string|int, V string|int] struct {
	occupied_entries uint
	table []ll.LinkedList[K,V] 
}

// Constructor
func NewHashTable[K string|int, V string|int]() *HashTable[K,V]{
	// Inicializo hash table con un slice de capacidad 8
	table := make([]ll.LinkedList[K, V], 4)
	ht := HashTable[K,V]{0,table}
	return &ht
}


// Methods
func (ht *HashTable[K, V]) Dump()string{
	var loadFactor float64 = float64(float64(ht.occupied_entries)/float64(len(ht.table)))
	ret := fmt.Sprintf("K/v ocupados: %d Tamaño tabla: %d Load factor: %.2f \n", ht.occupied_entries,len(ht.table),loadFactor)
	for index ,bucket := range ht.table {
		ret+=fmt.Sprintf("%d %s\n",index,bucket.Dump())
	}
	return ret
}

func (ht *HashTable[K, V]) Get(key K) (V,bool){
	var index int = hash_function(key, len(ht.table))
	bucket := &ht.table[index]
	return bucket.GetValueByKey(key)
}

func (ht *HashTable[K, V]) Set(key K, value V){
	var index int = hash_function(key, len(ht.table))

	// Revisar el load factor
	var loadFactor float64 = float64(float64(ht.occupied_entries)/float64(len(ht.table)))
	// Si es mayor a 0.7, redimensionar
	if loadFactor > 0.7 {
		fmt.Println(loadFactor, "Redimensionando")
		ht.resize_and_re_set()
	}


	// Que hay en ese indice?
	bucket := &ht.table[index]
	// Si no hay nada
	if bucket.IsEmpty(){
		bucket.Append(key,value)
		ht.occupied_entries++
		return
	// Si hay algo
	} else {
		// Revisar si la key existe en la lista
		_, exists := bucket.GetValueByKey(key)
		// Si existe la actualizo
		if exists {
			bucket.SetValueByKey(key,value)
		} else {
		// Caso contrario agrego un nuevo nodo al final
			bucket.Append(key,value)
			ht.occupied_entries++
		}
	}
}

// Borra el nodo del bucket que la contenga, si no encuentra nada, devuelve false
func (ht *HashTable[K, V]) Delete(key K) bool{
	var index int = hash_function(key, len(ht.table))
	bucket := &ht.table[index]
	return bucket.DeleteByKey(key)
}

func hash_function(key any, htSize int) int{
	switch k := key.(type) {
		case string:
			return len(k) % htSize
		case int:
			return k % htSize
		default:
			panic(errors.New("Tipo de key no soportado"))
	}
}

func (ht *HashTable[K,V]) resize_and_re_set(){
	var new_size int = len(ht.table)*2
	old_table := ht.table
	new_table := make([]ll.LinkedList[K, V], new_size)
	ht.table = new_table
	ht.occupied_entries=0
	// Recorrer la tabla actual
	for _,ll:= range old_table {
		// Recorrer cada lista
		if ll.Root == nil{
			continue
		}
		for aux:=ll.Root;aux!=nil;aux=aux.Next{
			
			ht.Set(aux.Key,aux.Value)
		}
	}
} 