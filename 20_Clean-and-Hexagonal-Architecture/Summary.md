# Summary

## (20) Clean and Hexagonal Architecture

### Clean Architecture
- Clean Architecture adalah pendekatan desain perangkat lunak yang menekankan pemisahan antara berbagai lapisan dalam aplikasi
- Tujuan dalam mengimplementasi Clean Architecture :
    - Independent of Frameworks
    - Testable
    - Independent of UI
    - Independent of Database
    - Independent of any External
- Clean Architecture memiliki empat lapisan utama, yaitu Entities, Use Cases, User Interface, dan Data Sources
- Entities:
    - Lapisan inti yang berisi definisi entitas bisnis dan logika bisnis murni
    - Lapisan ini tidak boleh bergantung pada lapisan lainnya dan harus stabil
- Use Cases:
    - Lapisan ini berisi implementasi dari use case atau alur kerja yang merepresentasikan operasi bisnis
    - Lapisan ini memproses input dari user interface dan mengembalikan hasilnya
- User Interface:
    - Lapisan ini mengatur interaksi antara pengguna dan aplikasi, termasuk antarmuka pengguna (UI) atau antarmuka program (API)
    - Lapisan ini mengambil input dari pengguna, meneruskannya ke lapisan use case, dan menampilkan hasilnya
- Data Sources:
    - Lapisan ini mengatur akses ke sumber data eksternal seperti database, layanan web, atau penyimpanan file
    - Lapisan ini menghubungkan aplikasi dengan sumber data eksternal
- Clean Architecture bertujuan untuk menciptakan perangkat lunak yang lebih mudah dimengerti, dipelihara, dan dioptimalkan dengan menghindari ketergantungan yang salah arah dan mengikuti prinsip-prinsip desain yang baik

### Hexagonal Architecture
- Hexagonal Architecture adalah pendekatan desain perangkat lunak yang bertujuan untuk menciptakan aplikasi yang mudah dipelihara, independen dari detail teknis, dan memudahkan pengujian
- Hexagonal Architecture menempatkan lapisan bisnis atau domain sebagai inti dari aplikasi. Logika bisnis utama ditempatkan di lapisan ini
- Konsep "port" digunakan untuk mendefinisikan antarmuka yang berfungsi sebagai kontrak untuk komunikasi dengan dunia luar
- konsep "Adapter" adalah implementasi dari kontrak ini dan bertanggung jawab untuk menghubungkan aplikasi dengan komponen eksternal seperti database atau antarmuka pengguna

### Context Golang
- Package context adalah suatu package yang membawa deadline, cancellation signal, atau value lain pada request/permintaan API

### Write Unit Test
- Buat berkas test terpisah untuk setiap package yang ingin diuji
- Berkas pengujian harus memiliki name_test.go
- Impor package 'testing; dalam berkas test
- Tulis fungsi yang dimulai dengan kata 'Test' diikuti dengan nama fungsi dan menerima parameter *testing.T
- Gunakan fungsi pengujian seperti t.Errorf untuk menandakan kesalahan 
- Gunakan perintah go test untuk menjalankan test