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
	"heimdallr/internal/evm"
	"heimdallr/internal/server"
	"heimdallr/internal/tezos"
	"heimdallr/internal/ton"
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

	// Connect to Ethereum

	ethRPCClient, err := ethclient.Dial(c.Ethereum.RPC)
	if err != nil {
		err = fmt.Errorf("evm rpc dial: %w", err)
		sugar.Fatal(err)
	}
	ethWSClient, err := ethclient.Dial(c.Ethereum.WS)
	if err != nil {
		err = fmt.Errorf("evm rpc dial: %w", err)
		sugar.Fatal(err)
	}
	eth := evm.New(ethRPCClient, ethWSClient, c.Ethereum.BridgeContract, c.Ethereum.PrivateKey)

	// Connect to Tezos

	tzsClient, err := rpc.NewClient(c.Tezos.URL, nil)
	if err != nil {
		err = fmt.Errorf("tezos dial: %w", err)
		sugar.Fatal(err)
	}
	tzs := tezos.New(tzsClient, c.Tezos.PrivateKey)

	if err := tzs.LoadContracts(
		ctx,
		c.Tezos.BridgeContract,
	); err != nil {
		err = fmt.Errorf("set tezos contract: %w", err)
		sugar.Fatal(err)
	}

	// Connect to TON

	pool := liteclient.NewConnectionPool()
	err = pool.AddConnection(ctx, c.TON.URL, c.TON.ServerKey)
	if err != nil {
		err = fmt.Errorf("add ton connection: %w", err)
		sugar.Fatal(err)
	}
	tonClient := tonUtils.NewAPIClient(pool)
	seed := strings.Split(c.TON.WalletSeed, " ")
	tonWallet, err := wallet.FromSeed(tonClient, seed, wallet.V3)

	ton := ton.New(tonClient, tonWallet, c.TON.BridgeContract)

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
