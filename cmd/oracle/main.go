package main

import (
	"context"
	"fmt"
	"log"

	"blockwatch.cc/tzgo/rpc"
	"bridge-oracle/internal/avalanche"
	"bridge-oracle/internal/bridge"
	"bridge-oracle/internal/tezos"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

const (
	avalancheURL      = ""
	avalancheContract = ""

	tezosURL      = ""
	tezosContract = ""
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %s", err.Error())
	}

	ctx := context.Background()

	avaClient, err := ethclient.Dial(avalancheURL)
	if err != nil {
		err = fmt.Errorf("avalanche dial: %w", err)
		log.Fatalln(err)
	}
	ava := avalanche.New(avaClient)

	tzsClient, err := rpc.NewClient(tezosURL, nil)
	if err != nil {
		err = fmt.Errorf("tezos dial: %w", err)
		log.Fatalln(err)
	}
	tzs := tezos.New(tzsClient)

	b := bridge.New(ava, tzs)
	b.SetAvalancheContract(avalancheContract)
	b.SetTezosContract(tezosContract)

	if err := b.Run(ctx); err != nil {
		err = fmt.Errorf("run bridge: %w", err)
		log.Fatalln(err)
	}
}
