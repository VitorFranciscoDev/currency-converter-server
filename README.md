# Currency Converter API (Go Backend)

A simple backend service written in Go that fetches data from a public API to convert currencies (e.g., BRL to USD).

# Features

- Fetches real-time exchange rates from a public API

- Converts between multiple currencies

- Lightweight and fast Go backend

# Installation

```
git clone https://github.com/VitorFranciscoDev/CurrencyConverterAPI

cd CurrencyConverterAPI

docker build -t currency-converter-api .

docker run -p 8080:8080 currency-converter-api
```

# Example response:

```
{
  "from": "BRL",
  "to": "USD",
  "amount": 100,
  "converted": 18.45
}
```
