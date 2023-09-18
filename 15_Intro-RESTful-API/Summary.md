# Summary

## (15) Intro RESTful API

### What's API ?
- API (Application Programming Interface) adalah  ekumpulan aturan dan protokol yang digunakan oleh perangkat lunak untuk berkomunikasi satu sama lain
- Client mengirim request, server memprosesnya dan mengirimkan respons
- Ini memungkinkan aplikasi berbeda untuk bekerja bersama serta berbagi data
- API Tools 
    - Katalon
    - Apache JMeter
    - SoapUI
    - Postman


### REST
- REST adalah singkatan dari Representational State Transfer
- REST (Representational State Transfer) adalah pendekatan arsitektur aplikasi untuk mengembangkan layanan web
- Ini berfokus pada sumber daya yang dapat diakses melalui URL dengan operasi HTTP 
- REST adalah stateless, mendukung berbagai representasi data, dan sesuai untuk sistem terdistribusi 
- Request dan Response format 
    - JSON
    - XML
    - SOAP
- HTTP Request Method
    - GET
    - POST 
    - PUT 
    - DELETE 
    - HEAD 
    - OPTION
    - PATCH

### JSON
- JSON (JavaScript Object Notation) adalah format data dalam teks yang digunakan untuk pertukaran data antar aplikasi
- Data dalam JSON dinyatakan dalam bentuk pasangan "key-value" 

### HTTP Respon Code
- 200 : OK
- 201 : Created
- 400 : Bad Request
- 404 : Not Found
- 401 : Unauthorized 
- 405 : Method Not Allowed 
- 500 : Internal Server Error

### Open API
- Open API (Application Programming Interface) adalah antarmuka perangkat lunak yang dibuka untuk umum yang memungkinkan developer perangkat lunak eksternal untuk mengakses dan berinteraksi dengan layanan atau sistem tertentu

### POSTMan
- Postman adalah sebuah aplikasi yang digunakan oleh pengembang perangkat lunak untuk menguji, mengelola, dan mendokumentasikan API (Application Programming Interface)
- Postman memungkinkan pengguna untuk mengirim permintaan HTTP ke API dan mengamati responsnya
- Hal ini memudahkan pengujian fungsionalitas API

### Package net/http
- Package net/http menyediakan implementasi client HTTP dan server

### Package Encoding/JSON
- berisi beberapa fungsi untuk operasi JSON 
    - Decode json to object struct
    - Decode json to map [string] interface {} & interface{}
    - Decode array json to array object
    - Encode object to json string
