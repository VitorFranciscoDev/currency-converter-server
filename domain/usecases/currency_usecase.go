package usecases

import "github.com/VitorFranciscoDev/currency-converter-server/infrastructure/datastore/interfaces"

type CurrencyUseCase struct {
	r interfaces.CurrencyRepository
}

func NewCurrencyUseCase(r interfaces.CurrencyRepository) *CurrencyUseCase {
	return &CurrencyUseCase{r}
}
