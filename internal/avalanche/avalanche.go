package avalanche

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"heimdallr/internal/avalanche/locker"
)

type Avalanche struct {
	// LockManager contract address
	contract common.Address
	// LockManager contract instance
	locker *locker.Locker

	privateKey string

	rpc *ethclient.Client
	ws  *ethclient.Client
}

func New(rpc *ethclient.Client, ws *ethclient.Client, contractAddr string, privateKey string) *Avalanche {
	return &Avalanche{
		contract:   common.HexToAddress(contractAddr),
		privateKey: privateKey,
		rpc:        rpc,
		ws:         ws,
	}
}

func (a *Avalanche) init() error {
	instance, err := locker.NewLocker(a.contract, a.ws)
	if err != nil {
		return fmt.Errorf("new locker: %w", err)
	}
	a.locker = instance

	return nil
}

// Subscribe creates subscription for the contract and returns Subscription instance.
func (a *Avalanche) Subscribe(ctx context.Context) (*Subscription, error) {
	if err := a.init(); err != nil {
		return nil, fmt.Errorf("init: %w", err)
	}

	opts := &bind.WatchOpts{Context: ctx}

	avaxEvents := make(chan *locker.LockerAVAXLocked)
	avaxSub, err := a.locker.WatchAVAXLocked(opts, avaxEvents)
	if err != nil {
		return nil, fmt.Errorf("watch avax: %w", err)
	}

	usdcEvents := make(chan *locker.LockerUSDCLocked)
	usdcSub, err := a.locker.WatchUSDCLocked(opts, usdcEvents)
	if err != nil {
		return nil, fmt.Errorf("watch avax: %w", err)
	}

	s := newSubscription()
	go s.loopAVAX(ctx, avaxSub, avaxEvents)
	go s.loopUSDC(ctx, usdcSub, usdcEvents)

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
