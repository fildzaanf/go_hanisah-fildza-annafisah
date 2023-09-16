# Summary

## (14) System Design

### Diagram Design
- Diagram adalah representasi visual dari data ataupun informasi yang digunakan untuk memvisualisasikan hubungan, pola, atau struktur dalam data
- Diagram Design Tools
    - Smartdraw
    - Lucidchart
    - Whimsical
    - draw.io 
    - Microsoft Visio 
- Berikut diagram yang umum digunakan untuk design software.
    - Flowchart
        - The theory - represent 1 process
        - Process, Decision, Terminator
    - Use Case
        - Diagram use case merangkum bagaimana pengguna sistem (aktor) berinteraksi dengan sistem dengan menjelaskan rincian aksi dan respons mereka dalam interaksi tersebut
    - Entity Relationship Diagram (ERD)
        - Diagram Entity Relationship (ER) adalah diagram untuk menggambarkan hubungan antara entitas seperti orang, objek, atau konsep dalam suatu sistem dan bagaimana mereka berinteraksi satu sama lain
    - High Level Architecture (HLA)
        - overall system design

### Characteristics of Distributed Systems

#### Scalability
- Scalability adalah kemampuan sistem, proses, atau jaringan untuk berkembang dan mengelola peningkatan permintaan
- Sistem terdistribusi yang dapat terus berkembang untuk mendukung peningkatan jumlah pekerjaan dianggap dapat diskalakan
- Sistem mungkin harus diskalakan karena berbagai alasan seperti peningkatan volume data atau peningkatan jumlah pekerjaan
- Sistem yang dapat diskalakan ingin mencapai penskalaan tanpa kehilangan performa

#### Reliability
- Reliability adalah probabilitas sistem akan gagal dalam periode tertentu
- Sistem terdistribusi dianggap dapat di reliability jika sistem tersebut tetap memberikan layanannya bahkan ketika satu atau beberapa komponen perangkat lunak atau perangkat kerasnya gagal
- Reliability mewakili salah satu karakteristik utama dari setiap sistem terdistribusi

#### Availability
- Availability adalah waktu sistem tetap beroperasi untuk menjalankan fungsi yang diperlukan dalam periode tertentu
- Ini adalah ukuran sederhana dari persentase waktu suatu sistem, layanan, atau mesin tetap beroperasi dalam kondisi normal

#### Efficiency
- Efficiency adalah kemampuan sistem untuk menghasilkan hasil yang diinginkan dengan menggunakan sumber daya yang tersedia secara optimal
- Efisiensi sistem terdistribusi berkaitan erat dengan bagaimana sistem mengelola sumber daya, mendistribusikan tugas, dan mengoptimalkan proses komunikasi antarkomponennya

#### Serviceaility or Manageability
- Serviceaility or Manageability adalah kesederhanaan dan kecepatan suatu sistem dapat diperbaiki atau dipelihara
- jika waktu untuk memperbaiki sistem yang gagal bertambah, maka ketersediaan akan berkurang
- Hal-hal yang perlu dipertimbangkan untuk pengelolaan adalah kemudahan mendiagnosis dan memahami masalah, kemudahan melakukan modifikasi, dan seberapa sederhana sistem dioperasikan

### Horizontal Scaling vs Vertical Scaling

#### Horizontal Scaling (Scaling out)
- Horizontal scaling adalah meningkatkan kapasitas sistem dengan menambahkan lebih banyak mesin atau server ke dalam infrastruktur 
- Sistem dapat menangani peningkatan beban dengan mendistribusikan beban kerja ke beberapa mesin, yang membuatnya lebih mudah untuk mengatasi permintaan yang lebih besar

#### Vertical Scaling (Scaling up)
- Vertical scaling adalah meningkatkan kapasitas sistem dengan meningkatkan daya komputasi pada satu mesin atau server yang ada
- Peningkatan kapasitas perangkat keras seperti CPU, RAM, atau penyimpanan pada server tunggal untuk meningkatkan kemampuan sistem dalam menangani beban yang lebih besar

### Job/Work Queue
- Job Queue/Batch Queue adalah seperti daftar pekerjaan yang diatur oleh perangkat lunak penjadwalan untuk dieksekusi
- Work Queue adalah semacam kerangka kerja untuk membuat aplikasi besar yang mengatur pekerjaan dan pengerjaannya melibatkan ribuan komputer dari berbagai kluster, cloud, dan jaringan

