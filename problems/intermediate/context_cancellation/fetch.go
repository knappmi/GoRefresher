package contextcancellation

import "context"

type Result struct { URL string; Err error }

// FetchAll concurrently fetches URLs respecting context cancellation.
func FetchAll(ctx context.Context, urls []string) ([]Result, error) {
	// TODO: implement concurrent fetch with select on context.Done()
	return nil, ctx.Err()
}
