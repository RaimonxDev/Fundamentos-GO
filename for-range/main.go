package main

import "fmt"

func main() {
	nums := []uint8{2, 4, 6, 8}

	// For range nos devuelve el indice y el valor
	for i, v := range nums {
		fmt.Println("Indice:", i, "valor:", v)
	}

	// omitimos la segunda variable por que no la estamos utilizando
	for i := range nums {
		// multiplicar y cambiar valores
		nums[i] *= 2
	}

}
