package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// https://jsonplaceholder.typicode.com/todos/1


func main(){

	fmt.Println("we are learning web request in go")
	res,err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer res.Body.Close()

	data,err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("The response is ", string(data))
	
}
