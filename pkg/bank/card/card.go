package card

import (
	"bank/pkg/bank/types"
)
// Total - Подсчитывает сумму денег на всех картах
func Total(cards []types.Card) types.Money {
	sum := types.Money(0)
	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance <= 0 {
			continue
		}
		sum += card.Balance
	}
	return sum
}
//PaymentSources - - Подсчитывает сумму денег на всех картах
//Отрицательные суммы и муммы на заблокированных картах игнорируются
func PaymentSources(cards []types.Card) []types.PaymentSource {
	var payment []types.PaymentSource
	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance <= 0 {
			continue
		}
		payment = append(payment, types.PaymentSource{
			Type:    "card",
			Number:  string(card.PAN),
			Balance: card.Balance,
		})
	}
	return payment
}
