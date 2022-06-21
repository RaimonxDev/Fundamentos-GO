package main

import "fmt"

func main() {

	divide(10, 5)

	// el codigo de detiene y finaliza la app con PANIC y las siguientes lineas ya no se ejecutan.
	divide(40, 0)

	// al tener panic esta linea no se ejecuta
	divide(60, 5)

}

func divide(param1, param2 int) {
	validarDivisor(param2)
	fmt.Println(param1 / param2)

}

func validarDivisor(value int) {
	if value == 0 {
		panic("Se genero un panic")
	}

}
