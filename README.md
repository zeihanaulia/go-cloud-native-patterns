# Cloud Native

Tulisan ini akan mengulas hasil pembelajaran dari beberapa sumber [#learn-from-books](#learn-from-books). 
Dan terinspirasi dari obrolan [The Pursuit of Production-Ready Software: Best Practices and Principles](https://www.youtube.com/watch?v=DR0yj7yEDdw) untuk mengulik lebih jauh.

Semua pembahasan tersebut mengerucut ke system [Dependability](https://en.wikipedia.org/wiki/Dependability). 
Konsep Dependability pertamakali didefinisikan oleh [Jean-Claude Laprie](https://ieeexplore.ieee.org/document/532603) sekitar 35 tahun yang lalu.
Jadi ini bukanlah konsep yang baru.

Mengutip dari gambar di [wikipedia](https://en.wikipedia.org/wiki/Dependability). Komponen dependability seperti ini:

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
    Kemampuan system berkerja dalam suatu waktu secara acak. Biasanya diekspresikan dari berapa banyak request yang diterima oleh system. 
    Uptime dibagi dengan total time.
- Reliability
    Kemampuan system berkerja dalam interval waktu. Biasanya diekspresikan dengan mean time beetwen  failures  (MTBF: total time dibagi denan total failures) atau failure rate (number or failures dibagi total time).
- Maintainability
    Kemampuan suatu sistem untuk mengalami modifikasi dan perbaikan. Ada beberapa jenis untuk mengukur maintainability.
    Mulai dari perhitungan cyclomatic complexity hingga pelacakan total waktu yang diperlukan untuk melakukan perbaikan system
    atau mengembalikan ke status sebelumnya.

Ada 4 kategori teknik yang dapat digunakan untuk mengimprove systems depentability:

- Fault prevention (Scalability / Loose Coupling)
    Digunakan selama pembuatan system untuk mencegah terjadinya kesalahan.
- Fault tolerance (Resilence / Loose Coupling)
    Digunakan selama mendesign system dan implementasi untuk mencegah kegagalan karena adanya kesalahan.
- Fault removal (Manageability)
    Digunakan untuk mengurangi jumlah dan tingkat keparahan(severity) kesalahan.
- Fault forecasting (Observability)
    Digunakan untuk menidentifikasi keberadaan, penciptaan dan kosekuensi dari kesalahan.

### Fault Prevention

Fault Prevention menjadi dasar dari semuanya. Kenapa? karena disini adalah tahap pembuatan system.
Lalu apa saja yang harus dilakukan:

- Good Programming Practice
    - Pair Programming
    - Test-Driven Development
    - Code Review Practice
- Language features
    Pemilihan bahasa cukup berpengaruh. Setiap bahasa punya paradigma dan karakteristik masing masing. Bisa baca dibuku [Clean Architechture - PARADIGM OVERVIEW](https://learning.oreilly.com/library/view/clean-architecture-a/9780134494272/ch3.xhtml#toclev_13)
- Scalability
    Kemampuan dari suatu sistem untuk terus memberikan layanan dalam menghadapai perubahan sesuai dengan permintaan.
- Loose Couplling
    Artinya setiap sistem saling terhubung tetapi atar service hanya tau interfacenya saja.

### Fault Tolrerance

Setelah Fault Prevention, maka masuk ke Fault Tolerance.

Ada beberapa nama lain seperti self-repair, self-healing, resillience. 
Semuanya menggamarkan kemampuan sistem untuk mendeteksi kesalahan dan mencegah menjadi kesalahan yang lebih besar.
Biasanya terdiri dari 2 bagian, *error detection* dan *recovery*

### Fault Removal

Selanjutnya adalah Fault Removal adakag proses mengurangi jumlah dan tingkat keparahan (severity) kesalahan.
Bahkan dikondisi ideal pun ada banyak cara yang membuat sistem melakukan kesalahan atau berperilaku tidak semestinya.
Mungkin gagal melakuka tindakan yang diharapkan atau melakukan tindakan yang salah.

Banyak kesalahan bisa diindentifikasi melalui testing yang memungkinkan untuk kita menverifikasi sistem yang kita buat.

#### Verification and Testing

Cuma ada satu cara buat menemukan kesalahan pada software. yaitu testing.
Kalau bukan kita yang nemuin ya user yang menggunakan sistem kita yang mengerikan jika usernya jahat.
Bisa berbahaya jika dia mengambil keuntungan dari kesalahan tersebut.

Ada 2 pendekatan:

1. Static Analisis
    Berguna untuk memberikan feedback diawal, memaksa untuk melakukan praktek yang konsisten dan menemukan error yang umum.
    Tanpa perlu bantuan manusia.
    Biasanya bisa menggunakan tools seperti , codeclimate, dll
2. Dynamic Analisis
    Kalau ini testing yang memerlukan manusia.

Kunci dari software testing adalah membuat software yang `designed for testability`.
function yang teatabilitynya tinggi adalah function yang memiliki tujuan tunggal dengan input dan output yang terdefinisi dengan baik atau sedikit efek samping.

#### Manageability

Kesalahan pada sistem ada karena sistem tidak berkerja sesuai requirement.
Dengan `Designing for manageability` mengizinkan perilaku system diubah tanpa ada perubahan dicode.
Manageable system pada dasarnya seperti sistem yang memiliki `knob` yang memungkinkan kontrol secara real-time untuk menjaga sistem kita tetap aman, bejalan lancar dan sesuai dengan requirement. Misalnya seperti feature flags yang bisa menyalakan atau mematikan fitur atau loading plug-in yang mengubah perilaku.

### Fault Forecasting

Ini adalah tahap terakhir, tahap ini dibangun berdasarkan pengetahuan dari kejadian yang ada dan kumpulan dari solusi solusi yang diterapkan.
Biasanya cuma menebak nebak saja. Tapi dengan membuat system `design for observability` indikator kesalahan bisa ditrack jadi kita bisa memprediksi dengan tepat sebelum berubah menjadi error.

## Twelve-Factor App

Jadi sekitar tahun 2010an, para developer dari heroku menyadari sesuatu karena seringnya mereka melihat development web memiliki masalah yang sama.
Lalu mereka menyusun [The Twelve-Factor App](https://12factor.net/), ini adalah kumpulan aturan main dan panduan yang merupakan metedologi pengembangan web.
Metedologinya sebagai berikut:

- Gunakan deklaratif format untuk setup automation, guna membantu developer baru yang join project.
- Memiliki `clean contract` pada sistem operasi yang digunakan.
- Cocok digunakan pada cloud platform modern sehingga tidak memerlukan server dan sysadmin.
- Meminimalisir perbedaan antara environment development dan production.
- Dapat `scale up` tanpa perubahan yang signifikan pada tools, architechture atau development practice.

Jadi apa saja isi dari cloud native:

1. Codebase

    > One codebase tracked in revision control, many deploys.

    Satu codebase untuk segala environtment, biasanya development, staging dan production.


2. Dependencies

    > Explicitly declare and isolate (code) dependencies.

    TODO: penjelasan

3. Configuration

    > Store configuration in the environment.

    Konfigurasi untuk setiap environment haruslah terpisah dari code, Jangan sampai ada konfigurasi yang dimasukan kedalam code.
    Contoh konfigurasi yang tidak boleh dimasukan kedalam code.

    - URL/DSN database atau apapun yang menjadi dependensi ke service kita.
    - Segala jenis secret, seperti password atau credential external service.
    - environment value, seperti hotname untuk deploy.

    Umumnya konfigurasi diletakan diYAML file dan tidak dimasukan kedalam repository. Tapi ini kurang ideal. Kenapa?
    Pertama, bisa jadi tidak sengaja masuk kedalam repository. Lalu, bisa terjadi miskonfigurasi karena lupa melakukan perubahan diproduction.

    Jadi, daripada membuat konfigurasi sebagai code atau sebagai konfigurasi external. 
    The Twelve-Factor App merekomendasikan agar konfigurasi diletakan sebagai *environment variables*.

    Keuntungan menggunakan *environment variables*:
        - Menjadi standard disemua OS dan language agnostic
        - Mudah dideploy tanpa melakukan perubahan disisi code
        - Mudah untuk diinject kedalam container

4. Backing Services

    > Treat backing services as attached resources.

    Backing service seperti datastore, messaging system, SMTP, dll harus diperlakukan sebagai seuatu yang mudah digantikan.
    Jadi perubahan disetiap environment hanya perlu ubah disisi *environment variables*.

5. Build, Release, Run

    > Strictly separate build and run stages.

    Pada saat deployment, biasanya ada tiga tahap terpisah.

    - Build
        Tahap ini biasanya proses automate dalam mengambil code dari spesifik versi, mengambil dependensi dan mencompile menjadi executeable artifact.
        Setiap build biasanya dilengkapi dengan identifier yang unik seperti timestamp atau urutan angka build.
    - Release
        Ditahap ini, setelah code dibuild lalu dimasukan konfigurasi yang menuju ke spesifik target deployment(develeopment, staging, production)
    - Run
        Tahap ini, proses menjalankan ke target.

6. Processes

    > Execute the app as one or more stateless processes.

    Service harus stateless dan tidak membagi apapun, data harus disimpan di datastore bukan diaplikasi.

7. Data Isolation

    > Each service manages its own data.

    Setiap service seharusnya memiliki data sendiri dan data hanya bisa diakses melalu API yang sudah didesign.

8. Scalability

    > Scale out via the process model.

9. Disposability

    > Maximize robustness with fast startup and graceful shutdown.

    - Service harus meminimalkan waktu untuk `start up`
    - Service harus bisa shutting down ketika menerima sinya `SIGTERM`. Stop request yang masuk, Selesaikan semua proses dan tutup semua koneksi.

    Baca juga, [belajar gracefull shutdown](https://github.com/zeihanaulia/go-learn-gracefull-shutdown).

10. Development/Production Parity

    > Keep development, staging, and production as similar as possible.

    Perbedaan antara development dan production harus semirip mungkin.

    - Code divergence
        Branch development harus kecil dan short-lived, harus segera ditest dan dideploy ke production ASAP.
    - Stack divergence
        Stack yang digunakan untuk menjalankan service harus sama mulai dari jenis os, versi os, jenis datastore dan versinya,  dll
    - Personnel divergence
        Libatkan pembuat code dalam proses deployment


11. Logs

    > Treat logs as event streams.

    TODO: penjelasan

12. Administrative Processes

    > Run administrative/management tasks as one-off processes.

    TODO: penjelasan


## Learn From Books:

- https://microservices.io/ by Chris Richardson
- [Cloud Native Go](https://learning.oreilly.com/library/view/cloud-native-go/9781492076322/) by [Matthew A. Titmus](https://www.linkedin.com/in/matthew-titmus/)
- [Designing Data-Intensive Applications](https://learning.oreilly.com/library/view/designing-data-intensive-applications/9781491903063/) by [Martin Kleppmann](https://martin.kleppmann.com/)
- [Site Reliability Engineering](https://learning.oreilly.com/library/view/site-reliability-engineering/9781491929117/) by Betsy Beyer, Chris Jones, Niall Richard Murphy, Jennifer Petoff
- [Clean Architecture: A Craftsman’s Guide to Software Structure and Design](https://learning.oreilly.com/library/view/clean-architecture-a/9780134494272/) By Robert C. Martin
- [Cloud Native](https://learning.oreilly.com/library/view/cloud-native/9781492053811/) By Boris Scholl, Trent Swanson, Peter Jausovec

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

