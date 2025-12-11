package main

import "fmt"

func divide(a int, b int) (int,error) {
	if(b == 0){
		return 0,fmt.Errorf("division by zero")
	}
	return a / b,nil
}

func main(){
	fmt.Println("we are learning error handling in go")

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("The result of the division is ", result)
}
