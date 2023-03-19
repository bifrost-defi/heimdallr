package chain

import (
	"math/big"
)

type Event struct {
	user        string
	amount      *big.Int
	coinId      int
	destination string
}

func NewEvent(
	user string,
	amount *big.Int,
	coinId int,
	destination string,
) Event {
	return Event{
		user:        user,
		amount:      amount,
		coinId:      coinId,
		destination: destination,
	}
}

func (e Event) User() string {
	return e.user
}

func (e Event) Amount() *big.Int {
	return e.amount
}

func (e Event) CoinID() int {
	return e.coinId
}

func (e Event) Destination() string {
	return e.destination
}
