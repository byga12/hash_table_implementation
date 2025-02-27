package main

import (
	// Ll "hash_table_implementation/linked_list"
	"fmt"
	Ht "hash_table_implementation/hash_table"
)

func main(){
	// ll := Ll.NewLinkedList[string,string]()
	// ll.Append("nombre", "Ana García")
	// ll.Append("edad", "28")
	// ll.Append("es_estudiante", "NO")
	// ll.Append("promedio", "9.2")
	// ll.Append("direccion", "Calle false 123")
	// ll.Append("temperatura", "25.5C")
	// ll.Append("codigo_postal", "54321")
	// ll.Append("mascota", "golden retriever")
	// ll.Append("cantidad_usuarios", "500")
	// ll.DeleteByKey("mascota")
	// ll.DeleteByKey("nombre")
	// ll.DeleteByKey("es_estudiante")
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
	ht.Delete("cantidad_usuarios")
	ht.Delete("mascota")
	ht.Delete("edad")
	value,exists := ht.Get("temperatura")
	fmt.Println(value,exists)
	fmt.Println(ht.Dump())
}