package tezos

type Subscription struct {
	onWAVAXBurned chan BurnEvent
	onWUSDCBurned chan BurnEvent

	errs chan error
}

type BurnEvent struct {
}

func (s *Subscription) OnWAVAXBurned() <-chan BurnEvent {
	return s.onWAVAXBurned
}

func (s *Subscription) OnWUSDCBurned() <-chan BurnEvent {
	return s.onWUSDCBurned
}

func (s *Subscription) Err() <-chan error {
	return s.errs
}
