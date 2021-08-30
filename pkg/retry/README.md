# Retry

## Implementation

1. Definisikan type requestor dan function Retry
    ```go
    type Requestor func(context.Context) (string, error)

    func Retry(requestor Requestor) Requestor {
        return func(c context.Context) (string, error) {
            response, err := requestor(c)
            return response, err
        }
    }
    ```

2. Buat logic retry
    ```go
    type Requestor func(context.Context) (string, error)

    func Retry(requestor Requestor, retries int, delay time.Duration) Requestor {
        return func(c context.Context) (string, error) {
            for r := 0; ; r++ {
                response, err := requestor(c)
                 if err == nil || r >= retries {
                    return response, err
                }

                log.Printf("Attempt %d failed; retrying in %v", r + 1, delay)

                select {
                case <-time.After(delay):
                case <-ctx.Done():
                    return "", ctx.Err()
                }
            }
        }
    }
    ```
    - Definidikan argument `retries int` dan `delay time.Duration` pada function `Retry`
    - Buat looping tanpa kondisi karena nanti akan di`break` dengan return
    - Jika requestor tidak error atau `r` lebih besar sama dengan `retries`, maka return response dan err
    - Jika error biarkan melooping.
        ```go
            log.Printf("Attempt %d failed; retrying in %v", r + 1, delay)
            select {
            case <-time.After(delay):
            case <-ctx.Done():
                return "", ctx.Err()
            }
        ```

## Example Usage

```go
var count int

func EmulateTransientError(ctx context.Context) (string, error) {
    count++

    if count <= 3 {
        return "intentional fail", errors.New("error")
    } else {
        return "success", nil
    }
}

func main() {
    r := Retry(EmulateTransientError, 5, 2*time.Second)

    res, err := r(context.Background())

    fmt.Println(res, err)
}
```