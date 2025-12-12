package main

// import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone int
}

type Address struct {
	Street string
	City   string
	State  string
	Zip    int
}

type Employee struct {
	personal_details Person
	contact_details  Contact
	address_details  Address
}

func main() {

	// var person_one Employee

	// person_one.personal_details = Person{
	// 	FirstName: "Ritik",
	// 	LastName:  "Singh",
	// 	Age:       22,
	// }

	// person_one.contact_details = Contact{
	// 	Email: "ritik@gmail.com",
	// 	Phone: 1234567890,
	// }

	// person_one.address_details = Address{
	// 	Street: "123 Main St",
	// 	City:   "New York",
	// 	State:  "NY",
	// 	Zip:    10001,
	// }

	// fmt.Println("The person one is ", person_one)

	pointers()
}
