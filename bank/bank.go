package main

import "fmt"

func menu() {
	fmt.Println()
	fmt.Println("*** Welcome to the bank *** ")
	fmt.Println("What do you want to do?")

	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Just a loop")
	fmt.Println("5. Exit")

	fmt.Println("Select your choice: ")
}

func justALoopFunction() {
	for i := 0; i < 10; i++ {
		fmt.Println("My index: ", i*2+1)
	}
}

func main() {
	var option int
	var accountBalance = 1000.0

	// Bucle infinito
	for {
		menu()
		fmt.Scan(&option)
		fmt.Println("Your choice: ", option)

		switch option {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Println("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. must be greater that 0")
				return
			}
			accountBalance += depositAmount
			fmt.Println("Your new balance is: ", accountBalance)
		case 3:
			fmt.Println("Your withdraw: ")
			var amount float64
			fmt.Scan(&amount)

			if amount >= accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				return
			}
			accountBalance -= amount
			fmt.Println("Your new balance is: ", accountBalance)
		case 4:
			justALoopFunction()
		default:
			fmt.Println("Goodbye")
			break
		}
	}
}
