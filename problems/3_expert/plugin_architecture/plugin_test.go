package pluginarch

import "testing"

func TestLoad(t *testing.T) {
	if err := Load("nonexistent.so"); err == nil { t.Fatal("expected error loading plugin") }
}
