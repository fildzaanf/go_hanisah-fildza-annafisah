# Summary

## (11) Database Schema, DDL, DML 

### Database 
- Database adalah sekumpulan data yang terorganisir
- Database Relationship
    - One to One
    - One to Many
    - Many to Many
- Tipe data MySQL : num, huruf, dan date

### DDL (Data Definition Language)
- DDL (Data Definition Language) digunakan untuk mendefinisikan, mengelola, dan mengubah struktur objek dalam database
- Pernyataan DDL mengenai struktur data dalam database
- Berikut adalah beberapa pernyataan DDL dan fungsinya.
    - CREATE digunakan untuk membuat objek baru dalam database, seperti tabel, indeks, atau tampilan
        ```sql
        CREATE DATABASE databasename;
        USE DATABASE databasename;
        CREATE TABLE table_name (
            column1 datatype,
            column2 datatype,
            column3 datatype,
            ....
        );
        DROP TABLE table_name;
        ```
    - ALTER digunakan untuk mengubah struktur objek yang sudah ada dalam database. Alter digunakan untuk menambahkan, mengubah, atau menghapus kolom dalam tabel atau mengubah tipe data kolom
        ```sql
        ALTER TABLE table_name 
        ADD column_name datatype;

        ALTER TABLE table_name 
        DROP COLUMN column_name;

        ALTER TABLE table_name 
        ALTER COLUMN column_name datatype;
        ```
    - DROP digunakan untuk menghapus objek dari database, seperti tabel, indeks, atau tampilan
        ```sql
        DROP TABLE table_name;
        ```
    - TRUNCATE digunakan untuk menghapus semua baris dari tabel, tetapi tetap mempertahankan struktur tabel itu sendiri. Truncate lebih cepat daripada DROP karena tidak perlu membangun ulang tabel
        ```sql
         TRUNCATE TABLE table_name;
        ```
### DML (Data Manipulation Language)
- DML (Data Manipulation Language) digunakan untuk mengelola atau memanipulasi data dalam tabel dalam database
- Pernyataan DML fokus pada operasi seperti menambah, mengubah, menghapus, dan mengambil data dari tabel
- Berikut adalah beberapa pernyataan DML dan fungsinya.
    - INSERT digunakan untuk menambahkan data baru ke dalam tabel
        ```sql
        INSERT INTO table_name (column1, column2, ...)
        VALUES (value1, value2, ...);
        ```
    - UPDATE digunakan untuk mengubah data yang sudah ada dalam tabel
        ```sql
        UPDATE table_name
        SET column1 = value1, column2 = value2, ...
        WHERE condition;
        ```
    - DELETE digunakan untuk menghapus data dari tabel
        ```sql
        DELETE FROM table_name
        WHERE condition;
        ```
    - SELECT digunakan untuk mengambil data dari tabel
        ```sql
        SELECT column1, column2, ...
        FROM table_name
        WHERE condition;
        
        SELECT * FROM table_name;
        ```
    - LIKE digunakan dalam pernyataan SELECT untuk mencocokkan nilai kolom dengan pola tertentu. Ini berfungsi ketika ingin mencari data berdasarkan substring atau karakteristik tertentu dalam kolom
        ```sql
        SELECT column1, column2, ...
        FROM table_name
        WHERE columnN LIKE pattern;
        ```
    - BETWEEN digunakan dalam pernyataan SELECT untuk mencari data yang berada dalam rentang nilai tertentu dalam kolom
        ```sql
        SELECT column_name(s)
        FROM table_name
        WHERE column_name BETWEEN value1 AND value2;
        ```
    - AND digunakan dalam pernyataan WHERE untuk menggabungkan dua atau lebih kondisi. semua kondisi harus benar agar baris data dipilih
        ```sql
        SELECT column1, column2, ...
        FROM table_name
        WHERE condition1 AND condition2 AND condition3 ...;
        ```
    - OR digunakan dalam pernyataan WHERE untuk menggabungkan dua atau lebih kondisi. setidaknya salah satu kondisi harus benar agar baris data dipilih.
        ```sql
        SELECT column1, column2, ...
        FROM table_name
        WHERE condition1 OR condition2 OR condition3 ...;
        ```
    - ORDER BY digunakan dalam pernyataan SELECT untuk mengurutkan hasil berdasarkan satu atau lebih kolom (mengurutkan secara naik (ASC) atau menurun (DESC))
        ```sql
        SELECT column1, column2, ...
        FROM table_name
        ORDER BY column1, column2, ... ASC|DESC;
        ```
    - LIMIT digunakan dalam pernyataan SELECT untuk membatasi jumlah baris data yang dikembalikan oleh queri
        ```sql
        SELECT * FROM table_name LIMIT 10;
        ```

### DCL (Data Control Language)
- DCL (Data Control Language) digunakan untuk mengelola hak akses dan izin dalam database
- Pernyataan DCL memungkinkan pengelola database untuk mengontrol siapa yang dapat mengakses dan melakukan operasi tertentu dalam database
- Berikut adalah beberapa pernyataan DML dan fungsinya.
    - GRANT digunakan untuk memberikan izin atau hak akses tertentu kepada pengguna dalam database
        ```sql 
        GRANT permission_type ON object_name TO user_or_role;
        ```
    - REVOKE digunakan untuk mencabut atau menghapus izin yang sebelumnya diberikan kepada pengguna dalam database
        ```sql
        REVOKE permission_type ON object_name FROM user_or_role;
        ```