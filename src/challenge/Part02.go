package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		pin := i%3 == 0
		pan := i%5 == 0

		switch true {
		case pin && pan:
			fmt.Println("PINPAN")
		case pin:
			fmt.Println("PIN")
		case pan:
			fmt.Println("PAN")
		default:
			fmt.Println(i)
		}

	}
}
