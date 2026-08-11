package main

import (
	"fmt"
	"strconv"
)

func main() {
	println("Hello World")

	print0()
	print1()
}

func print0() {
	n := 65
	s := string(n)
	fmt.Println(s) // 输出：A
}

func print1() {
	n := 65
	s := strconv.Itoa(n)
	fmt.Println(s) // 输出：65
}
