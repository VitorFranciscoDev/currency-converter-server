package usecases

// func to convert the initial amount to a new currency based on the exchange rate
func Convert(amount float64, rate float64) float64 {
	return amount * rate
}
