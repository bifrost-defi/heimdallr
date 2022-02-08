package main

import (
	"bridge-oracle/internal/avalanche"
	"bridge-oracle/internal/bridge"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("")
	if err != nil {
		err = fmt.Errorf("ethereum dial: %w", err)
		log.Fatalln(err)
	}

	bridge.New(
		avalanche.New(client),
	)
}
