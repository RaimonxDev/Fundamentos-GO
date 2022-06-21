package main

import "fmt"

func main() {
	func(text string) {
		fmt.Println("hello " + text)
	}("ramon")
}
