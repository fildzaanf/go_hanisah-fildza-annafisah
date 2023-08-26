# Summary

## (04) Time Complexity and Space Complexity

### Big-O Notation
- Big-O Notation digunakan untuk mengukur tingkat kompleksitas suatu algoritma
- Kompleksitas suatu algoritma dibagi menjadi 2, yaitu Time Complexity dan Space Complexity

### Time Complexity
- Time Complexity adalah seberapa waktu yang diperlukan oleh suatu algoritma untuk menyelesaikan suatu masalah berdasarkan input yang diberikan. 
- Berikut beberapa macam Time Complexity.
      - O(1) — Constant Time <br>
         Constant Time artinya jumlah input yang diberikan pada suatu algoritma tidak akan mempengaruhi waktu proses (runtime) dari algoritma tersebut.
         ```go
         func dominant (n int) int {
            var result int = 0
            result = result + 10
            return result
         }
         ```
      - O(log n) — Logarithmic Time
         Logarithmic Time artinya ketika memberikan input sebesar n terhadap sebuah fungsi, jumlah tahapan yang dilakukan oleh fungsi tersebut berkurang berdasarkan suatu faktor.
         ```go
         func logarithmic(n int)int {
            result := 0
            for n > 1 {
               n /= 2
               result += 1
            }
            return resulr
         }
         ```
      - O(n+m) — Linear Time 
         ```go
         func linear (n int, m int) int {
            result := 0
            for i := 0; i < n; i++ {
               result += 1
            }
            for j := 0; j < m; j++ {
               result += 1
            }
            return result
         }
         ```
      - O(n) — Linear Time 
         Linear Time artinya ketika proses (runtime) dari fungsi berbanding lurus dengan jumlah input yang diberikan.
         ```go
         func dominant (n int) int {
            result := 0

            for i := 0; i < n; i++ {
               result += 1
            }
            result = result + 10
            return result
         }
         
         ```
      - O(n^2) — Quadratic Time
         Quadratic Time adalah ketika proses (runtime) dari fungsi adalah sebesar n^2, dengan n adalah jumlah input dari fungsi tersebut.
         ```go
         func quadratic(n int) int {
         result := 0
            for i := 0; i < n; i++ {
               for j := 0; j < n; j++ {
               result += 1
               } 
            }
            return result
         }
         ```
      - O(2^n) — Exponential Time
         Exponential Time digunakan dalam situasi  tidak terlalu tahu terhadap permasalahan yang dihadapi, sehingga mengharuskan  mencoba setiap kombinasi dan permutasi dari semua kemungkinan.
      - O(n!) — Factorial Time
      
### Space Complexity
- Space Complexity adalah seberapa efisien suatu algoritma dalam penggunaan memori ketika menyelesaikan suatu masalah berdasarkan ukuran input yang diberikan.
- Berikut beberapa macam Space Complexity.
      - O(1) — Constant Space
      - O(log n) — Logarithmic Space
      - O(n) — Linear Space 
      - O(n+m) — Linear Space 
      - O(n^2) — Quadratic Space 
      - O(2^n) — Exponential Space
      - O(n!) — Factorial Space

