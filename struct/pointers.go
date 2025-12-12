package main

import "fmt"

func pointers() {
	fmt.Println("we are learning pointers in go")

	var age int = 22
	var ptr *int = &age

	fmt.Println("The value of age is ", age)
	fmt.Println("The value of ptr is ", ptr)
}
