package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generateRandomNumber generates a random number between 1 and 100
func generateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 1
}

// checkNumber checks the generated random number and prints the appropriate message
func checkNumber(num int) {
	fmt.Printf("Generated Random Number: %d\n", num)

	if num == 50 {
		fmt.Println("It's 50!")
	} else if num > 50 && num%2 == 0 {
		fmt.Println("It's closer to 100, and it's even!")
	} else if num > 50 {
		fmt.Println("It's closer to 100")
	} else {
		fmt.Println("It's closer to 0")
	}
}

func main() {
	randomNumber := generateRandomNumber()
	checkNumber(randomNumber)
}