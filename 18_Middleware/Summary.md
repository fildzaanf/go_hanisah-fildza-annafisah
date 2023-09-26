# Summary

## (17) Middleware

### Middleware
- Middleware adalah perangkat lunak atau kode yang digunakan untuk memproses permintaan HTTP di antara permintaan masuk dan respon yang keluar dari server web atau perangkat lunak
- Middleware berfungsi sebagai perantara yang dapat melakukan berbagai tugas seperti autentikasi, autorisasi, logging, kompresi data, caching, atau modifikasi permintaan dan respons sebelum mencapai handler utama perangkat lunak
- Example third party Middleware :
    - Negroni
    - Echp
    - Interpose
    - Alice


### Echo Middleware
- Basic Auth
- Body Dump
- Body Limit
- CORS
- CSRF
- Casbin Auth
- Gzip
- JWT
- Key Auth
- Logger
- Method Override
- Proxy
- Recover
- Redirect
- Request ID
- Rewrite
- Secure
- Session
- Static
- Trailing Slash


### Implementation Middleware
- Berikut contoh penggunaan Middleware
    ```go

    e := echo.New()

    // Middleware global
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Middleware untuk grup rute tertentu
    adminGroup := e.Group("/admin")
    adminGroup.Use(AuthMiddleware)

    // Rute
    e.GET("/", Handler)
    adminGroup.GET("/dashboard", AdminDashboardHandler)
    ```
-  Pre dan Use dalam Middleware
    - Pre untuk melakukan tugas yang harus dieksekusi sebelum middleware lainnya, seperti validasi awal atau penyaringan permintaan
    - Use untuk middleware global yang harus diterapkan pada semua rute atau middleware yang harus dijalankan setelah Pre middleware
    
        ```go
        e := echo.New()

        // Middleware ini dijalankan sebelum middleware yang ditetapkan dengan Use
        e.Pre(PreMiddleware) 

        // Middleware ini dijalankan setelah middleware yang ditetapkan dengan Pre
        e.Use(GlobalMiddleware) 
        ```
- Log Middleware untuk mencatat informasi terkait permintaan HTTP yang masuk dan respons yang keluar dari aplikasi web
    ```go
    package middlewares

    import (
        "github.com/labstack/echo/v4"
        "github.com/labstack/echo/v4/middleware"
    )

    func Logger(e *echo.Echo) {

        e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
            Format: "method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
        }))

    } 
    ```
- JWT (JSON Web Token) Middleware untuk mengamankan rute atau endpoint dalam aplikasi web dengan memverifikasi token JWT yang dikirim oleh klien
    ```go
    package middlewares

    import (
        "Praktikum/constants"
        "time"

        "github.com/golang-jwt/jwt/v4"
    )

    func CreateToken(userId int, name string) (string, error) {

        payload := jwt.MapClaims{}
        payload["userId"] = userId
        payload["name"] = name
        payload["exp"] = time.Now().Add(time.Hour * 24).Unix()

        token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

        return token.SignedString([]byte(constants.SECRET_JWT))
    }
    ```
- Auth Middleware untuk melakukan autentikasi pengguna sebelum memberikan akses ke rute atau endpoint tertentu
    ```go
    package middlewares

    import (
        "github.com/labstack/echo/v4"
        "github.com/labstack/echo/v4/middleware"
        "Praktikum/constants"
        "net/http"
    )

    func AuthMiddleware() echo.MiddlewareFunc {
        
        return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
           
            if username == constants.ValidUsername && password == constants.ValidPassword {
                return true, nil
            }
            return false, nil
        })
    }

    
    ```