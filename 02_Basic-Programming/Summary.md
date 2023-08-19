# Summary 

## Basic Programming

### Introduction to Golang
- History of Golang <br>
  Go is a new, general-purpose programming language that makes it easy to build, simple, reliable, and efficient software.

- Installation <br>
  - Go to https://golang.org/dl/ and download the latest version
  - Follow Instruction Installation

- Go command terminal
  - go run : Execute the program with extension .go
  - go build : Compile program files.
  - go install : Like go build and continue the install process
  - go test : For testing with suffix _test.go
  - go get : download the Go package
  - go mod init : go mod init command creates a go.mod file to track code's dependencies

### Basic Programming with Golang 
- Syntax <br>
  Main syntax 
  - Package declaration
  - Import packages
  - Functions
  - Statements and expressions

```go
// main package
package main
// import package “fmt” 
import ("fmt")
// main function golang
func main() {
  // statements and expressions
  fmt.Println("Hello World!")
}
```

 - Output <br>
Go memiliki tiga fungsi untuk menampilkan output:
   - Print()<br>
    Fungsi Print() mencetak argumennya dengan format defaultnya
   - Println()<br>
    Fungsi Println() mirip dengan Print() dengan perbedaan whitespace ditambahkan di antara argumen dan newline ditambahkan di bagian akhir
   - Printf()<br>
    Fungsi Printf() pertama memformat argumennya berdasarkan kata formatting verbs yang diberikan dan kemudian mencetaknya. Berikut beberapa contoh formatting verbs.
     - %v untuk mencetak nilai argumen
     - %T untuk mencetak jenis argumen
     - %s untuk mencetak nilai string 
     - %d untuk mencetak nilai integer 

- Variables <br>
Variabel adalah wadah untuk menyimpan nilai data <br>
  - Declaring (Creating) Variables <br>

```go
// with the var keyword
  var variablename type = value
// with the := sign 
  variablename := value
```
- Data Types <br>
  Tipe data menentukan ukuran dan tipe nilai variabel. Go statically typed artinya setelah sebuah tipe variabel didefinisikan, ia hanya dapat menyimpan data dari tipe tersebut. Go memiliki tiga tipe data dasar:
  - bool : nilai boolean (true or false). nilai default tipe data boolean adalah false
  - numeric : tipe integer, floating point, dan tipe kompleks
  - string : nilai string. tipe data string digunakan untuk menyimpan urutan karakter (teks). nilai string harus diapit oleh tanda kutip ganda

- Operators <br>
Operator digunakan untuk melakukan operasi pada variabel dan nilai. Go membagi operator menjadi beberapa grup berikut.
  - Arithmetic operators <br>
  Operator aritmatika digunakan untuk melakukan operasi matematika umum. Berikut beberapa contoh operator aritmatika.
   > +	, -, *, /, %, ++, --

  - Assignment operators <br>
  Operator assignment digunakan untuk menetapkan nilai ke variabel. Berikut beberapa contoh operator sssignment.
  > =, +=, -=, *=, /=, %=, &=, >>=, <<=, >>= 

  - Comparison operators <br>
  Operator comparison digunakan untuk membandingkan dua nilai. Berikut beberapa contoh operator comparison.
  > ==, !=, >, <, >=, <=

  - Logical operators<br>
  Operator logical digunakan untuk menentukan logika antara variabel atau nilai. Berikut beberapa contoh operator logical.
  > ==, ||, !

  - Bitwise operators
Operator bitwise digunakan pada angka (biner). Berikut beberapa contoh operator bitwise.
  > &, |, <<, >>


- Constans <br>
variabel harus memiliki nilai tetap yang tidak dapat diubah dan read-only
```go
const CONSTNAME type = value
```
### Control Structures 

- Branching <br>
Conditional statements digunakan untuk melakukan tindakan yang berbeda berdasarkan kondisi yang berbeda. Kondisi dapat berupa true atau false. Go memiliki conditional statements berikut.<br>
  - If untuk menentukan blok kode yang akan dieksekusi, jika kondisi yang ditentukan benar
  ```go
  if condition {
  // code to be executed if condition is true
  }
  ```
  - Else untuk menentukan blok kode yang akan dieksekusi, jika kondisi yang sama salah
  ```go
  if condition {
  // code to be executed if condition is true
  } else {
  // code to be executed if condition is false
  }
  ```
  - Else if untuk menentukan kondisi baru yang akan diuji, jika kondisi pertama salah
  ```go
  if condition1 {
   // code to be executed if condition1 is true
  } else if condition2 {
   // code to be executed if condition1 is false and condition2 is true
  } else {
   // code to be executed if condition1 and condition2 are both false
  }
  ```
  - Nested if untuk pernyataan if di dalam pernyataan if
  ```go
  if condition1 {
   // code to be executed if condition1 is true
  if condition2 {
     // code to be executed if both condition1 and condition2 are true
  }
  }
  ```
  - Switch case untuk menentukan banyak blok kode alternatif yang akan dieksekusi
  ```go
   // single case
   switch expression {
  case x:
   // code block
  case y:
   // code block
  case z:
  ...
  // opsional
  default:
   // code block
  }

  // multi case
  switch expression {
  case x,y:
   // code block if expression is evaluated to x or y
  case v,w:
   // code block if expression is evaluated to v or w
  case z:
  ...
  default:
   // code block if expression is not found in any cases
  }
  ```
- Looping<br>
  - For Loop<br>
    For loop mengulang melalui blok kode beberapa kali. For loop adalah satu-satunya perulangan yang tersedia di Go. Loop berguna jika Anda ingin menjalankan kode yang sama berulang kali, setiap kali dengan nilai yang berbeda. Setiap eksekusi loop disebut iterasi.
  ```go
    for statement1; statement2; statement3 {
     // code to be executed for each iteration
    } 
  ``` 
     - statement1 : Menginisialisasi nilai penghitung loop
     - statement2 : Mengevaluasi untuk setiap iterasi loop. Jika hasilnya TRUE , pengulangan berlanjut. Jika bernilai FALSE, loop berakhir
     - statement3 : Meningkatkan nilai penghitung loop
     - continue Statement :  melewatkan satu atau lebih iterasi dalam loop. Kemudian dilanjutkan dengan iterasi berikutnya dalam loop
     - break statement : menghentikan/menghentikan eksekusi loop
     - nested loops : menempatkan loop di dalam loop lain. "inner loop" akan dieksekusi satu kali untuk setiap iterasi "outer loop"
     - range : lebih mudah mengulang array, irisan, atau peta. mengembalikan indeks dan nilainya
  
  ```go
     for index, value := array|slice|map {
   // code to be executed for each iteration
    }
  ```

    

