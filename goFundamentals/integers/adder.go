package main

import "fmt"

// Add принимает 2 числа и возвращает их сумму
func Add(x, y int) int {
	return x + y
}

func main() {
	x, y := 2, 3
	fmt.Println(Add(x, y))

}
