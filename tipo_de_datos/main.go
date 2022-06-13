package main

import "fmt"

func main() {

	a := true

	// printf nos ayuda imprimir en en consola los tipos y valores de las variable que le pasemos

	// Se usan los verbos y se identifican con con %
	// Usamos "T" para el tipo, y "v" para el valor
	// luego le pasamos la variable a consultar en este caso queremos saber el tipo de la variable a
	// y el valor de la variable a

	// NOTA NO SE DEBEN DE USAR COMILLAS SIMPLES '' para los string
	var b string = "letra b"

	// Numeros
	/*
		int8 tiene un tamaño de 8 bits (de -128 a 127).
		int16 tiene un tamaño de 16 bits (de -32768 a 32767).
		int32 tiene un tamaño de 32 bits (de -2147483648 a 2147483647).
		int64 tiene un tamaño de 64 bits (de -9223372036854775808 a 9223372036854775807).
	*/

	/*
		uint8 tiene un tamaño de 8 bits (de 0 a 255).
		uint16 tiene un tamaño de 16 bits (de 0 a 65535).
		uint32 tiene un tamaño de 32 bits (de 0 a 4294967295).
		uint64 tiene un tamaño de 64 bits (de 0 a 18446744073709551615).
	*/

	// Alias
	/*
		El tipo byte es un alias para uint8.
		 Y el tipo rune es un alias para int32.
	*/
	fmt.Printf("Tipo: %T , Valor: %v ", a, a)
	fmt.Println(b)

	// No se puede hacer operaciones con diferentes tipos:
	var d uint8 = 100
	var f uint16 = 20000
	// la siguiente lineas dan error en la variable g
	// // var g = d + f

	// fmt.Println(g)

	// para eso lo solucionamos usando un casting, es decir transformar el tipo de dato para la operacion.
	// PERO NO MODIFICA EL TIPO DE DATO DECLARADO.

	var g = uint16(d) + f

	fmt.Println(g)
	fmt.Printf("Tipo: %T , Valor: %v ", d, d)

	// Operador Blank

}
