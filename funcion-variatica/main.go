package main

import "fmt"

func main() {
	result := sum(1, 2, 90)
	fmt.Println(result)

}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
