package main

import "fmt"

func main(){
	// fmt.Println("we are learning arrays in go")

	// var name[5]string
	// name[0] = "Ritik"
	// name[1] = "Raj"
	// name[2] = "Rajesh"
	// name[3] = "Rajeshwari"
	// name[4] = "Rajeshwari"

	// fmt.Println("The name of the array is ", name)

	// var numbers = [5]int{1,2,3,4,5}
	// fmt.Println("The numbers of the array is ", len(numbers))


	// *******IF ELSE STATEMENT*****

	// var age int = 22

	// if age >= 18 {
	// 	fmt.Println("You are an adult")
	// }else {
	// 	fmt.Println("You are not an adult")
	// }


	// *******SWITCH STATEMENT*****

	day := "Mondaydwd"

	switch day {
		case "Monday":
			fmt.Println("Today is Monday")
		case "Tuesday":
			fmt.Println("Today is Tuesday")
		case "Wednesday":
			fmt.Println("Today is Wednesday")
		case "Thursday":
			fmt.Println("Today is Thursday")
		case "Friday":
			fmt.Println("Today is Friday")
		case "Saturday":
			fmt.Println("Today is Saturday")
		case "Sunday":
			fmt.Println("Today is Sunday")
		default:
			fmt.Println("Invalid day")
	}


}
