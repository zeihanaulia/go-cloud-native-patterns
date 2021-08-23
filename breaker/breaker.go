package breaker

import (
	"context"
	"errors"
	"time"
)

type Circuit func(context.Context) (string, error)

func Breaker(circuit Circuit, failureThreshold uint) Circuit {
	var consecutiveFailures int = 0
	var lastAttempt = time.Now()
	return func(ctx context.Context) (string, error) {
		d := consecutiveFailures - int(failureThreshold)
		if d >= 0 {
			shouldRetryAt := lastAttempt.Add(time.Second * 2 << d)
			if !time.Now().After(shouldRetryAt) {
				return "", errors.New("service unreachable")
			}
		}

		response, err := circuit(ctx)
		lastAttempt = time.Now()
		if err != nil {
			consecutiveFailures++
			return response, err
		}
		consecutiveFailures = 0

		return response, nil
	}
}
