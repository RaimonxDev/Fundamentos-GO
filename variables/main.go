package main

import "fmt"

func main() {

	//Todas las variables que declaramos las debemos de usar
	var dog string
	dog = "perro"

	// Podemos asignar en una sola linea
	var nameDog string = "frank"
	// Go tambien infiere en los tipo segun su valor al momento de asignar
	var lastNameDog = "francisco"

	// Podemos asignar varias variables en una sola liena
	var esposo, esposa string = "ramon", "meilyn"
	// Go infiere el tipo de la segunda variable
	var espocito, edadEspocito = "ramon", 28

	//Podemos eliminar la palabra var y asignar con :=
	mama, papa := "nora", "jose"

	// Los tipos de datos son estaticos no se pueden cambiar

	// esto marca error =>> edadEspocito = 28

	fmt.Println(dog)
	fmt.Println(nameDog)
	fmt.Println(lastNameDog)
	fmt.Println(esposo)
	fmt.Println(esposa)
	fmt.Println(espocito)
	fmt.Println(edadEspocito)
	fmt.Println(mama)
	fmt.Println(papa)

}
