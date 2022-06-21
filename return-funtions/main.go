package main

import "fmt"

func main() {

	//nums := []int{1, 4, 6, 7, 8, 90, 70}

	/*result := filter(nums, func(num int) bool {
		if num <= 10 {
			return true
		}
		return false
	})

	fmt.Println(result)*/

	//Los segundos parentesis ejecutara la funcion "x" que es el retorno de la funcion hello
	x := hello("ramon")()
	fmt.Println(x)

}

func hello(name string) func() string {
	return func() string {
		return "hello" + name
	}
}

// Funcion que retornan funciones

/*func filter(nums []int, callback func(int) bool) []int {

	var result []int

	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result

}*/
