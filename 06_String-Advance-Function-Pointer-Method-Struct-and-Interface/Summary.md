# Summary

## (06) String, Advance Function, Pointer, Struct, Method,  Interface, Package, and Error Handling

### String
- String adalah tipe data yang digunakan untuk merepresentasikan urutan karakter
- String diwakili dengan tanda kutip dua ("").
- Len, Compare, and Contains
    ```go
    // 1. len string
    sentence := "Hello";
    lenSentence := len(sentence)
    fmt.Println(lenSentence)

    // 2. compare string
    str1 := "abc"
    str2 := "abd"
    fmt.Println(str1 == str2)

    // 3. Contains
    res := strings.Contains(str, substr)
    fmt.Println(res) // true
    ```
- Substring, Replace, and Insert
    ```go
    // 4. substring
    value := "cat;dog"
    // Take substring from index 4 to length of string.
    substring := value[4:len(value)]
    fmt.Println(substring)

    // 5. Replace
    s := "this[things]I would like to remove"
    t := strings.Replace(s, "[", "", -1)
    fmt.Printf("%s\n", t)  

    // 6. Insert
    p := "green"
    index := 2
    q := p[:index] + "hi" + p[index:]
    fmt.Println(p, q)
    ```
### Advance Function

- Variadic Function
    - Fungsi dapat menerima jumlah argumen yang tidak terbatas. Membuat variadic function dengan menggunakan '...' setelah tipe data parameter

        ```go
        func variadicFunction(number ...int) int{

        }
        ```
- Anonymous Function
    - Fungsi tanpa nama yang dapat digunakan sebagai nilai variabel atau sebagai argumen dalam fungsi lain

        ```go
        // Anonymous function
        func() {
            fmt.Println("Welcome! to Alta")
        }()

        // Assigning an anonymous function to a variable
        value := func() {
            fmt.Println("Welcome! to Alta")
        }
        value()

        // Passing arguments in anonymous function
        func(sentence string) {
            fmt.Println(sentence)
        }("Alta")
        ```
- Closure
    - Fungsi yang dapat mengakses variabel dari lingkungan di mana mereka didefinisikan
    - Tipe khusus dari fungsi anonim yang mereferensikan variabel yang dideklarasikan di luar fungsi itu sendiri
    - Dalam hal ini akan menggunakan variabel yang tidak dimasukkan ke dalam fungsi sebagai parameter, namun tersedia saat fungsi dideklarasikan

        ```go
        func newCounter() func() int {
        count := 0
            return func() int {
            count += 1
            return count
            }
        }

        ```

- Defer Function
    - Fungsi yang diatur untuk dieksekusi setelah fungsi selesai berjalan, biasanya digunakan untuk membersihkan sumber daya atau menangani kesalahan

        ```go
        func main() {
        defer func() {
            fmt.Println("Later")
        }()

        fmt.Println("First")
        }

        // output :
        // First
        // Later
        ```

### Pointer
- Pointer adalah variabel yang menyimpan alamat memori variabel lain
- Pointer memiliki kekuatan untuk mengubah data yang ditunjuknya
- Operator pada pointer :
    - (*) operator : dereferencing 
    - (&) operator : referencing

    ```go
    Declaration : 
        var <variable_name> *<variable_type>
        var nameAddress *string 

    Use : 
        var <variable_name> *<variable_type>
    	var name = “John” 
        var nameAddress *string
        nameAddress = &name 
    ```

### Method
- Method adalah fungsi yang melekat pada suatu tipe data (struct atau tipe data lainnya)
- Method digunakan untuk melakukan operasi pada nilai yang terkait dengan tipe data tersebut
- Method membantu menulis kode OOP di Go
- Method membantu Anda menghindari konflik penamaan
- Pemanggilan method lebih mudah dibaca dan dipahami dibandingkan pemanggilan fungsi

    ```go
    // method
    func (receiver StructType) functionName(input type) returnType {
    // block statement method
    }
    // function
    func functionName(input type) returnType {
    // block statement function
    }

    ```

### Struct
- Struct adalah tipe yang ditentukan user yang berisi kumpulan bidang/properti atau fungsi (metode) bernama
- Struct adalah tipe data yang digunakan untuk mengelompokkan beberapa nilai dengan tipe data berbeda dalam satu variabel
- Untuk mengakses field struct, gunakan operator dot (.) antara nama variabel struct dan field struct

    ```go
    // declaration
    type struct_variable_name struct {
    field <data_type>
    field <data_type>
    ...
    field <data_type>
    }
    
    type Person struct {
    name string
    age int
    }

    func main() {

    // initialization
    person := Person{"John", 30}
   
    // access struct 
    fmt.Println("Name: ", person.name)
    fmt.Println("Age: ", person.age)
  }
    ```

### Interface
- Interface adalah kumpulan definisi metode yang dapat diimplementasikan ke objek
- Dalam Go, sebuah tipe data secara otomatis dianggap sebagai implementasi dari sebuah interface jika tipe data tersebut mengimplementasikan semua metode yang didefinisikan dalam interface
- Interface mendefinisikan perilaku objek
- zero value interface adalah nil

    ```go
    type interface_name interface {
    method_name1 <return_type>
    method_name2 <return_type>
    method_name3 <return_type>
    ...
    method_namen <return_type>
    }

    type calculate interface {
    large() int
    }

    type square struct {
    side int
    }

    func (s square) large() int {
    return s.side * s.side
    }

    ```
### Package
- Package adalah kumpulan fungsi dan data
- Package adalah cara untuk mengorganisir dan mengelompokkan kode Go
- File Go harus berada dalam package tertentu
- Package dapat diimpor dan digunakan oleh package lain

    ```go
    package main

    import (
    "fmt"
    "math/rand"
    )
    ```

### Error Handling
- Error handling adalah cara mengatasi kesalahan atau kondisi yang tidak diinginkan dalam kode
- Fungsi panic() digunakan untuk menghentikan eksekusi program dengan pesan kesalahan
- Ketika runtime Go mendeteksi kesalahan ini, ia menjadi panic
- Fungsi recover() digunakan untuk mengatasi panic dan melanjutkan eksekusi program

    ```go
    package main

    import "fmt"

    func myMethod() {
    defer func() {
        if err := recover(); err != nil {
        fmt.Println("Error Message:", err)
        }
    }()

    anOddCondition := true
    if anOddCondition {
        panic("I am panicking")
    }
    }

    func main() {
    myMethod()
    }
    ```


