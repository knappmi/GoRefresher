package customerror

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type ParseError struct {
	Line int
	Msg  string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("parse error at line %d: %s", e.Line, e.Msg)
}

func Lookup(k string) (string, error) {
	return "", fmt.Errorf("lookup %s: %w", k, ErrNotFound)
}

func Parse(line string, ln int) error {
	if line == "" {
		return &ParseError{Line: ln, Msg: "empty"}
	}
	return nil
}
