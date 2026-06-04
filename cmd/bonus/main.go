package main

import (
	"bonus/internal/commission"
	"bonus/internal/receipt"
	"fmt"
	"strings"
)

func main() {
	var senderName, senderLastName string
	var receiverName, receiverLastName string
	var cardNumber string
	var amount float64
	var alifChoice int

	fmt.Print("Enter sender name: ")
	fmt.Scan(&senderName)
	fmt.Print("Enter sender last name: ")
	fmt.Scan(&senderLastName)

	fmt.Print("Enter receiver name: ")
	fmt.Scan(&receiverName)
	fmt.Print("Enter receiver last name: ")
	fmt.Scan(&receiverLastName)
	fmt.Println("Enter card number: ")
	fmt.Scan(&cardNumber)

	fmt.Println("Enter the amount: ")
	fmt.Scan(&amount)

	fmt.Println("Alif card? (1-yes/0-no): ")
	fmt.Scan(&alifChoice)
	isAlifCard := alifChoice == 1

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
	total := amountInt + commission1

	fmt.Printf("=============================================\n")
	fmt.Printf("Sender name: %s %s\n", strings.ToUpper(senderLastName), strings.ToUpper(senderName))
	fmt.Printf("Receiver name: %s %s\n", strings.ToUpper(receiverLastName), strings.ToUpper(receiverName))
	fmt.Printf("Transaction number: %d\n", receipt.GenerateTransactionID())
	fmt.Printf("Credit account: %s\n", receipt.MaskCard(cardNumber))
	fmt.Printf("Date of transaction: %s\n", receipt.CurrentDateTime())
	fmt.Printf("Amount(min500 - max15 000 000): %s sum\n", receipt.SpaceNumber(amountInt))
	fmt.Printf("Commission: %s sum\n", receipt.SpaceNumber(commission1))
	fmt.Printf("Total: %s sum\n", receipt.SpaceNumber(total))
	fmt.Printf("Status: Done\n")
	fmt.Printf("=============================================\n")

}
