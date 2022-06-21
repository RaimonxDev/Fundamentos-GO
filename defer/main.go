package main

import (
	"fmt"
	"os"
)

func main() {
	// Defer aplaza la ejecucion de la funcion
	// defer fmt.Println(3)
	// fmt.Println(1)
	// fmt.Println(2)
	// fmt.Println(2)

	// Si todas las funciones tienen defer entonces las guardara en una pila
	// al finalizar ejecutara la ultima guardada como la primera y la primera guardada sera la ultima
	// defer fmt.Println(3)
	// defer fmt.Println(2)
	// defer fmt.Println(1)

	// a := 5

	// Apesar que se aplaza la ejecucion con defer los valores son evaluados inmediatamente por eso se queda el valor de 5 y no de 10
	// defer fmt.Println("aplazado", a)

	// a = 10
	// fmt.Println(a)

	// Ejemplos para practico defer

	file, err := os.Create("hello.txt")

	if err != nil {
		fmt.Println("ocurrio un error al crear el archivo")
		return
	}

	// Una vez creado el archivo nos aseguramos de cerrarlo. pero al aplicar defer lo dejara de ultimos
	// si no ocurre ningun error entonces se cerrar si o si
	// Si ocurre un error igualmente se ejecutara
	// sirve para limpiar los recursos y asegurarnos que siempre se cierre
	defer file.Close()

	// De file usamos el metodo write y este devuelve el numero de bytes escritos y el error
	// y reasignamos el error
	_, err = file.Write([]byte("hola ramon"))

	if err != nil {
		fmt.Println("error al escribir el archivo ")
		return
	}

}
