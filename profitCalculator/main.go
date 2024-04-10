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

func main() {
	//** Build a profit calculator **
	var revenue float64
	var expenses float64
	var taxRate float64

	//- Ask for revenue, expenses & tax rate
	outputText("Revenue: ")
	fmt.Scan(&revenue)

	outputText("Expenses: ")
	fmt.Scan(&expenses)

	outputText("taxRate: ")
	fmt.Scan(&taxRate)

	//- Calculate earnings before tax (EBT) and earnings after tax (profit)
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate)

	//- Calculate ratio (EBT/profit)
	ratio := ebt / profit

	//- Output EBT, profit and the ratio
	fmt.Println("EBT: ", ebt)
	fmt.Println("profit: ", profit)
	fmt.Println("ratio: ", ratio)

	fmt.Print(alternativeReturnTwoValues())

	// Create a module
	// go mod init example.com/basics-practice
}
