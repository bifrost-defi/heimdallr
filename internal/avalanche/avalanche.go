package avalanche

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Avalanche struct {
	client *ethclient.Client
	sub    ethereum.Subscription
}

var ErrSubscriptionExists = errors.New("subscription already exists")

func New(client *ethclient.Client) *Avalanche {
	return &Avalanche{
		client: client,
	}
}

// Subscribe creates subscription for the contract and returns
func (a *Avalanche) Subscribe(ctx context.Context, contract string) (*Subscription, error) {
	if a.sub != nil {
		return nil, ErrSubscriptionExists
	}

	address := common.HexToAddress(contract)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}

	logs := make(chan types.Log)

	sub, err := a.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return nil, err
	}
	a.sub = sub

	abiReader, err := loadABI("internal/avalanche/abi.json")
	if err != nil {
		return nil, fmt.Errorf("load abi: %w", err)
	}

	contractAbi, err := abi.JSON(abiReader)
	if err != nil {
		return nil, fmt.Errorf("abi json: %w", err)
	}

	s := newSubscription(contractAbi)
	go s.loop(logs)

	return s, nil
}
