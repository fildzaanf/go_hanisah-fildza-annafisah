# Summary

## (16) Intro Echo Golang

### Whats Echo 
- Echo adalah framework web ringan untuk bahasa pemrograman Go (Golang) dengan routing sederhana, dukungan middleware, dan kinerja tinggi
- Membuat Instance Echo Untuk memulai penggunaan Echo sebagai server HTTP
    ```go
    e := echo.New()
    ```
- Menentukan route HTTP dan handler fungsi yang sesuai untuk setiap rute menggunakan metode GET, POST, PUT, DELETE, dan lainnya
    ```go
    e.GET("/route", handlerFunction)
    ```
- Handler functions adalah fungsi yang akan dieksekusi ketika rute tertentu diakses
- Handler functions menerima context sebagai argumen dan dapat digunakan untuk menangani permintaan HTTP
-  Echo mendukung middleware yang dapat digunakan untuk memodifikasi  respons sebelum atau setelah melewati handler utama
- Dapat mengakses parameter dari URL dan data binding dari permintaan HTTP, seperti JSON atau data formulir, ke dalam struktur data Go
- Mengelola respons dengan mudah menghasilkan respons HTTP dengan menggunakan metode c.JSON, c.HTML, dan lainnya
- Akhiri dengan memulai server Anda dan mendengarkan pada port tertentu
    ```go
    e.Logger.Fatal(e.Start(":8000"))
    ```

### Advantage Using Echo
- Optimized Router 
- Data Rendering
- Data Binding
- Middleware
- Scalable

### How to Install Echo

```cli
$ ~/workspacename
$ go get github.com/labstack/echo/v4
```
