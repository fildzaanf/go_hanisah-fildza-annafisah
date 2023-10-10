# Summary

## (23) CI/CD

### Continuous Integration (CI)
- Continuous integration merupakan praktik pengembangan perangkat lunak di mana Developer secara teratur mengunggah (push) atau menggabungkan (merge) perubahan kode (code changes) mereka ke sebuah repositori terpusat (central repository) atau ke mainline trunk (seperti branch master/main), setelah itu proses pembentukan (build) dan pengujian (test) secara otomatis akan dijalankan
- Tujuan utama dari continuous integration untuk menemukan dan mengatasi bug lebih cepat, meningkatkan kualitas perangkat lunak, serta mengurangi waktu yang diperlukan untuk memvalidasi dan merilis pembaruan perangkat lunak baru
- Ada beberapa kategori tools yang bisa digunakan untuk menghadirkan CI untuk aplikasi yang dibuat, antara lain sebagai berikut.
    - Source control version management : GIT
    - Build automation : Jenkins
    - Automated testing : Jenkins

### Continuous Delivery/Deployment (CD)
- Continuous Delivery
    - Continuous Delivery merupakan kelanjutan dari continuous integration karena secara otomatis melakukan deploy terhadap semua perubahan kode ke testing atau production environment sesaat setelah tahapan build dan test selesai
    -  Jika diterapkan dengan benar maka developer akan selalu memiliki produk build yang siap deploy yang telah melewati proses pengujian standar 
    - Continuous delivery menggunakan pendekatan di mana suatu persetujuan manual (manual approval) perlu dilakukan sebelum proses deploy ke production environment guna memastikan bahwa perubahan kode siap disajikan ke pengguna
- Continuous Deployment
    - Continuous deployment adalah praktik yang berupaya mengotomatiskan deployment produk dari awal hingga akhir
    - Dengan praktik ini, semua perubahan yang berhasil melewati setiap fase pada alur CI/CD akan langsung di-deploy ke production environment dan tersaji ke pengguna
    - Hanya pengujian yang gagal yang akan mencegah perubahan baru di-deploy ke production environment

### CI/CD Pipeline 
-  CI/CD Pipeline adalah praktik yang berfokus pada peningkatan software delivery di seluruh siklus pengembangan perangkat lunak (build, test, dan deploy kode) melalui automasi
- Alur CI/CD pada dasarnya memiliki 8 tahapan atau fase, antara lain <br>
Plan -> Code -> Build -> Test -> Release -> Deploy -> Operate -> Monitor