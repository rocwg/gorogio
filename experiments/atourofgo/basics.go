package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func basics00() {
	// welcome
	fmt.Println("Hello, 世界")
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())

	// Packages
	fmt.Println("My favorite number is", rand.Intn(10))

	// Imports
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// Exported names
	fmt.Println(math.Pi)

	// Functions
	fmt.Println(add(42, 13))

	// Multiple results
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// Named return values
	fmt.Println(split(17))

}

func add(x, y int) int {
	//func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
