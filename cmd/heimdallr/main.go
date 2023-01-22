package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"blockwatch.cc/tzgo/rpc"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"heimdallr/config"
	"heimdallr/internal/bridge"
	"heimdallr/internal/evm"
	"heimdallr/internal/server"
	"heimdallr/internal/tezos"
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

	go runServer(ctx, sugar)

	b := bridge.New(eth, tzs, sugar)
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
