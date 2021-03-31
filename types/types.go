package types

//Money type int64
type Money int64

//Currency type string
type Currency string
//
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)
//PAN type string
type PAN string

//Card model
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

//Payment int and money
type Payment struct {
	ID     int
	Amount Money
}
//PaymentSource str and money
type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
