package evm

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"heimdallr/internal/evm/locker"
)

type EVM struct {
	// Bridge contract address
	contract common.Address
	// Bridge contract instance
	// TODO
	locker *locker.Locker

	privateKey string

	rpc *ethclient.Client
	ws  *ethclient.Client
}

func New(rpc *ethclient.Client, ws *ethclient.Client, contractAddr string, privateKey string) *EVM {
	return &EVM{
		contract:   common.HexToAddress(contractAddr),
		privateKey: privateKey,
		rpc:        rpc,
		ws:         ws,
	}
}

func (a *EVM) init() error {
	instance, err := locker.NewLocker(a.contract, a.ws)
	if err != nil {
		return fmt.Errorf("new locker: %w", err)
	}
	a.locker = instance

	return nil
}

// Subscribe creates subscription for the contract and returns Subscription instance.
func (a *EVM) Subscribe(ctx context.Context) (*Subscription, error) {
	if err := a.init(); err != nil {
		return nil, fmt.Errorf("init: %w", err)
	}

	opts := &bind.WatchOpts{Context: ctx}

	ethEvents := make(chan *locker.LockerAVAXLocked)
	ethSub, err := a.locker.WatchAVAXLocked(opts, ethEvents)
	if err != nil {
		return nil, fmt.Errorf("watch eth: %w", err)
	}

	s := newSubscription()
	go s.loop(ctx, ethSub, ethEvents)

	return s, nil
}

func (a *EVM) UnlockETH(ctx context.Context, user string, amount *big.Int) (string, *big.Int, error) {
	opts, err := a.createTransactor(ctx)
	if err != nil {
		return "", nil, err
	}

	userAddress := common.HexToAddress(user)
	tx, err := a.locker.UnlockAVAX(opts, userAddress, amount)
	if err != nil {
		return "", nil, fmt.Errorf("call unlock: %w", err)
	}

	return tx.Hash().Hex(), tx.Cost(), nil
}

func (a *EVM) createTransactor(ctx context.Context) (*bind.TransactOpts, error) {
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
	nonce, err := a.rpc.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, fmt.Errorf("pending nonce: %w", err)
	}

	gasPrice, err := a.rpc.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("suggest gas price: %w", err)
	}

	chainID, err := a.rpc.ChainID(ctx)
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
