package main

import "fmt"

func main() {

	// Usamos make para crear un map , donde entre corchetes le pasamos el tipo de dato de la key y luego el tipo de dato del valor
	animals := make(map[string]string)
	animals["cat"] = "toto"
	animals["dog"] = "frank"

	fmt.Println(animals)

	// otra manera de crear un map
	fruits := map[string]string{
		"apple":  " :apple ",
		"banana": ":banana",
	}

	fmt.Println(fruits)
	// Eliminar elemento de un map con la funcion delete, indicamos el map y la llave a eliminar
	delete(fruits, "banana")
	fmt.Println(fruits)

	//Obtener un elemento de un map
	fmt.Println(animals["cat"])

	// Si consultamos un valor que no existe go nos devolvera el valor 0 del tipo de dato asignado
	// animal, ok := animals["gorila"]
	// fmt.Println(animal, ok)

	// Ejemplo crear una key que no existe
	if _, ok := animals["gorila"]; !ok {
		animals["gorila"] = ":gorila"
	}
	fmt.Println(animals)

}
