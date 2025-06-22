package main

import "fmt"

func Help() {
	fmt.Println()
	fmt.Println("This application is specially made to turn the application to Single User Mode and vice versa!")
	fmt.Println()
	fmt.Println("You should supply required arguments for its proper functioning. Please follow the below details")
	fmt.Println()
	fmt.Println("Command							Action")
	fmt.Println()
	fmt.Println("-config						-	To config the application,Note: Donot change the config file's name")
	fmt.Println("-release					-	To release the application from Single User Mode")
	fmt.Println("-down						-	To switch the application to Single User Mode")
	fmt.Println()
	fmt.Println()
}
