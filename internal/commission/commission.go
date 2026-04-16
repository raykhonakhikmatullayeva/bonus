package commission

const CommissionRate = 29
const Percentage = 100

func Commission(amountInt int) int {
	if IsValidAmount(amountInt) {

		return amountInt * CommissionRate / Percentage / Percentage
	}
	return 0
}
