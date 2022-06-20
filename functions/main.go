package main

import (
	"fmt"
	"strings"
)

func main() {
	hello("ramon", "martinez")
	name := "ramon"

	// tenemos que crear un puntero a la variable name
	change(&name)
	fmt.Println(name)

	total := sum(5, 6)

	fmt.Println(total)

	minusculas, mayuscula := convert("RaMOn")

	fmt.Println(minusculas, mayuscula)

}

// Con parametros
func hello(firtName, lastName string) {
	fmt.Printf("Hello %s %s\n", firtName, lastName)
}

// Cambiar valores por referencia
// Si no usamos un puntero no se va a cambiar el valor ya que la funcion recibe una copia de ese valor y no modifica la varible principal
func change(value *string) {
	*value = "raimonxx"
}

// Retornar valores

func sum(num1, num2 int) int {
	return num1 + num2
}

// Retornar 2 valores

func convert(text string) (string, string) {

	min := strings.ToLower(text)
	may := strings.ToUpper(text)

	return min, may

}
