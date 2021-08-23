package retry

import (
	"context"
	"log"
	"time"
)

type Requestor func(context.Context) (string, error)

func Retry(requestor Requestor, retries int, delay time.Duration) Requestor {
	return func(ctx context.Context) (string, error) {
		for r := 0; ; r++ {
			response, err := requestor(ctx)
			if err == nil || r >= retries {
				return response, err
			}

			log.Printf("Attempt %d failed; retrying in %v", r+1, delay)

			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return "", ctx.Err()
			}
		}
	}
}
