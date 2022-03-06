package tezos

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type Storage struct {
	load       storageLoader
	value      map[string]interface{}
	lastUpdate time.Time
}

type Burning struct {
	User        tezos.Address `json:"user"`
	Amount      tezos.Z       `json:"amount"`
	Destination string        `json:"destination"`
	Ts          time.Time     `json:"ts"`
}

type storageLoader func(ctx context.Context) (map[string]interface{}, error)

func newStorage(load storageLoader) *Storage {
	return &Storage{
		load:       load,
		value:      make(map[string]interface{}),
		lastUpdate: time.Time{},
	}
}

// UpdateBurnings compares last storage state with new one and returns new entries of burnings.
// New state will be saved for the next usage.
func (s *Storage) UpdateBurnings(ctx context.Context) ([]Burning, error) {
	current, err := s.load(ctx)
	if err != nil {
		return nil, fmt.Errorf("load storage: %w", err)
	}

	burningsData, ok := current["burnings"]
	if !ok {
		return nil, fmt.Errorf("invalid map")
	}

	data, err := json.Marshal(burningsData)
	if err != nil {
		return nil, fmt.Errorf("marshal burnings: %w", err)
	}

	var burnings map[string]Burning
	if err := json.Unmarshal(data, &burnings); err != nil {
		return nil, fmt.Errorf("unmarshal burnings: %w", err)
	}

	newBurnings := make([]Burning, 0)
	for _, v := range burnings {
		if v.Ts.After(s.lastUpdate) {
			newBurnings = append(newBurnings, v)
		}
	}

	s.value = current
	s.lastUpdate = time.Now()

	return newBurnings, nil
}
