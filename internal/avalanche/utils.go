package avalanche

import (
	"fmt"
	"io"
	"os"
)

func loadABI(path string) (io.Reader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	return f, nil
}
