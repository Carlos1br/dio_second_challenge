package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("o número", i, "é divisível por 3")
		}
	}
}
