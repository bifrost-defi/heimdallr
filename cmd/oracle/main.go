package main

import (
	"context"
	"fmt"
	"log"

	"blockwatch.cc/tzgo/rpc"
	"bridge-oracle/config"
	"bridge-oracle/internal/avalanche"
	"bridge-oracle/internal/bridge"
	"bridge-oracle/internal/tezos"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %s", err.Error())
	}
	c, err := config.LoadConfig()
	if err != nil {
		err = fmt.Errorf("load config: %w", err)
		log.Fatalln(err)
	}

	ctx := context.Background()

	avaClient, err := ethclient.Dial(c.Avalanche.URL)
	if err != nil {
		err = fmt.Errorf("avalanche dial: %w", err)
		log.Fatalln(err)
	}
	ava := avalanche.New(avaClient, c.Avalanche.Contract, c.Avalanche.PrivateKey)

	tzsClient, err := rpc.NewClient(c.Tezos.URL, nil)
	if err != nil {
		err = fmt.Errorf("tezos dial: %w", err)
		log.Fatalln(err)
	}
	tzs := tezos.New(tzsClient)

	if err := tzs.SetContracts(
		c.Tezos.WAVAXContract,
		c.Tezos.WUSDCContract,
	); err != nil {
		err = fmt.Errorf("set tezos contract: %w", err)
		log.Fatalln(err)
	}

	b := bridge.New(ava, tzs)

	if err := b.Run(ctx); err != nil {
		err = fmt.Errorf("run bridge: %w", err)
		log.Fatalln(err)
	}
}
