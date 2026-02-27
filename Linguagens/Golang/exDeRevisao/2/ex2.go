package main

import "fmt"

func main() {
	perm := true
	if perm {
		fmt.Println("Lista autorizada")
		for i := 0; i <= 10; i++ {
			fmt.Println(i)
		}
	} else {
		fmt.Println("Lista não autorizada")
	}
}
