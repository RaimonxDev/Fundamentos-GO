package main

import (
	"errors"
	"fmt"
)

func main() {

	// content, err := ioutil.ReadFile("./algo.txt")

	// debemos de manejar el error cuando este se presenta, se realiza con un if
	// if err != nil {
	// 	fmt.Printf("Ocurrio un error: %v", err)
	// 	// si hay un error termina la aplicacion y no hay necesidad de un else
	// 	return
	// }
	// fmt.Println(string(content))

	result, err := division(10, 0)

	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	println(result)

}

// Podemos asignar un nombre a los valores que retornamos
func division(dividendo, divisor int) (result int, err error) {

	if divisor == 0 {
		// al tener nombre a los valores que retornamos podemos hacer lo siguiente
		err = errors.New("No se puede dividir por 0")
		// en este caso GO devolvera result con el valor "CERO" por default segun el tipo de dato
		return
		// return 0, errors.New("No se puede dividir por 0")
	}
	// aqui asignamos el resultado y retornamos. Entonces error se devolvera con su valor "CERO" que es nil por el tipo de dato
	result = dividendo / divisor
	return
	// return dividendo / divisor, nil
}
