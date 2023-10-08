# Summary

## (22) Compute Service

### Introduction to Deployment
- Deployment adalah kegiatan yang bertujuan untuk menyebarkan aplikasi/produk yang telah dikerjakan oleh para pengembang dan seringkali untuk mengubah dari status sementara ke permanen. 
- Penyebarannya dapat melalui beragam cara tergantung dari jenis aplikasinya, aplikasi web & api ke server sedangkan aplikasi mobile ke Playstore/Appstore.

### Deployment Strategy
- Bing-Bang Deployment Strategy atau sering disebut Replace Deployment Strategy
- Rollout Deployment Strategy
- Blue/Green Deployment Strategy
- Canary Deployment Strategy

#### Big-Bang / Replace Deployment Strategy
Kelebihan 
- Mudah diimplementasikan. cara klasik, tinggal replace
- Perubahan kepada sistem langsung 100% secara implementasikan

Kekurangan
- Terlalu beresiko, rata rata downtime cukup lama

#### Rollout Deployment Strategy
Kelebihan
- Lebih aman dan less downtime dari versi sebelumnya

Kekurangan 
- Akan ada 2 versi aplikasi berjalan secara barengan sampai semua server terdeploy, dan bisa membuat bingung
- Karena sifatnya perlahan satu persatu untuk deployment dan rollback lebih lama dari yang Bigbang karena prosesnya perlahan lahan sampai semua server terkena efeknya
- Tidak ada kontrol request, Server yang baru ter-deploy dengan aplikasi versi bary, langsung mendapatkan request yang sama banyaknya dengan server lain.

#### Blue/Green Deoloyment Strategy
Kelebihan :
- Perubahan sangat cepat, sekali switch service langsung berubah 100%
- Tidak ada issue beda versi pada service seperti yang terjadi pada Rollout Deployment

Kekurangan :
- Resource yang dibutuhkan lebih banyak karena untuk setiap deployment kita harus menyediakan service yang serupa environtmentnya dengan yang sedang berjalan di production
- Testing harus benar-benar sangat diprioritaskan sebelum di switch, aplikasi harus kita pastikan aman dari request yang tiba-tiba banyak

#### Canary Deployment Strategy
Kelebihan :
- Cukup aman
- Mudah untuk rollback jika terjadi error/bug, tanpa berimbas kesemua user

Kekurangan :
- Untuk mencapai 100% cukup lama dibanding dengan Blue/Green deployment
- Dengan Blue/Green deployment, aplikasi langsung 100% terdeploy keseluruh user

### Amazon Web Services (AWS)
- AWS merupakan platform cloud yang paling komprehensif dan digunakan secara luas
- Menurut AWS, cloud computing adalah penggunaan sesuai kebutuhan (on-demand) sumber daya IT melalui internet dengan harga sesuai pemakaian (pay-as-you-go)
- Terdapat 3 model penerapan  cloud computing adalah cloud-based, on-premises (lokal), dan hybrid
- Amazon Elastic Compute Cloud (Amazon EC2) adalah layanan yang dapat digunakan untuk mendapatkan akses ke server virtual 
- Amazon RDS (Relational Database Service) adalah layanan basis data relasional yang dikelola seluruhnya oleh AWS. Layanan ini menyederhanakan operasi dan pemeliharaan basis data SQL
