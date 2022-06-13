package main

import "fmt"

func main() {

	// [n] tipo, declaramos el tamaño del array y luegos su tipo de datos
	// array de 3 posiciones de tipo string
	var food [3]string
	food[0] = " burger"
	food[1] = "coca"
	food[2] = "papa"
	// no podemos cambiar el tamaño del array

	// food[3] ="hot dog" MARCA ERROR

	// Declarar array en 1 sola linea con :=
	// los valores van en llaves {}
	food2 := [3]string{"burger2", "coca2", "papa2"}

	fmt.Println(food)
	fmt.Println(food2)

	// los valores van en llaves
	// SLICES son APUNTADORES de RRAYS

	zoo := [4]string{"leon", "tigre", "delfin", "aguila"}

	felines := zoo[0:2]      // Tambien podemos omitir el 0 y dejarlo vacio entonces GO lo entendera que debe de tomarlo desde el principio quedaria [:2]
	otherAnimals := zoo[2:4] // Tambien podemos omitir el 4 y dejarlo vacio entonces GO lo entendera que debe de tomarlo desde el principio quedaria [2:]
	fmt.Println("ARRAY ORIGINAL:", zoo)
	fmt.Println("felinos", felines)
	fmt.Println("others", otherAnimals)

	/*
		Como los slices son apuntadores de arrays si cambiamos el valor en uno de estos tambien
		cambiamos el valor del array original y todos los demas slices que apunten al array original
		tambien se veran afectados
	*/
	otherAnimals[0] = "vaca"
	fmt.Println("ARRAY ORIGINAL afectado:", zoo)
	fmt.Println("felinos no aceptado", felines)
	fmt.Println("others afectado", otherAnimals)

}
