package main

import "fmt"

func main() {

	animal := "gato"

	switch animal {
	case "perro":
		fmt.Println("es un perro")
	case "gato":
		fmt.Println("es un gato")

	default:
		fmt.Println("no es un animal")

	}

	// Case agrupados
	switch animal {
	case "perro", "gato":
		fmt.Println("es un animal")
	default:
		fmt.Println("no es un animal")
	}
	// Case agrupados con operador
	switch {
	case animal == "perro" || animal == "gato":
		fmt.Println("es un animal domestico")

	case animal == "tigre" || animal == "leon":
		fmt.Println("es un animal salvaje")
	default:
		fmt.Println("no es un animal")
	}
}
