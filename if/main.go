package main

import "fmt"

func main() {
	// dog := "frank"
	// if dog == "frank" {
	// 	fmt.Println("es tu mascota")
	// } else if dog == "sasha" {
	// 	fmt.Println("amiga de frank")
	// } else {
	// 	fmt.Println("Ups no se quien es")
	// }

	// Cuando declaramos en el if, esta solo existe en el scope de if
	if emoji := "🐕"; emoji == "🐕" {
		fmt.Println("es tu gato")
	} else if emoji == "🔨" {
		fmt.Println("es un martillo")
	} else {
		fmt.Println("Ups no se que es")
	}
	// fmt.Println(emoji) // undefined

}
