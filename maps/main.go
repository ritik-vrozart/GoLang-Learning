package main
import "fmt"



func main(){
	fmt.Println("we are learning maps in go")


	studentScores := make(map[string]int)

	studentScores["Ritik"] = 90
	studentScores["Raj"] = 80
	studentScores["Rajesh"] = 70
	studentScores["Rajeshwari"] = 60
	studentScores["Rajeshwari"] = 60

	fmt.Println("The scores of the students are ", studentScores)
}