### Load Balancing
- Load Balancer membantu menyebarkan lalu lintas ke sekelompok server untuk meningkatkan daya tanggap dan ketersediaan aplikasi, situs web, atau database
- Load Balancer juga melacak status semua sumber daya saat mendistribusikan permintaan
- Jika server tidak tersedia untuk menerima permintaan baru atau tidak merespons atau memiliki tingkat kesalahan yang tinggi, Load Balancer akan berhenti mengirimkan lalu lintas ke server tersebut

### Monolithic vs Microservices
#### Monolithic
- Monolithic memiliki basis kode tunggal dengan banyak modul
- Modul dibagi berdasarkan fitur bisnis atau fitur teknis
- Ini memiliki sistem build tunggal yang membangun seluruh aplikasi dan/atau ketergantungan
- Ia juga memiliki biner tunggal yang dapat dieksekusi atau diterapkan

#### Microservices
- Microservices adalah layanan yang dapat diterapkan secara independen yang dimodelkan pada domain bisnis
- Mereka berkomunikasi satu sama lain melalui jaringan
- Sebagai pilihan arsitektur, mereka menawarkan banyak pilihan untuk memecahkan masalah 
- Arsitektur microservices didasarkan pada beberapa microservices yang berkolaborasi

### SQL vs NoSQL
- ACID adalah empat sifat penting dalam pengelolaan transaksi dalam basis data:
    - Atomicity : Transaksi adalah entitas tunggal yang harus dilaksanakan sepenuhnya atau gagal sama sekali
    - Consistency : Transaksi harus menjaga konsistensi basis data sebelum dan setelah pelaksanaan, mematuhi semua aturan integritas
    - Isolation : Transaksi harus berjalan seolah-olah adalah satu-satunya transaksi yang berjalan, terisolasi dari transaksi lainnya
    - Durability : Hasil transaksi harus bertahan bahkan setelah terjadi kegagalan sistem, memastikan data tetap aman dan tidak hilang

#### SQL
- SQL mengacu pada basis data relasional
- Basis data relasional terstruktur dan memiliki skema yang telah ditentukan sebelumnya
- SQL menggunakan tabel dan memiliki skema yang telah ditentukan sebelumnya
- Data dalam SQL diorganisasi dalam tabel dengan baris dan kolom
- SQL lebih baik digunakan ketika struktur data tetap dan terdefinisi dengan baik
- DBMS yang menggunakan SQL : MySQL, PostgreSQL, Oracle

#### NoSQL
- NoSQL mengacu pada basis data non-relasional
- Basis data non-relasional tidak terstruktur dan memiliki skema dinamis
- NoSQL memiliki skema yang lebih dinamis dan fleksibel, yang memungkinkan penyimpanan data yang berbeda dalam dokumen, grafik, etc
- Data dalam NoSQL tidak harus diorganisasi dalam tabel, dan model datanya dapat lebih variatif
- NoSQL lebih cocok digunakan untuk data yang tidak memiliki struktur yang ketat atau perlu diakses secara horizontal
- DBMS yang menggunakan NoSQL : MongoDB, Cassandra, Redis

### Caching
- Caching adalah metode penyimpanan sementara data dalam memori atau penyimpanan yang lebih cepat
- Caching untuk meningkatkan kinerja dan responsivitas sistem dengan mengurangi waktu akses ke data yang sering digunakan
- Data yang sering diakses disimpan dalam cache agar akses lebih cepat

### Database Indexing
- Database indexing adalah proses pembuatan struktur data tambahan
- Index digunakan untuk mempercepat pencarian dan penyortiran data dalam basisdata
- Database indexing menyimpan data yang telah diurutkan dengan cara tertentu sehingga query dapat menemukan data lebih efisien

### Database Replication
- Database replication adalah proses menggandakan data dari satu basisdata ke basisdata lainnya secara real-time 
- Database replication bertujuan untuk redundancy, ketersediaan, dan skalabilitas
- Database replication memastikan data yang sama tersedia di berbagai server, meningkatkan ketersediaan dan keandalan data