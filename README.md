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
    - Debounce
    - Retry
    - Throttle
    - Timeout
- Concurrency Pattern
    - Fan-In
    - Fan-Out
    - Future
    - Sharding

Learn From Books:

- [Cloud Native Go](https://learning.oreilly.com/library/view/cloud-native-go/9781492076322/) by (Matthew A. Titmus)(https://www.linkedin.com/in/matthew-titmus/)

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

Untuk implementasi bisa dilihat [Breaker](https://github.com/zeihanaulia/go-cloud-native-patterns/tree/main/breaker)

Beberapa repository dan implementasi circuit breaker:

- https://github.com/sony/gobreaker
- https://github.com/streadway/handy
- https://github.com/afex/hystrix-go
- https://github.com/go-kit/kit/tree/master/circuitbreaker


#### Referece

* https://microservices.io/patterns/reliability/circuit-breaker.html