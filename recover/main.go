package main

import "fmt"

func main() {

	divide(10, 5)
	divide(40, 0)
	// Ejecutamos esta linea despues del panico con recover
	divide(60, 5)

}

func divide(param1, param2 int) {

	// Recover se usa con defer y una funcion anonima autoinvocada
	// Se apila la funcion para que se ejecute de ultimo, y cuando entre en panico se ejecuta la funcion y con recover se recupera la app
	defer func() {
		if r := recover(); r != nil {
			// r seria el panico que se ejecuto
			fmt.Println("recuperamos la ejecucion", r)
		}
	}()

	validarDivisor(param2)
	fmt.Println(param1 / param2)

}

func validarDivisor(value int) {
	if value == 0 {
		panic("Se genero un panic")
	}

}
