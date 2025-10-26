package customerror

import (
	"errors"
	"testing"
)

func TestErrors(t *testing.T) {
	_, err := Lookup("k")
	if !errors.Is(err, ErrNotFound) { t.Fatal("Is failed") }
	perr := Parse("", 3)
	var target *ParseError
	if !errors.As(perr, &target) || target.Line!=3 { t.Fatal("As failed or wrong line") }
}
