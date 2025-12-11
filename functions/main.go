package main

import "fmt"


func sum(a int , b int) int {
	return a + b
}


func main(){

	fmt.Println("we are learning functions in go")

	a:= 25
	b:= 30

	var result int = sum(a,b)

	fmt.Println("The sum of ", a, " and ", b, " is ", result)
}