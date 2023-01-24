package ton

import "github.com/xssnick/tonutils-go/ton"

type TON struct {
	client     *ton.APIClient
	privateKey string
}

func New(client *ton.APIClient, privateKey string) *TON {
	return &TON{}
}
