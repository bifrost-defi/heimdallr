package avalanche

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"bridge-oracle/internal/avalanche/locker"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Avalanche struct {
	// LockManager contract address
	contract common.Address
	// LockManager contract instance
	locker *locker.Locker

	privateKey string

	client *ethclient.Client
	sub    ethereum.Subscription
}

var ErrSubscriptionExists = errors.New("subscription already exists")

func New(client *ethclient.Client, contractAddr string, privateKey string) *Avalanche {
	return &Avalanche{
		contract:   common.HexToAddress(contractAddr),
		privateKey: privateKey,
		client:     client,
	}
}

// Init setups contract.
func (a *Avalanche) Init() error {
	instance, err := locker.NewLocker(a.contract, a.client)
	if err != nil {
		return fmt.Errorf("new locker: %w", err)
	}
	a.locker = instance

	return nil
}

// Subscribe creates subscription for the contract and returns Subscription instance.
func (a *Avalanche) Subscribe(ctx context.Context) (*Subscription, error) {
	if a.sub != nil {
		return nil, ErrSubscriptionExists
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{a.contract},
	}

	logs := make(chan types.Log)

	sub, err := a.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return nil, err
	}
	a.sub = sub

	abiReader, err := loadABI("internal/avalanche/abi/LockManager.abi")
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

func (a *Avalanche) UnlockAVAX(ctx context.Context, user string, amount *big.Int) (string, *big.Int, error) {
	opts, err := a.createTransactor(ctx)
	if err != nil {
		return "", nil, err
	}

	userAddress := common.HexToAddress(user)
	tx, err := a.locker.UnlockAVAX(opts, userAddress, amount)
	if err != nil {
		return "", nil, fmt.Errorf("call unlockAVAX: %w", err)
	}

	return tx.Hash().Hex(), tx.Cost(), nil
}

func (a *Avalanche) UnlockUSDC(ctx context.Context, user string, amount *big.Int) (string, *big.Int, error) {
	opts, err := a.createTransactor(ctx)
	if err != nil {
		return "", nil, err
	}

	userAddress := common.HexToAddress(user)
	tx, err := a.locker.UnlockUSDC(opts, userAddress, amount)
	if err != nil {
		return "", nil, fmt.Errorf("call unlockAVAX: %w", err)
	}

	return tx.Hash().Hex(), tx.Cost(), nil
}

func (a *Avalanche) createTransactor(ctx context.Context) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(a.privateKey)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := a.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, fmt.Errorf("pending nonce: %w", err)
	}

	gasPrice, err := a.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("suggest gas price: %w", err)
	}

	chainID, err := a.client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get chain id: %w", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("new keyed transactor: %w", err)
	}
	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0)
	opts.GasLimit = uint64(300000)
	opts.GasPrice = gasPrice

	return opts, nil
}
