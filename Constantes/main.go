package main

import "fmt"

func main() {

	// el valor que se asigna no se puede cambiar
	const pi = 3.14
	// marca un error => pi = 3.1416
	fmt.Println(pi)

}
