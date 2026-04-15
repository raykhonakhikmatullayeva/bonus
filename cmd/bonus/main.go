package main

import (
	"bonus/internal/commission"
	"fmt"
)

func main() {
	name := "SHOKIROV SHUHRAT"
	var amount float64
	var alifChoice int
	fmt.Println("Enter the amount: ")
	fmt.Scan(&amount)
	fmt.Println("Alif card? (1-yes/0-no): ")
	fmt.Scan(&alifChoice)
	isAlifCard := (alifChoice == 0)
	amountInt := int(amount)
	if !commission.IsValidAmount(amountInt) {
		fmt.Println("Amount must be between 500 and 15 000 000!")
		return
	}
	var commission1 int
	if isAlifCard {
		commission1 = 0
	} else {
		commission1 = commission.Commission(amountInt)
	}
	fmt.Printf("=====================Reciept========================\n")
	fmt.Printf("Service: %s\n", name)
	fmt.Printf("Amount(min500 - max15 000 000): %d\n", amountInt)
	fmt.Printf("Commission: %d\n", commission1)
	fmt.Printf("Total: %d\n", amountInt+commission1)
	fmt.Printf("Status: Done\n")
	fmt.Printf("=============================================\n")

}
