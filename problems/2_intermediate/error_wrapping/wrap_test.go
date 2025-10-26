package errorwrapping

import (
	"errors"
	"testing"
)

func TestWrap(t *testing.T) { if !errors.Is(Wrap(), ErrRoot) { t.Fatal("expected root") } }
