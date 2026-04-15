package commission

const MinAmount = 500
const MaxAmount = 15000000

func IsValidAmount(amount int) bool {
	if amount < MinAmount {
		return false
	}
	if amount > MaxAmount {
		return false
	}
	return true
}
