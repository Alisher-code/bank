package card

import (
	"bank/pkg/bank/types"
	"fmt"
)
func ExamplePaymentSources() {
	cards := []types.Card{
		{
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 30_000_03,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 9999",
			Balance: -50_000_05,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 2020",
			Balance: 40_000_04,
			Active:  false,
		},
	}
	payment := PaymentSources(cards)
	for _, payments := range payment {
		fmt.Println(payments.Number)
	}

	//Output:
	//5058 xxxx xxxx 8888

}
