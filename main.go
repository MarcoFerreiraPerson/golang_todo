package main

import ("fmt"
		"strings"
		)


func add_todo(task []string){

}


func main() {
	// create a todo cli application

	var option string

	for option != "4" && option != "exit" {
		fmt.Print("\n1. List out TODOs\n2. Add a TODO\n3. Remove a TODO\n4. Exit\n\n")
		fmt.Print("Please select an option: ")
		fmt.Scan(&option)

		switch option {
			case "1":
				
			case "2":
			case "3":
			default:
				fmt.Print("Invalid input, please try again")

				 
		}

		option = strings.ToLower(option)
		fmt.Print("\033[H\033[2J") //clears the menu
	}
	

}