package main

import "fmt"

func main() {
	users := 0
	for i := 0; i < 11; i++ {
		users++
		if users%2 == 0 {
			fmt.Println("Usuario com numero par:", users)
		}
	}
}
