package main

import ("fmt"
		"strconv"
		"strings"
		"os"
		"bufio"
		)


func add_todo(task []string){
	
}

func parse_tasks(tasks_string string) []string{

	var tasks = []string{}

	var buff string = "" 
	for i:=0; i < len(tasks_string); i++{
		if tasks_string[i] == ','{
			tasks = append(tasks, buff)
			buff = ""
			continue
		}
		buff += string(tasks_string[i])
	}
	buff = strings.TrimSpace(buff)
	tasks = append(tasks, buff)

	return tasks
}


func main() {
	// create a todo cli application

	var unsaved_tasks = []string{}
	var option string
	reader := bufio.NewReader(os.Stdin)

	for option != "4" && option != "exit" {
		fmt.Print("\n1. List out TODOs\n2. Add a TODO\n3. Remove a TODO\n4. Exit\n\n")
		fmt.Print("Please select an option: ")
		fmt.Scan(&option)

		switch option {
			case "1":
				fmt.Print("Unsaved Tasks: \n")
				for idx, val:= range unsaved_tasks {
					fmt.Print(strconv.Itoa(idx+1)+". " +val+ ",\n")
				}
				var buf string
				fmt.Print("Press any key to continue\n")
				fmt.Scan(&buf)


			case "2":
				var tasks string
				fmt.Print("Please list tasks that you would like to add in comma separated format: ")
				tasks, _ = reader.ReadString('\n')
				tasks = strings.TrimSpace(tasks)
				unsaved_tasks = append(unsaved_tasks, parse_tasks(tasks)...)
				
				// add_todo(tasks)
			case "3":

				var tasks string
				fmt.Print("Please list tasks that you would like to remove in comma separated format: ")
				tasks, _ = reader.ReadString('\n')

				var removed_tasks = parse_tasks(tasks)

				for idx, val:= range unsaved_tasks {
					for _, rval:= range removed_tasks {
						if val == rval {
							unsaved_tasks = append(unsaved_tasks[:idx], unsaved_tasks[idx+1:]...)
							fmt.Print("Removed: "+rval+"\n")
						}
					}
				}
				
			default:
				fmt.Print("Invalid input, please try again")

				 
		}

		option = strings.ToLower(option)
		fmt.Print("\033[H\033[2J") //clears the menu
	}
	

}