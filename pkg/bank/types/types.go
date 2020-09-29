package types

// Money represents the amount of money in (diram, cents, kopeyka end etc.) 
type Money int64

//Currency represents the code of currency
type Currency string

//codes of currencies
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

//PAN represents the card number
type PAN string

//Card rpresents information about card
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}