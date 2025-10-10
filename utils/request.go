package utils

import (
	"currency-converter/entity"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// function to fetch the exchange rates from the open.er api
func GetRates(base string, symbols []string) (*entity.RatesResponse, error) {

	// api url
	url := fmt.Sprintf("https://open.er-api.com/v6/latest/%s", base)

	// http get request
	resp, err := http.Get(url)

	// if the requests return an error
	if err != nil {
		return nil, err
	}

	// close the connection when the func ends
	defer resp.Body.Close()

	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// variable of the entity
	var rates entity.RatesResponse

	// parse the json into the struct
	if err := json.Unmarshal(body, &rates); err != nil {
		return nil, err
	}

	return &rates, nil
}
