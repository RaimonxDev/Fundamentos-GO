package main

import "fmt"

func main() {

	// len() # de elementos de un array
	// cap() # de elementos del array donde apunta, a partir del indice donde
	//se creo el slices
	// food := [5]string{"hotdog", "manzana", "limon", "burger", "pizza"}
	// fruits := food[1:3]

	// fmt.Println("Food", food)
	// fmt.Println("Frutas", fruits)
	// fmt.Println("Frutas", len(fruits)) // 2
	// fmt.Println("Frutas", cap(fruits)) // 4
	/*
		len result 2:  len cuenta la cantidad de elementos que hay en ese array en este caso son 2 manzan y limon
		cap result 4 : como fruits es un slice de food y los slice hacen referencia al array original. entonces con cap contara desde el indice que se creo le nuevo array "fruits" hasta el tamaño del array original food en este caso 4
		porque comenzo del del indice 1 y food tiene 5 indices .

	*/
	// Agregar nuevo elemento con append
	// fruits = append(fruits, "piña", "pera")

	// fmt.Println("Food", food)
	// Tambien modifico el valor en el array original "food"
	// fmt.Println("Frutas", fruits)
	// fmt.Println("Frutas", len(fruits)) // 4
	// fmt.Println("Frutas", cap(fruits)) // 4

	// Que sucede si sobrepasamos la capacidad maxima

	// fruits = append(fruits, "piña", "pera", "banana")
	// fmt.Println("food", food)
	// fmt.Println("Frutas", fruits)
	// fmt.Println("Frutas", len(fruits)) // 5
	// fmt.Println("Frutas", cap(fruits)) // 8

	// Cuando llega a su capicidad maxima crea un nuevo array y lo duplica
	// y pierde la referencia con el array original

	// fruits := []string{"fresa", "arandano"}
	// fruits = append(fruits, "piña")

	//Podemos crear array con la funcion make donde pasamos su tamaño y capacidad
	// fruits := make([]string, 0, 3)

	// fmt.Println("Frutas", len(fruits)) // 0
	// fmt.Println("Frutas", cap(fruits)) // 3

	// valor 0 de los slices su valor es nil o nulo

	var fruits []string
	fmt.Println(fruits)
	fmt.Println(fruits == nil) //true
}
