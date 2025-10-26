package contextcancellation

import (
	"context"
	"testing"
)

func TestFetchAll(t *testing.T) { t.Skip("TODO: implement context cancellation tests") }

func TestFetchAllCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // immediate cancel
	_, err := FetchAll(ctx, []string{"http://example.com"})
	if err == nil { t.Fatal("expected cancellation error") }
}
