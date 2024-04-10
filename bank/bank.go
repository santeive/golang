package main

import "fmt"

func menu() {
	fmt.Println("*** Welcome to the bank *** ")
	fmt.Println("What do you want to do?")

	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	fmt.Println("Select your choice: ")
}

func main() {
	var option int

	menu()
	fmt.Scan(&option)

	fmt.Println("Your choice: ", option)
}
