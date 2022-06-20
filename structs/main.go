package main

import "fmt"

func main() {

	// indicamos el type o nombre de la estructura seguido de la palabra struct
	type course struct {
		Name      string
		Professor string
		Duration  uint8
	}

	db := course{
		Name:      "Bases de datos",
		Professor: "Alexys",
		// Sino asignamos un valor, se asigna automaticamente el valor 0 segun el tipo de dato
		// Duration:  25,
	}

	// git := course{"GIT", "BETO", 15}
	// fmt.Printf("%+v\n", db)
	// fmt.Printf("%+v\n", git)

	// acceder un valor en especifico
	// fmt.Printf("%+v\n", db.Name)

	// Crear un puntero a una structu

	p := &db
	// fmt.Printf("%+v\n", p)

	// cambiar el valor desde el puntero
	// * usamos la desreferenciacion y en caso de las estruturas tenemos que colocarlo entre ()
	(*p).Professor = "Juan"

	// O una manera mas facil, GO lo hace por nosotros las desreferenciacion
	// p.Professor = "Ramon"
	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", p)

}
