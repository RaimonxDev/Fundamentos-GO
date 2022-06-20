package main

import "fmt"

func main() {
	i := 0
	// solo se declara la palabra for y se ejecuta infinitamente
	for {
		fmt.Println(i)
		i++
	}
}
