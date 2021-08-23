# Cloud Native

Tulisan ini akan mengulas hasil pembelajaran dari beberapa sumber [#learn-from-books](#learn-from-books). 
Dan terinspirasi dari obrolan [The Pursuit of Production-Ready Software: Best Practices and Principles](https://www.youtube.com/watch?v=DR0yj7yEDdw)

Semua pembahasan tersebut mengerucut ke system [Dependability](https://en.wikipedia.org/wiki/Dependability).

* Dependability:
    - Attributes
        - Availability
        - Reliability
        - Maintainability
        - Safety
        - Confidentiality
        - Integrity
    - Threats
        - Faults
        - Errors
        - Failures
    - Means
        - Prevention
        - Tolerance
        - Removal
        - Forecasting


- Availability 
- Reliability
- Maintainability

## Learn From Books:

- https://microservices.io/
- [Cloud Native Go](https://learning.oreilly.com/library/view/cloud-native-go/9781492076322/) by [Matthew A. Titmus](https://www.linkedin.com/in/matthew-titmus/)
- [Designing Data-Intensive Applications](https://learning.oreilly.com/library/view/designing-data-intensive-applications/9781491903063/) by [Martin Kleppmann](https://martin.kleppmann.com/)

## Common Failure

Asumsi-asumsi yang salah dalam distributed computing, menurut  [L Peter Deutsch](https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing), 

- The network is reliable: switches fail, routers get misconfigured
- Latency is zero: it takes time to move data across a network
- Bandwidth is infinite: a network can only handle so much data at a time
- The network is secure: don’t share secrets in plain text; encrypt everything
- Topology doesn’t change: servers and services come and go
- There is one administrator: multiple admins lead to heterogeneous solutions
- Transport cost is zero: moving data around costs time and money
- The network is homogeneous: every network is (sometimes very) different
- Services are reliable: services that you depend on can fail at any time

Maka dari itu kita perlu belajar tentang stability pattern yang bisa digunakan.

- Stability Pattern
    - Circuit Breaker
    - Retry
    - Debounce & Throttle
    - Timeout

Sebenarnya beberapa hal diatas bisa tercover dengan tools-tools seperti [istio](https://istio.io/latest/docs/tasks/traffic-management/), [nginx](https://www.nginx.com/resources/library/api-traffic-management-101-monitoring-beyond/), dll. Tidak ada salahnya coba membuat dari sisi script jika tidak menggunakan tools tersebut.


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

### Debounce && Throttle

Kedua pattern ini kurang lebih sama, hanya saja cara kerjanya berbeda.

Throttle menjaga input masuk dalam durasi tertentu, misalnya 100 request per menit.
Debounce hanya menerima input yang sama dalam 1 kali,hanya diawal atau diakhir.

#### Throttle

Throttle membatasi jumlah function direquest. Contoh penggunaan:

- User hanya boleh request dalam 10 kali per detik.
- Membatasi jeda klik per 500 millisecond
- Hanya boleh 3 kali gagal login dalam 24 jam

Tapi biasanya kita menggunakan Throttle memperhitungkan lonjakan aktifitas yang dapat memenuhi system.
Jika system tidak sanggup, maka akan terjadi kegagalan.

#### Debounce

Debounce membatasi jumlah function direquest. Contoh penggunaan:

- User hanya boleh request dalam 10 kali per detik.
- Membatasi jeda klik per 500 millisecond
- Hanya boleh 3 kali gagal login dalam 24 jam

Tapi biasanya kita menggunakan Throttle memperhitungkan lonjakan aktifitas yang dapat memenuhi system.
Jika system tidak sanggup, maka akan terjadi kegagalan.



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

