package errorwrapping

import "fmt"
import "errors"

var ErrRoot = errors.New("root")

func Wrap() error {
	return fmt.Errorf("layer2: %w", fmt.Errorf("layer1: %w", ErrRoot))
}
