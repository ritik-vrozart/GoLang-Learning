package main

import (
	"fmt"
	"net/url"
)


func main(){
	fmt.Println("we are learning url in go")

	url_string := "https://jsonplaceholder.typicode.com/todos/1"


	fmt.Println("The url is ", url_string)

	parsedUrl,err := url.Parse(url_string)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("The parsed url is ", parsedUrl)
}