# Summary

## (19) Unit Testing

### Software Testing
- Proses menganalisis item perangkat lunak untuk mendeteksi perbedaan antara kondisi yang ada dan yang diperlukan serta untuk mengevaluasi fitur item perangkat lunak
- Tujuan Software Testing 
    - Mencegah regresi
    - Tingkat kepercayaan dalam Refactoring
    - Tingkatkan desain kode
    - Dokumentasi Kode
    - Meningkatkan skala tim
- Level dari Testing
    - UI (End to End) :  menguji interaksi antara keseluruhan melalui antarmuka pengguna
    - Integrasi :  menguji modul atau sub sistem tertentu melalui aplikasi
    - Unit : menguji bagian terkecil yang dapat diuji dari suatu aplikasi melalui metode

### Framework Testing
- Framework menyediakan alat dan struktur yang diperlukan untuk menjalankan pengujian secara efisien, contoh Testify (Go)
- Structure pola yang biasa digunakan
    - Sentralisasi file pengujian Anda di dalam folder pengujian
    - Simpan file pengujian bersama dengan file produksi
- Runner adalah alat yang menjalankan file tes

### Mocking
- Kebutuhan Anda untuk membuat objek tiruan (dan) objek palsu yang mensimulasikan perilaku objek nyata

### Coverage
- Pembuat kode perlu memastikan apakah mereka telah membuat pengujian yang cukup
- Alat cakupan menganalisis kode aplikasi saat pengujian sedang berjalan
- Coverange Report 
    - Function
    - Statment
    - Beranch
    - Lines
- Coverage Format 
    - CLI
    - XML
    - HTML
    - LCOV

## Implementasi Testing
- Buat nama file name_test.go
- Fungsi test ditulis dengan PascalCase
- Run Testing
    ```go
    go test ./lib/calculate -cover
    ```
- Run Testing with report Coverage
    ```go
    go test ./lib/calculate -coverprofil=cover.out && go tool cover -html=cover.out
    ```

