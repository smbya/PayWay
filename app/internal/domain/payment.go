package domain

type Payment struct {
	UserId      int
	Amount      string
	Currency    string
	Destination string
}
