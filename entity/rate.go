package entity

// struct of the response
type RatesResponse struct {
	Base  string             `json:"base"`  // base currency
	Rates map[string]float64 `json:"rates"` // exchange rates
}
