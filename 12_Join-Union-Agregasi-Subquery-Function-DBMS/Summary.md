# Summary

## (12) Join, Union, Agregasi, Subquery, Function - DBMS

### Join
- Join adalah klausa untuk mengkombinasikan record dari dua atau lebih table
- Berikut adalah tipe SQL Join.
    - Inner Join
        - Inner Join menggabungkan baris yang memiliki nilai yang cocok di kedua tabel
        - Ini menghasilkan hanya baris yang memiliki pertandingan dalam kedua tabel
        ```sql
        SELECT column_name(s)
        FROM table1
        INNER JOIN table2
        ON table1.column_name = table2.column_name;
        ```
    - Left Join
        - Left Join menggabungkan semua baris dari tabel kiri dengan baris yang cocok di tabel kanan 
        - Jika tidak ada pertandingan di tabel kanan, nilai-nilai di tabel kanan akan menjadi NULL
        ```sql
        SELECT column_name(s)
        FROM table1
        LEFT JOIN table2
        ON table1.column_name = table2.column_name;
        ```
    - Right Join
        - Right Join adalah kebalikan dari Left Join
        - Ini menggabungkan semua baris dari tabel kanan dengan baris yang cocok di tabel kiri
        - Jika tidak ada pertandingan di tabel kiri, nilai-nilai di tabel kiri akan menjadi NULL
        ```sql
        SELECT column_name(s)
        FROM table1
        RIGHT JOIN table2
        ON table1.column_name = table2.column_name;
        
        ```
---

### Union
- Union adalah operasi dalam SQL yang digunakan untuk menggabungkan hasil dari dua atau lebih pernyataan SELECT menjadi satu hasil tunggal
        ```sql
        SELECT column1, column2 FROM table1
        UNION
        SELECT column1, column2 FROM table2;
        ```
---

### Agregasi
- Fungsi Agregasi adalah fungsi di mana nilai beberapa baris dikelompokkan bersama untuk membentuk nilai ringkasan tunggal
- Berikut beberapa fungsi Agregasi SQL.
    - MIN and MAX
        ```sql
        SELECT MIN(column_name)
        FROM table_name
        WHERE condition;

        SELECT MAX(column_name)
        FROM table_name
        WHERE condition;
        ```

    - SUM
        ```sql
        SELECT SUM(column_name)
        FROM table_name
        WHERE condition;
        ```

    - AVG
        ```sql
        SELECT AVG(column_name)
        FROM table_name
        WHERE condition;
        ```

    - COUNT
        ```sql
        SELECT COUNT(column_name)
        FROM table_name
        WHERE condition;
        ```

    - HAVING
        ```sql
        SELECT column_name(s)
        FROM table_name
        WHERE condition
        GROUP BY column_name(s)
        HAVING condition
        ORDER BY column_name(s);
        ```

---

### Subquery
- Subquery atau Inner query atau Nested query adalah query di dalam query SQL lain
- Subquey digunakan untuk mengembalikan data yang akan digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil
        ```sql
        (SELECT [DISTINCT] subquery_select_argument
        FROM {table_name | view_name}
        {table_name | view_name} ...
        [WHERE search_conditions]
        [GROUP BY aggregate_expression [, aggregate_expression] ...]
        [HAVING search_conditions])
        ```
---

### Function 
- Function adalah sebuah kumpulan statement yang akan mengembalikan sebuah nilai balik pada pemanggilnya
        ```sql
        DELIMITER //

        CREATE FUNCTION function_name [ (parameter datatype [, parameter datatype]) ]
        RETURNS return_datatype

        BEGIN

        declaration_section

        executable_section

        END;

        DELIMITER ;
        ```

        ```sql
        CREATE TRIGGER trigger_name
        {BEFORE | AFTER} {INSERT | UPDATE| DELETE }
        ON table_name FOR EACH ROW
        trigger_body;
        ```
