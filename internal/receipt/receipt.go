package receipt

import (
	"fmt"
	"math/rand"
	"time"
)

func MaskCard(cardNumber string) string {
	if len(cardNumber) < 4 {
		return "invalid card number"
	}
	last4 := cardNumber[len(cardNumber)-4:]
	return fmt.Sprintf("Uzcard**%s", last4)
}
func GenerateTransactionID() int {
	//rand.Seed(time.Now().UnixNano())
	return rand.Intn(900000000) + 100000000
}
func CurrentDateTime() string {
	return time.Now().Format("02-01-2006 15:04")
}
func SpaceNumber(n int) string {
	s := fmt.Sprintf("%d", n)
	if len(s) <= 3 {
		return s
	}
	result := ""
	for i, ch := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result += " "
		}
		result += string(ch)
	}
	return result
}
