package retry

import "context"

type Requestor func(context.Context) (string, error)

func Retry(requestor Requestor) Requestor {
	return func(c context.Context) (string, error) {
		response, err := requestor(c)
		return response, err
	}
}
