package main

import "fmt"

func main() {
	resultado := dobro(2)
	fmt.Println("Resultado", resultado)
}

func dobro(a int) int {
	return a * 2
}
