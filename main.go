package main

import (
	// Ll "hash_table_implementation/linked_list"
	"fmt"
	Ht "hash_table_implementation/hash_table"
)

func main(){
	// ll := Ll.NewLinkedList[string,string]()
	// fmt.Println(ll.Dump())

	ht := Ht.NewHashTable[string,string]()
	ht.Set("nombre", "Ana García")
	ht.Set("edad", "28")
	ht.Set("es_estudiante", "NO")
	ht.Set("promedio", "9.2")
	ht.Set("direccion", "Calle false 123")
	ht.Set("temperatura", "25.5C")
	ht.Set("codigo_postal", "54321")
	ht.Set("mascota", "golden retriever")
	ht.Set("cantidad_usuarios", "500")
	fmt.Println(ht.Dump())
}