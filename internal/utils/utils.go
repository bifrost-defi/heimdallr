package utils

import (
	"fmt"
	"math/big"
)

func BigIntToHex(i *big.Int) string {
	return fmt.Sprintf("0x%x", i)
}
