package main

import  (
	"os"
	"fmt"
)

func main(){
	file,err := os.Create("example.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	file.WriteString("Hello, World!")
	defer file.Close()

	fmt.Println("File created successfully")
}
