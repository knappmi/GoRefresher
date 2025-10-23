package buildtags

import "testing"

func TestMode(t *testing.T) { if Mode()=="" { t.Fatal("empty mode") } }
