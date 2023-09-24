# Summary

## (17) ORM and Code Structure MVC

### ORM
- ORM adalah Object-Relational Mapping yang merupakan teknik yang memetakan objek dalam bahasa pemrograman berorientasi objek (seperti Go) ke tabel database relasional
- ORM memungkinkan developer untuk berinteraksi dengan database menggunakan objek pemrograman daripada kueri SQL mentah
- ORM menyederhanakan operasi database seperti pengambilan data, penyisipan, dan pembaruan, serta mengurangi kerumitan dalam mengelola data

### GORM 
- GORM adalah library Go yang populer untuk ORM (Object-Relational Mapping) yang menyederhanakan interaksi dengan database
- GORM mendukung berbagai sistem manajemen database (seperti MySQL, PostgreSQL, SQLite) dan menyediakan API tingkat tinggi untuk operasi database
- GORM menyederhanakan tugas-tugas seperti membuat, membaca, memperbarui, dan menghapus data dari database

### Code Structure MVC
- Struktur MVC adalah pendekatan arsitektur perangkat lunak yang membagi aplikasi menjadi tiga komponen utama: Model, View, dan Controller
- Model mewakili data dan logika bisnis aplikasi
- Model berinteraksi dengan database melalui ORM untuk mengelola data
- View bertanggung jawab untuk antarmuka pengguna dan logika presentasi
- View menampilkan data kepada pengguna dan menangkap input pengguna
- Controller berperan sebagai perantara antara Model dan View
- Controller memproses permintaan pengguna, berinteraksi dengan Model untuk mengelola data, dan mengirimkan respons ke View
- Struktur MVC meningkatkan modularitas kode, memungkinkan pemisahan tugas yang jelas antara komponen, dan mempermudah pemeliharaan dan skalabilitas aplikasi