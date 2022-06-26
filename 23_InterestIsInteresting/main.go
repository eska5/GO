package main

// Task 1
func InterestRate(balance float64) float32 {
	if balance < 0 {
		return 3.213
	} else if balance >= 0 && balance < 1000 {
		return 0.5
	} else if balance >= 1000 && balance < 5000 {
		return 1.621
	} else {
		return 2.475
	}
}

// Task 2
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100
}

// Task 3
func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
}

// Task 4
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var counter int = 0
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		counter++
	}
	return counter
}
