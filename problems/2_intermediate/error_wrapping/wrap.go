package errorwrapping

import "errors"

var ErrRoot = errors.New("root")

func Wrap() error {
	// TODO: implement wrapping chain
	return ErrRoot
}
