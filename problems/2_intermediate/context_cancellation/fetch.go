package contextcancellation

import (
	"context"
	"net/http"
	"time"
)

type Result struct { URL string; Err error }

// FetchAll concurrently fetches URLs respecting context cancellation.
func FetchAll(ctx context.Context, urls []string) ([]Result, error) {
	res := make([]Result, len(urls))
	ch := make(chan Result)
	client := http.Client{ Timeout: 2 * time.Second }
	for _, u := range urls {
		go func(u string){
			req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
			_, err := client.Do(req)
			ch <- Result{URL: u, Err: err}
		}(u)
	}
	count := 0
	for count < len(urls) {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case r := <-ch:
			res[count] = r
			count++
		}
	}
	return res, nil
}
