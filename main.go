package main

import (
	"fmt"
	"os"
)

// simple student manage system
// 1. show welcome messenge
// 2. add student information
// 3. edit student information
// 4. show student information
// 5. quit

func showMenu() {
	fmt.Println("Welcome simple student manage system.")
	fmt.Println("1. add student information")
	fmt.Println("2. modify student information")
	fmt.Println("3. show all students information")
	fmt.Println("0. quit")
}

func main() {
	var stus studentMgr
	for {
		// Show welcome messenge && Show menu
		showMenu()

		var input int
		fmt.Printf("Please choice option: ")
		fmt.Scanf("%d", &input)
		switch input {
		// add student information
		case 1:
			var (
				id          int
				name, class string
			)
			fmt.Printf("Please type student ID: ")
			fmt.Scanf("%d\n", &id)
			fmt.Printf("Please type student name: ")
			fmt.Scanf("%s\n", &name)
			fmt.Printf("Please type student class: ")
			fmt.Scanf("%s\n", &class)
			stus.addStudent(newStudent(id, name, class))
		// edit student information
		case 2:
			var (
				id          int
				name, class string
			)
			fmt.Printf("Please type student ID: ")
			fmt.Scanf("%d\n", &id)
			fmt.Printf("Please type student name: ")
			fmt.Scanf("%s\n", &name)
			fmt.Printf("Please type student class: ")
			fmt.Scanf("%s\n", &class)
			stus.modifyStudent(newStudent(id, name, class))
		// show student information
		case 3:
			stus.showStudent()
		// quit
		case 0:
			fmt.Println("user cancelled")
			os.Exit(0)
		}

	}

}
