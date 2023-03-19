package bridge

import (
	"heimdallr/internal/chain"
)

func rollbackEvent(event chain.Event) chain.Event {
	// Swap user and destination because event for rolling back
	// must be mirrored to use it as an event from another blockchain.
	return chain.NewEvent(
		event.Destination(),
		event.Amount(),
		event.CoinID(),
		event.User(),
	)
}
