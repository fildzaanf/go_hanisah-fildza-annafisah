# Summary

## (05) Data Structure

### Array
Array adalah struktur data yang dapat menyimpan beberapa nilai bertipe sama dalam satu variabel dengan ukuran alokasi tetap.

 - Array declaration
    ```go
    // with the var keyword
    var array_name = [length]datatype{values} // length is defined
    var array_name = [...]datatype{values} // length is inferred
    
    // with the := sign
    array_name := [length]datatype{values} // length is defined
    array_name := [...]datatype{values} // length is inferred
    ```
- Assign and access array element
    ```go
    // assign values to array elements
    array_name[index] = value
    array_name := [length]datatype{values}

    // access and print array elements
    fmt.Println(array_name[index]) 
    ```
- Initialize with array literal
    ```go
    array_name := [5]int{} //not initialized
    array_name := [5]int{1,2} //partially initialized
    array_name := [5]int{1,2,3,4,5} //fully initialized
    array_name := [5]int{1:10,2:40} // initialize only specific elements
    // 1:10 >> assign 10 to array index 1
    // 2:40 >> assign 40 to array index 2 
    ```
- Iterate array using for loop
    ```go
    primes := [5]int{2, 3, 5}

    // technique 1
    for index := 0; index < len(primes); index++ {
    fmt.Println(primes[index])
    }

    // technique 2
    for index, element := range primes {
    fmt.Println(index, "=>", element)
    }

    for _, value := range primes {
    fmt.Println(value)
    }

    // technique 3
    index := 0
    for range primes {
    fmt.Println(primes[index])
    index++
    }

    ```
- Length of an array
    ```go
    // len(array_name)
    fmt.Println(len(array_name))
    ```

### Slice
Slice adalah struktur data yang dapat menyimpan beberapa nilai bertipe sama dalam satu variabel (seperti Array) tetapi memiliki ukuran alokasi dinamis.

- Create slice
    ```go
    // long declaration
    var even_numbers []int
    fmt.Printf("elements = %v, len = %d, cap = %d\n", even_numbers, len(even_numbers), cap(even_numbers))

    // long declaration with values
    var odd_numbers = []int{1, 3, 5, 7, 9}
    fmt.Printf("elements = %v, len = %d, cap = %d\n", odd_numbers, len(odd_numbers), cap(odd_numbers))

    // short declaration with values
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Printf("elements = %v, len = %d, cap = %d\n", numbers, len(numbers), cap(numbers))

    // using make function
    var primes = make([]int, 5, 10)
    fmt.Printf("elements = %v, len = %d, cap = %d\n", primes, len(primes), cap(primes))
    ```
- Create slice from array
    ```go
    myarray := [length]datatype{values} // array
    myslice := myarray[start:end] // slice made from the array
    ```
- Length and capacity of slice
    ```go
    len() // function - returns the length of the slice (the number of elements in the slice)
    cap() // function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
    ptr // reference to an underlying array
    ```
- Append and copy slices
    ```go
    // append elements to the end of a slice using the append() function
    slice_name = append(slice_name, element1, element2, ...) 
    // to append all the elements of one slice to another slice, use the append() function
    slice3 = append(slice1, slice2...)

    // copy() function creates a new underlying array with only the required elements for the slice
    copy(destination, source)
    copy(copied_colors, colors) // copy from colors to copied_colors
    ```
- slice zero value
the zero value of slice is nil.

### Map
Map adalah struktur data yang menyimpan data berupa pasangan key dan value dimana setiap key bersifat unik.

- Map declaration
    ```go
    // long declaration
  var salary = map[string]int{}
  fmt.Println(salary)

  // long declaration with value
  var salary_a = map[string]int{"umam": 1000, "iswanul": 2000}
  fmt.Println(salary_a)

  // short declaration
  salary_b := map[string]int{}
  fmt.Println(salary_b)

  // using make
  var salary_c = make(map[string]int)
  salary_c["doe"] = 7000 // assign value
  fmt.Println(salary_c)
    ```
- Working with maps
    ```go
    // long declaration with value
    var salary_a = map[string]int{"umam": 1000, "iswanul": 2000}
    fmt.Println(salary_a, len(salary_a))

    salary_a["nabilah"] = 7000 // assign value
    fmt.Println(salary_a)

    delete(salary_a, "iswanul") // remove value by key
    fmt.Println(salary_a)

    value, exist := salary_a["umam"] // check key is exist
    fmt.Println(value, exist)

    for key, value := range salary_a { // loop over maps
    fmt.Println("->", key, value)
    ```
### Function
Fungsi adalah sepotong kode yang dipanggil dengan nama. Fungsi adalah cara mudah untuk membagi kode Anda menjadi blok-blok yang berguna. Fungsi membuat penulisan kode yang bersih, rapi, dan modular.

- Function declaration
    ```go
    func <name_function> () { <statements>}
    func main() {}

    func <name_function> () <type_return> { <statements> }
    func phi() float64 { return 3.14 }

    func <name_function> (<parameter>) { <statements> }
    func square(value int) int {
    return value * value
    }
    ```
- Automatic type assignment
    ```go
    func scale(width, height, scale int) (int, int)
    func scale(width int, height int, scale int) (int, int)
    ```