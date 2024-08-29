package main

import (
	"fmt"
	"os"

	"go_basics/calculator"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("ENV"))

	fmt.Println(calculator.Offset)
	fmt.Println("Sum: ", calculator.Sum(1, 2))
	fmt.Println("Multiply: ", calculator.Multiply(1, 2))
}
