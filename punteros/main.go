package main

import "fmt"

func main() {
	fruit := "manzana"
	// & operador de direccion lo usamos para obtener la direccion de memoria de una variable

	var p *string
	p = &fruit

	fmt.Printf("Tipo: %T, Valor: %v, Direccion: %v\n", fruit, fruit, &fruit)
	fmt.Printf("Variable P =, Tipo: %T, Valor: %v\n", p, p)

	// Podemos ver el valor almacenado en la direccion de memoria con el operador
	// desreferenciacion con *

	fmt.Printf("Variable P =, Tipo: %T, Valor: %v, Valor del dato en memoria: %v\n", p, p, *p)

	// Podemos cambiar el valor de la variable en memoria con el puntero desde su direccion en memoria
	*p = "pera"
	// comprobamos si cambio el valor
	fmt.Println(fruit)
}
