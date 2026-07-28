package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// ## Packages
//
// Every Go program is made up of packages.
//
// Programs start running in package `main`.
//
// This program is using the packages with import paths `"fmt"` and `"math/rand"`.
//
// By convention, the package name is the same as the last element of the import path.
// For instance, the `"math/rand"` package comprises files that begin with the statement `package rand`.
//
// ## Imports
//
// This code groups the imports into a parenthesized, "factored" import statement.
//
// You can also write multiple import statements, like:
//
// ```go
// import "fmt"
// import "math"
// ```
//
// But it is good style to use the factored import statement.
func main() {
	fmt.Println("Hello, 世界")

	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println(math.Pi)
}
