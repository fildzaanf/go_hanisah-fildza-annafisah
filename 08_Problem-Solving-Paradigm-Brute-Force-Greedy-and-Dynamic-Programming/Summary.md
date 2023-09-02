# Summary

## (08) Problem Solving Paradigm - Brute Force, Greedy, and Dynamic Programming

### What is Problem Solving Pradigm ?
- Problem solving paradigm adalah pendekatan untuk memecahkan permasalahan
- Setiap paradigma memiliki kelebihan dan kelemahan tertentu 
- Pemilihan paradigma yang tepat tergantung pada sifat masalah yang akan dipecahkan
- Berikut beberapa paradigma yang sering digunakan dalam problem-solving.
    - Complete Search (Brute Force)
    - Divide and Conquer
    - The Greedy approach
    - Dynamic Programming

---

### Brute Force
- Complete Search (Brute Force) adalah metode memecahkan masalah dengan mencari solusi dengan memeriksa semua kemungkinan solusi secara berurutan
- Brute Force digunakan ketika tidak ada algoritma lain yang tersedia 
- Biasanya mudah ditulis karena lugas
- Kompleksitas waktu tinggi jika jumlah kemungkinan solusi besar
- Secara teoritis semua masalah dapat diselesaikan dengan menggunakan pendekatan Brute Force terutama ketika memiliki waktu tidak terbatas
- Contoh : mencari nilai minimal dan maksimal, etc.

---

### Divide and Conquer
- Divide and Conquer adalah paradigma pemecahan masalah dimana suatu masalah dibuat lebih sederhana dengan 'membagi' menjadi submasalah yang lebih kecil
- Setiap submasalah diselesaikan secara terpisah dan hasilnya digabungkan untuk mendapatkan solusi akhir
- Berikut adalah tahapan Divide and Conquer.
    - Divide : membagi masalah yang besar menjadi masalah yang lebih kecil
    - Conquer : ketika masalah sudah cukup kecil untuk diselesaikan maka langsung selesaikan
    - Combine : jika dibutuhkan maka perlu menggabungkan solusi dari submasalah yang lebih kecil menjadi solusi untuk masalah yang besar
- Contoh : algoritma pencarian Binary search, etc.

---

### The Greedy approach
- Pendekatan ini memilih langkah terbaik yang tersedia pada setiap langkah agar mencapai solusi yang optimal secara global
- Efektif dan efisien, tetapi tidak selalu menghasilkan solusi optimal
- Algoritma ini digunakan dalam masalah optimasi yang memungkinkan pemilihan langkah terbaik secara lokal
- Contoh : algoritma Huffman Coding, Activity Selection, Djikstra algorithm, etc.

---

### Dynamic Programming
- Dynamic Programming adalah paradigma problem solving yuntuk memecahkan masalah kompleks dengan memecahkannya menjadi serangkaian submasalah yang lebih kecil dan menghindari perhitungan berulang dengan menyimpan hasil dari setiap submasalah
- Dynamic Programming digunakan untuk mengatasi masalah dengan struktur rekursif atau tumpang tindih
- Berikut adalah karakteristik Dynamic Programming.
    - overlapping subproblem
        - Karakteristik ini mengacu pada sifat masalah yang memiliki submasalah yang rekursif atau tumpang tindih
        - Beberapa submasalah perlu dipecahkan lebih dari sekali
        - Solusi dari submasalah yang sama dapat digunakan kembali
        - Dynamic programming menyimpan hasil perhitungan submasalah untuk menghindari perhitungan berulang
    - optimal substructure property
        - Solusi optimal dari masalah besar dapat didapatkan dari solusi optimal submasalah yang lebih kecil
        - Struktur masalah memungkinkan penggunaan solusi optimal submasalah untuk mencapai solusi optimal secara global
- Berikut adalah metode pada Dynamic Programming.
    - Top-Down with Memoization
        - Penyelesaian masalah rekursif dengan menyimpan hasil perhitungan submasalah dalam memoization table (tabel memoization)
        - Ketika submasalah ditemukan lagi, solusinya dapat ditemukan pada tabel daripada melakukan perhitungan ulang
        - Mengurangi kompleksitas waktu dengan menghindari perhitungan berulang
    - Bottom-Up with Tabulation
        - Menyelesaikan submasalah terkecil terlebih dahulu dan menyimpan hasilnya dalam tabel
        - Solusi untuk masalah yang lebih besar dibangun berdasarkan hasil yang telah disimpan dalam tabel
        - Efisien dalam penggunaan memori dan dapat diterapkan pada masalah dengan pola tertentu
- Contoh : algoritma Fibonacci dengan penyimpanan hasil perhitungan sebelumnya, etc.
