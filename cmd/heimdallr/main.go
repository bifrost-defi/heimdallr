package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"blockwatch.cc/tzgo/rpc"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/xssnick/tonutils-go/liteclient"
	tonUtils "github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"go.uber.org/zap"
	"heimdallr/config"
	"heimdallr/internal/bridge"
	"heimdallr/internal/chain/evm"
	"heimdallr/internal/chain/tezos"
	"heimdallr/internal/chain/ton"
	"heimdallr/internal/server"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	if err := godotenv.Load(); err != nil {
		sugar.Fatalf("No .env file found: %s", err.Error())
	}
	c, err := config.LoadConfig()
	if err != nil {
		err = fmt.Errorf("load config: %w", err)
		sugar.Fatal(err)
	}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-exit
		cancel()

		sugar.Debug("shutdown")
	}()

	eth, err := setupEthereum(ctx, c.Ethereum)
	if err != nil {
		err = fmt.Errorf("setup ethereum: %w", err)
		sugar.Fatal(err)
	}

	tzs, err := setupTezos(ctx, c.Tezos)
	if err != nil {
		err = fmt.Errorf("setup tezos: %w", err)
		sugar.Fatal(err)
	}

	ton, err := setupTON(ctx, c.TON)
	if err != nil {
		err = fmt.Errorf("setup ton: %w", err)
		sugar.Fatal(err)
	}

	go runServer(ctx, sugar)

	b := bridge.New(eth, tzs, ton, sugar)
	if err := b.Run(ctx); err != nil {
		err = fmt.Errorf("run bridge: %w", err)
		sugar.Fatal(err)
	}
}

func runServer(ctx context.Context, logger *zap.SugaredLogger) {
	router := chi.NewRouter()
	router.HandleFunc("/live", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	srv := server.New(router, 8080)
	if err := srv.Run(ctx); err != nil {
		err = fmt.Errorf("run server: %w", err)
		logger.Fatal(err)
	}
}

func setupEthereum(ctx context.Context, config config.Ethereum) (*evm.EVM, error) {
	ethRPCClient, err := ethclient.DialContext(ctx, config.RPC)
	if err != nil {
		return nil, fmt.Errorf("evm rpc dial: %w", err)
	}
	ethWSClient, err := ethclient.DialContext(ctx, config.WS)
	if err != nil {
		return nil, fmt.Errorf("evm rpc dial: %w", err)
	}
	eth := evm.New(ethRPCClient, ethWSClient, config.BridgeContract, config.PrivateKey)

	return eth, nil
}

func setupTezos(ctx context.Context, config config.Tezos) (*tezos.Tezos, error) {
	tzsClient, err := rpc.NewClient(config.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("tezos dial: %w", err)
	}
	tzs := tezos.New(tzsClient, config.PrivateKey)

	if err := tzs.LoadContracts(
		ctx,
		config.BridgeContract,
	); err != nil {
		return nil, fmt.Errorf("set tezos contract: %w", err)
	}

	return tzs, nil
}

func setupTON(ctx context.Context, config config.TON) (*ton.TON, error) {
	pool := liteclient.NewConnectionPool()
	err := pool.AddConnectionsFromConfigUrl(ctx, config.ConfigURL)
	if err != nil {
		return nil, fmt.Errorf("add ton connection: %w", err)
	}
	tonClient := tonUtils.NewAPIClient(pool)

	seed := strings.Split(config.WalletSeed, " ")
	tonWallet, err := wallet.FromSeed(tonClient, seed, wallet.V4R2)
	if err != nil {
		return nil, fmt.Errorf("wallet from seed: %w", err)
	}

	return ton.New(tonClient, tonWallet, config.BridgeContract), nil
}
