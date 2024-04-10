package main

import "fmt"

func main() {
	//** Build a profit calculator **
	var revenue float64
	var expenses float64
	var taxRate float64

	//- Ask for revenue, expenses & tax rate
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("taxRate: ")
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
}
