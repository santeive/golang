package main

import "fmt"

func outputText(text string) {
	fmt.Print(text)
}

func returnTwoValues() (int, int) {
	// Multiple value return
	aVar := 100
	bVar := 200

	return aVar, bVar
}

func alternativeReturnTwoValues() (aVar int, bVar int) {
	// Multiple value return
	aVar = 1
	bVar = 1

	// Good but not explicit
	return
}

func getUserInput(infoText string) float64 {
	//- Ask for revenue, expenses & tax rate
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	//- Calculate earnings before tax (EBT) and earnings after tax (profit)
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func main() {
	//** Build a profit calculator **
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("expenses: ")
	taxRate := getUserInput("taxRate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	//- Output EBT, profit and the ratio
	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("profit: %.1f\n", profit)
	fmt.Printf("ratio: %.1f\n", ratio)
	fmt.Println(alternativeReturnTwoValues())

	// Create a module
	// go mod init example.com/basics-practice
}
