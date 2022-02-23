package tezos

import (
	"context"
	"fmt"

	"blockwatch.cc/tzgo/contract"
	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/rpc"
	"blockwatch.cc/tzgo/tezos"
	"github.com/ethereum/go-ethereum/common/math"
)

type Contract struct {
	client *rpc.Client
}

func NewContract(client *rpc.Client) *Contract {
	return &Contract{
		client: client,
	}
}

func loadContract(ctx context.Context, c *rpc.Client, addr string, resolve bool) (*contract.Contract, error) {
	a, err := tezos.ParseAddress(addr)
	if err != nil {
		return nil, err
	}
	if a.Type != tezos.AddressTypeContract {
		return nil, fmt.Errorf("Invalid contract address")
	}
	con := contract.NewContract(a, c)

	if resolve {
		if err := con.Resolve(ctx); err != nil {
			return nil, err
		}
	}
	return con, nil
}

func (c *Contract) GetBalance(ctx context.Context, addr, owner string) ([]byte, error) {
	con, err := loadContract(ctx, c.client, addr, false)
	if err != nil {
		return nil, err
	}
	own, err := tezos.ParseAddress(owner)
	if err != nil {
		return nil, err
	}
	fa1 := con.AsFA1()
	res, err := fa1.GetBalance(ctx, own)
	if err != nil {
		return nil, err
	}
	return res.MarshalBinary()
}

func (c *Contract) GetAllowance(ctx context.Context, addr, owner, spender string) ([]byte, error) {
	con, err := loadContract(ctx, c.client, addr, false)
	if err != nil {
		return nil, err
	}
	own, err := tezos.ParseAddress(owner)
	if err != nil {
		return nil, err
	}
	spend, err := tezos.ParseAddress(spender)
	if err != nil {
		return nil, err
	}
	fa1 := con.AsFA1()
	res, err := fa1.GetAllowance(ctx, own, spend)
	if err != nil {
		return nil, err
	}
	return res.MarshalBinary()
}

func (c *Contract) GetTotalSupply(ctx context.Context, addr string) ([]byte, error) {
	con, err := loadContract(ctx, c.client, addr, false)
	if err != nil {
		return nil, err
	}
	fa1 := con.AsFA1()
	res, err := fa1.GetTotalSupply(ctx)
	if err != nil {
		return nil, err
	}
	return res.MarshalBinary()
}

func (c *Contract) Transfer(ctx context.Context, addr, fromAddr, toAddr, amount string) ([]byte, error) {
	var transfer tezos.Z
	con, err := loadContract(ctx, c.client, addr, false)
	if err != nil {
		return nil, err
	}
	from, err := tezos.ParseAddress(fromAddr)
	if err != nil {
		return nil, err
	}

	to, err := tezos.ParseAddress(toAddr)
	if err != nil {
		return nil, err
	}

	am, _ := math.ParseBig256(amount)
	amT := tezos.Z(*am)

	prim, err := con.RunView(ctx, "transfer",
		micheline.NewPair(
			micheline.NewBytes(from.Bytes22()),
			micheline.NewPair(
				micheline.NewBytes(to.Bytes22()),
				micheline.NewNat(amT.Big()),
			),
		),
	)

	if err != nil {
		return nil, err
	}

	transfer.Set(prim.Int)

	return transfer.MarshalBinary()
}

func (c *Contract) Approve(ctx context.Context, addr, spenderAddr, amount string) ([]byte, error) {
	var approve tezos.Z
	con, err := loadContract(ctx, c.client, addr, false)
	if err != nil {
		return nil, err
	}
	spender, err := tezos.ParseAddress(spenderAddr)
	if err != nil {
		return nil, err
	}

	am, _ := math.ParseBig256(amount)
	amT := tezos.Z(*am)

	prim, err := con.RunView(ctx, "approve",
		micheline.NewPair(
			micheline.NewBytes(spender.Bytes22()),
			micheline.NewNat(amT.Big()),
		),
	)

	if err != nil {
		return nil, err
	}

	approve.Set(prim.Int)

	return approve.MarshalBinary()
}

func (c *Contract) Burn(ctx context.Context, addr, value string) ([]byte, error) {
	var burn tezos.Z
	con, err := loadContract(ctx, c.client, addr, false)
	if err != nil {
		return nil, err
	}

	am, _ := math.ParseBig256(value)
	amT := tezos.Z(*am)

	prim, err := con.RunView(ctx, "burn", micheline.NewNat(amT.Big()))

	if err != nil {
		return nil, err
	}

	burn.Set(prim.Int)

	return burn.MarshalBinary()
}

func (c *Contract) Mint(ctx context.Context, addr, value string) ([]byte, error) {
	var mint tezos.Z
	con, err := loadContract(ctx, c.client, addr, false)
	if err != nil {
		return nil, err
	}

	am, _ := math.ParseBig256(value)
	amT := tezos.Z(*am)

	prim, err := con.RunView(ctx, "mint", micheline.NewNat(amT.Big()))

	if err != nil {
		return nil, err
	}

	mint.Set(prim.Int)

	return mint.MarshalBinary()
}
