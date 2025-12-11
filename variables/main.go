package main
import "fmt"


func main(){
	// var name string = "Ritik"
	// var age int = 22
	// var isStudent bool = true
    // var salary float64 = 100000.0

	// fmt.Println("Name: ", name)
	// fmt.Println("Age: ", age)
	// fmt.Println("Is Student: ", isStudent)
	// fmt.Println("Salary: ", salary)

	// age:=25
	// name:="Ritik"
	// isStudent:=false

	// fmt.Println("Age is ", age)
	// fmt.Println("Name is ", name)
	// fmt.Println("Is Student is ", isStudent)

	// fmt.Printf("Age is %d, Name is %s, Is Student is %t\n", age, name, isStudent)

	// fmt.Printf("Age is %v, Name is %v, Is Student is %v\n", age, name, isStudent)

	// fmt.Printf("Age is %#v, Name is %#v, Is Student is %#v\n", age, name, isStudent)

	// fmt.Printf("Age is %T, Name is %T, Is Student is %T\n", age, name, isStudent)

	// fmt.Printf("Age is %f, Name is %f, Is Student is %f\n", age, name, isStudent)

	// fmt.Printf("Age is %q, Name is %q, Is Student is %q\n", age, name, isStudent)


	// ******TAKING INPUT FROM USER*****

	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello, ", name)
}