package main

import "fmt"

func main() {
	// Aritmeticos +, - , *, /
	var a = 4 + 2*3
	fmt.Println(a)
	var b = 10

	b += a
	fmt.Println(b)

	var c = 3
	// post-incremento ++  , post-decremento --
	c++ // c--
	fmt.Println(c)
	// no valido
	// fmt.Println(c++)

	// Comparacion <,>, ==, != , >= , <=
	fmt.Println(4 < 3)

	// Logicos && , ||
	var age = 30

	fmt.Println(age >= 18 && age <= 80)

	// Unario negamos el valor booleano
	fmt.Println(!(4 == 4))

}
