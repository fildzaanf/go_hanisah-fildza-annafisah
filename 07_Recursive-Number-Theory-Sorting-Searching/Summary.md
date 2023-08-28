# Summary

## (07) Recursive, Number Theory, Sorting, and  Searching

### Recursive
- Rekursif adalah fungsi memanggil dirinya sendiri untuk menyelesaikan masalah yang lebih besar dengan memecahkannya menjadi masalah yang lebih kecil sehingga mempersingkat kode
- Jika masalahnya cukup kecil, fungsi rekursi dapat menghasilkan jawaban langsung
- Jika masalahnya terlalu besar, fungsi tersebut akan memanggil dirinya sendiri dengan cakupan masalah yang lebih kecil
- Dengan solusi rekursif, akan lebih mudah untuk melihat dan merancang jalur penyelesaian
- Strategi rekursi : Base case dan Recurrence relations
- Contoh problem Factorial
    ```go
    func factorial (n int) int {
        if n ==1 {
        return 1
        }else {
        return n*factorial(n-1)
        }
    }
    ```

### Number Theory
- Number theory adalah salah satu cabang matematika yang mempelajari bilangan bulat (integers)
- Berikut beberapa topik dalam bidang Number theory.
    - Prime Number
        ```go
            func checkPrime(number int) bool{
            if number<2{
                return false
            }
            sqrtNumber := int(math.Sqrt(float(number)))
            for i := 2; i<=sqrtNumber; i++{
                if number%i == 0{
                    return false
                }
            }
            return true
        }
        ```
    - Greatest Common Divisor (GCD)
        ```go
        func gcd(a int, b int) int{
            if b==0{
                return a
            }
            return gcd (b, a%b)
        }
        
        ```
    - Least Common Multiple (LCM)
        ```go
        func lcm(a int, b int) int {
            return a*(b/gcd(a,b))
        }
        ```
    - Factorial
         ```go
        func factorial (n int) int {
            if n ==1 {
            return 1
            }else {
            return n*factorial(n-1)
            }
        }
        ```
### Sorting
- Sorting adalah proses menyusun data dalam urutan tertentu dengan mengurutkan berdasarkan nilai elemennya
- Urutan yang paling sering digunakan adalah urutan numerik dan urutan abjad
- Beberapa contoh sorting.
    - Selection Sort - O(n^2)
        ```go
        func selectionSort(elements []int) []int {
        n := len (elements)
        for k := 0; k < n; k++ {
        minimal := k
        for j := k + 1; j < n; j++ {
        if elements[j] < elements [minimal] {
        minimal = j
            }
        }
        elements[k], elements[minimal] = elements [minimal], elements[k]
        }
        return elements
        }
        ```
    - Counting Sort - O(n + k)
        ```go
        // CountingSort Time: O(n+k) , Space: 0(n+k)
        func countingSort(elements []int, k int) []int {
        count := make([]int, k + 1) 
        for i := 0; i < len (elements); i++{
        count[elements[i]]1++
        }
        counter := 0
        for i := 0; 1 < k + 1; i++ {
            for j := 0; j < count[i]: j++ {
        elements [counter] = i
        counter += 1
            }
        }
        return elements
        }
        ```
    - Merge Sort - O(log n)
    - Builtins Sort in Golang
        ```go
        strs := []string{"c", "a", "b"} 
        sort.Strings(strs) 
        fmt.PrintIn("Strings:", strs)
        // An example of sorting
        ints := []int{7, 2, 4} 
        sort.Ints(ints) 
        fmt. PrintIn("Ints:", , ints)
        // We can also use sort to check if a slice is
        // already in sorted order
        s := sort. IntsAreSorted(ints) 
        fmt .PrintIn("Sorted: ", s)
        ```


### Searching
- Searching adalah proses menemukan posisi nilai tertentu dalam daftar nilai
- Berikut beberapa contoh searching.
    - Linear Search - O(n)
        ```go
        func linierSearch (elements []int, x int) int {
            for i := 0; i < len (elements); i++ {
            if elements[i] == x {
            return i
                }
            }
            return -1
        }
        ```
    - Builtins Search
        ```go
        elements := []int {12, 15, 15, 19, 24, 31, 53, 59, 60}
        value = 31
        findIndex := sort.SearchInts (elements, value)
        if value == elements [findIndex] {
        fmt. PrintIn("value found in elemenets")
        } else {
        fmt. PrintIn("value not found in elemenets")
        }
        ```
