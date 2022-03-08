package bridge

import "math/big"

type RollbackEvent struct {
	user        string
	amount      *big.Int
	destination string
}

var _ Event = (*RollbackEvent)(nil)

func (e RollbackEvent) User() string {
	return e.user
}

func (e RollbackEvent) Amount() *big.Int {
	return e.amount
}

func (e RollbackEvent) Destination() string {
	return e.destination
}

func rollbackEvent(event Event) RollbackEvent {
	// Swap user and destination because event for rolling back
	// must be mirrored to use it as an event from another blockchain.
	return RollbackEvent{
		user:        event.Destination(),
		amount:      event.Amount(),
		destination: event.User(),
	}
}
