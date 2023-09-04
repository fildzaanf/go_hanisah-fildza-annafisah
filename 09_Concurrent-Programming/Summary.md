# Summary

## (09) Concurrent Programming 

### Concurrent Programming
- Concurrent Programming adalah paradigma pemrograman yang memungkinkan beberapa tugas atau proses dapat dieksekusi secara bersamaan dan independen tanpa harus menunggu tugas atau proses lainnya diselesaikan terlebih dahulu
- Beberapa bahasa pemrograman menyediakan mekanisme bawaan seperti threads , goroutines, atau processes untuk mengimplementasikan Concurrent Programming
- Concurrent Programming juga terdapat tantangan tertentu, seperti race conditions dan masalah sinkronisasi

---
### Big Search Website
- Big Search website menyediakan fungsionalitas pencarian untuk berbagai jenis konten, yaitu websites, images, dan videos
- Mesin pencari ini akan mencari website sehingga menampilkan hasil websites, images, dan videos di web
- Terdapat 3 cara dalam melakukan pencarian tersebut, yaitu Sequential, Parallel, dan Concurrent

---
### Different Sequential, Parallel, and Concurrent
#### Sequential Program
- Dalam Sequential Program, tugas atau proses dieksekusi satu per satu, dalam satu alur eksekusi

#### Parallel Program
- Parallel Program memungkinkan beberapa tugas atau proses dieksekusi secara bersamaan, seringkali menggunakan beberapa prosesor (multicore)

#### Concurrent Program
- Concurrent Program memungkinkan beberapa tugas atau proses dieksekusi secara independen dan bersamaan dalam eksekusinya, memberikan tampilan eksekusi bersamaan 

---

### Feature GO
#### Goroutines (Concurrent execution)
- Goroutine adalah fungsi atau metode yang berjalan secara bersamaan (independen) dengan fungsi atau metode lain
- Cost pembuatan goroutine sangat kecil jika dibandingkan dengan thread
- Thread adalah proses yang ringan, atau thread adalah unit yang mengeksekusi kode di bawah program
- Go concurrency (goroutines) mempermudah pembuatan program paralel yang memanfaatkan mesin dengan banyak prosesor
- Gomaxprocs digunakan untuk mengontrol jumlah thread sistem operasi yang tersedia untuk eksekusi goroutine pada program go tertentu

```go
    time GOMAXPROCS=1 go run main.go
```
```go
    package main 
    import (
        "fmt"
        "time"
    )

    func hello(){
        fmt.Println("Hello Fildza")
    }

    // the main function is goroutine
    func main(){
        // goroutine with keyword go
        go hello()
        fmt.Println("main function")
        time.Sleep(1 * time.Second)
    }

```

```go
    package main

    import (
        "fmt"
        "time"
    )

    func hello(s string) {
        for i := 0; i < 5; i++ {
            time.Sleep(100 * time.Millisecond)
            fmt.Println(s)
        }
    }

    func main() {
        go hello("world") // using goroutine
        hello("hello")    // using thread
    }

```

#### Channels (Synchronization and Messaging)
- Channels dalam Go memfasilitasi komunikasi dan sinkronisasi antara goroutines
- Unbuffered and buffered channels
```go
    // unbuffered
    c := make(chan string)
    
    // buffered
    c := make(chan strings, 42)

```

#### Select (Multi-way concurrent control)
- Select memudahkan untuk mengontrol komunikasi data melalui satu atau banyak saluran

---

### Race Condition
- Race condition adalah kondisi dimana dua threads digunakan oleh memori disaat bersamaan
- Berikut contoh race condition

```go
    package main

    import "fmt"

    func getNumber() int {
        var i int
        go func() {
            i = 5
        }()

        return i
}


```

---
### Fixing Data Races
#### WaitGroups
- WaitGroups merupakan cara yang lurus ke depan dalam mengatasi data race, cara kerjanya dengan memblokir akses read sampai semua operasi write telah selesai. Contoh
```go
    package main

    import (
        "fmt"
        "sync"
    )

    func getNumber() int {
        var i int
        var wg sync.WaitGroup
        wg.Add(1)
        go func() {
            i = 5
            wg.Done()
        }()
        wg.Wait()
        return i
    }
```

#### Channels Blocking
- Channels Blocking mengizinkan goroutines untuk melakukan sinkronisasi dengan yang lain untuk sesaat. Contoh
```go
    package main

    import "fmt"

    func getNumber() int {
        var i int
        done := make(chan struct{})
        go func() {
            i = 5
            done <- struct{}{}
        }()
        <-done
        return i
    }
    func main(){
        fmt.Println(getNumber())
    }
```

#### Mutex
- Mutex mengabaikan urutan untuk reads dan writes dan hanya mengharuskan tidak terjadinya secara serentak
```go
    package main

    import (
        "fmt"
        "sync"
        "time"
    )

    func main(){
        fmt.Println(getNumber())
    }
    type safeNumber struct {
        var int
        m   sync.Mutex
    }
    func (i *SafeNumber) Get() int {
        i.m.Lock()
        defer i.m.Unlock()
        return i.val
    }

    func (i *SafeNumber) Set(val int) {
        i.m.Lock()
        defer i.m.Unlock()
        i.val = val
    }

    func getNumber() int{
        i := &SafeNumber{}
        go func(){
        i.Set(5)
        }()
        time.Sleep(time.Second)
        return i.Get()
    }


```