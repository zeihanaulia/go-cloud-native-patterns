# Cloud Native Pattern

Asumsi-asumsi yang salah dalam distributed computing:

- The network is reliable: switches fail, routers get misconfigured
- Latency is zero: it takes time to move data across a network
- Bandwidth is infinite: a network can only handle so much data at a time
- The network is secure: don’t share secrets in plain text; encrypt everything
- Topology doesn’t change: servers and services come and go
- There is one administrator: multiple admins lead to heterogeneous solutions
- Transport cost is zero: moving data around costs time and money
- The network is homogeneous: every network is (sometimes very) different
- Services are reliable: services that you depend on can fail at any time

Dari asumsi-asumsi diatas bisa dikelompokan kedalam 2 pattern, yaitu `Stability Pattern` dan `Concurrency Pattern`.

- Stability Pattern
    - Circuit Breaker
    - Retry
    - Timeout
    - Debounce
    - Throttle

Learn From Books:

- [Cloud Native Go](https://learning.oreilly.com/library/view/cloud-native-go/9781492076322/) by [Matthew A. Titmus](https://www.linkedin.com/in/matthew-titmus/)

## Stability Pattern

### Circuit Breaker

Circuit Breaker (CB), secara otomatis akan memutus jika terjadi kesalahan secara terus menerus. 
Untuk mencegah terjadinya kegagalan yang lebih besar dan lebih banyak.
CB digunakan untuk menangkap error dan jika sudah mencapai batasnya akan `open circuit` atau memtutus request karena terlalu banyak error.
CB terinspirasi dari kelistrikan, setiap instalasi listrik biasanya dipasang circuit breaker, 
Ketika terjadi konsleting, circute akan terbuka dan mematikan aliran listrik kerumah.

Sekarang bayangkan jika itu didalam system, Service yang kita buat akan mengquery ke database, lalu databasenya mati atau gagal meresponse. 
Jika terjadi cukup lama service kita akan dibanjiri dengan logs error yang sebetulnya tidak berguna.
Ada baiknya jika kita stop saja semua request dari awal hingga database kembali berkerja.

CB pada dasarnya hanyalah function yang berupa [Adapter Pattern](https://refactoring.guru/design-patterns/adapter) diantara request dan response ke service lain. Function `Breaker` akan membungkus Function `Circuit` untuk menambahkan error handling. 
Function `Breaker` memiliki status `open` dan `close`. 
Jika status `close` maka function akan berjalan secara normal dengan meneruskan request yang diterima.
Sebaliknya jika status `open` maka function tidak akan meneruskan dan membuat service gagal lebih cepat.

Dan biasanya akan ada logic auto `close` breaker, Untuk mengecek apakah service sudah berjalan dengan normal.

Untuk implementasi bisa dilihat [Breaker](https://github.com/zeihanaulia/go-cloud-native-patterns/tree/main/breaker).

Beberapa repository dan implementasi circuit breaker:

- https://github.com/sony/gobreaker
- https://github.com/streadway/handy
- https://github.com/afex/hystrix-go
- https://github.com/go-kit/kit/tree/master/circuitbreaker

#### Reference

* https://microservices.io/patterns/reliability/circuit-breaker.html

### Retry

Retry adalah mekanisme pengulangan request ketika terjadi kegagalan ketika membuat request.
Biasanya retry juga memiliki batas mengulang dan juga periode pengulangannya.

Sama seperti CB, untuk membuat `Retry` juga menggunakan [Adapter Pattern](https://refactoring.guru/design-patterns/adapter).
Function `Retry` akan membungkus `Requestor`, untuk mengandle error dari requestor.
Lalu function `Retry` bisa mengkontrol berapa kali retry hingga akhirnya gagal dan juga delay setiap requestnya.

Untuk implementasi bisa dilihat [Retry](https://github.com/zeihanaulia/go-cloud-native-patterns/tree/main/retry).

Beberapa repository dan implementasi retry:

- https://github.com/avast/retry-go
- https://github.com/sethvargo/go-retry
- https://github.com/go-kit/kit/blob/master/sd/lb/retry.go

#### Reference

* http://thinkmicroservices.com/blog/2019/retry-pattern.html

### Timeout

Pada dasarnya timeout akan menghentikan proses dalam durasi waktu.
Biasanya ketika ada masalah didalam service seperti query lambat atau konek ke service lain lambat.
Sehingga proses berjalan menjadi lama. 
Agar tidak dalam proses terus menerus dan client mengunggu lama ada bagusnya kita kasih durasi akses.
Bukan berarti proses lama itu pasti gagal, bisa jadi memang prosesnya perlu waktu sehingga harus dihandle secara asyc, 
bisa dibaca di [long process API](https://github.com/zeihanaulia/go-long-process-api) untuk melihat proof of concept dari handle api yang memproses lama.

Untuk membuat timeout pada service kita hanya perlu memainkan context

```go
ctx := context.Background()
ctxt, cancel := context.WithTimeout(ctx, 10 * time.Second)
defer cancel()

result, err := SomeFunction(ctxt)
```

