# Breaker

## Implementation

1. Definisikan Circuit Type

    ```go
        type Circuit func(context.Context) (string, error)
    ```

2. Buat function `Breaker` yang menerima circuit dan batas kegagalan. Lalu mengembalikan type `Circuit`.
    ```go
        func Breaker(circuit Circuit, failureThreshold uint) Circuit {
            return func(c context.Context) (string, error) {
                response, err := circuit(ctx)
                return response, err
            }
        }
    ```

3. Definisikan variable untuk mencatat kegagalan secara berurutan dan waktu request terakhir
    ```go
        func Breaker(circuit Circuit, failureThreshold uint) Circuit {
            var consecutiveFailures int = 0
            var lastAttempt = time.Now()
            return func(c context.Context) (string, error) {
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
    ```

    - Lalu assign `lastAttempt` setelah function `response, err := circuit(ctx)` dijalankan
    - Setelah itu check tambahkan nilai `consecutiveFailures` jika error dan buat jadi `0` jika tidak error
        ```go
            if err != nil {
                consecutiveFailures++
                return response, err
            }
            consecutiveFailures = 0
        ```

4. Buat logic untuk membuka breaker jika error lebih dari batas
    ```go
        func Breaker(circuit Circuit, failureThreshold uint) Circuit {
            var consecutiveFailures int = 0
            var lastAttempt = time.Now()
            return func(c context.Context) (string, error) {
                d := consecutiveFailures - int(failureThreshold)
                if d >= 0 {
                    shouldRetryAt := lastAttempt.Add(time.Second * 2 << d)
                    if !time.Now().After(shouldRetryAt) {
                        m.RUnlock()
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
    ```

    - Definisikan variable `d` =  `consecutiveFailures - int(failureThreshold)`
    - Jika variable `d` lebih besar atau sama dengan 0
    - Maka check apakah waktu sekarang setelah  nilai `shouldRetryAt`.
    - Jika tidak maka `open circuit` atau return error