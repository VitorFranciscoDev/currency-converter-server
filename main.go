package main

import (
	"currency-converter/usecases"
	"currency-converter/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/convert", convertHandler) // endpoint

	http.ListenAndServe(":8080", nil) // start the server
}

func convertHandler(w http.ResponseWriter, r *http.Request) {

	// query parameters
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	amountStr := r.URL.Query().Get("amount")

	// validate parameters
	if from == "" || to == "" || amountStr == "" {
		http.Error(w, "Missing 'from', 'to', or 'amount' parameter", http.StatusBadRequest)
		return
	}

	// convert amount to float
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	// fetch exchange rate
	rates, err := utils.GetRates(from, []string{to})
	if err != nil {
		http.Error(w, "Error fetching rates: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// calculate the converted amount
	converted := usecases.Convert(amount, rates.Rates[to])

	// map to json response
	response := map[string]interface{}{
		"from":      from,
		"to":        to,
		"amount":    amount,
		"converted": converted,
		"rate":      rates.Rates[to],
	}

	w.Header().Set("Content-Type", "application/json") // response type to json
	json.NewEncoder(w).Encode(response)                // encode the map to json and send it
}
