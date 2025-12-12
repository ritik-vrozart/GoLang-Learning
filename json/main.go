package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}


func main(){
	fmt.Println("we are learning json in go")

	user := User {
		Name: "Ritik",
		Age: 22,
		Email: "ritik@gmail.com",
	}

	jsonData, err  := json.Marshal(user)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("The json data is ", string(jsonData))
}
