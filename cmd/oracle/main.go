package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"blockwatch.cc/tzgo/rpc"
	"bridge-oracle/config"
	"bridge-oracle/internal/avalanche"
	"bridge-oracle/internal/bridge"
	"bridge-oracle/internal/tezos"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
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

	avaClient, err := ethclient.Dial(c.Avalanche.URL)
	if err != nil {
		err = fmt.Errorf("avalanche dial: %w", err)
		sugar.Fatal(err)
	}
	ava := avalanche.New(avaClient, c.Avalanche.Contract, c.Avalanche.PrivateKey)

	tzsClient, err := rpc.NewClient(c.Tezos.URL, nil)
	if err != nil {
		err = fmt.Errorf("tezos dial: %w", err)
		sugar.Fatal(err)
	}
	tzs := tezos.New(tzsClient)

	if err := tzs.SetContracts(
		c.Tezos.WAVAXContract,
		c.Tezos.WUSDCContract,
	); err != nil {
		err = fmt.Errorf("set tezos contract: %w", err)
		sugar.Fatal(err)
	}

	b := bridge.New(ava, tzs, sugar)

	if err := b.Run(ctx); err != nil {
		err = fmt.Errorf("run bridge: %w", err)
		sugar.Fatal(err)
	}
}
