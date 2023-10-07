# Summary

## (21) Docker


### Container vs. Virtual Machine 
- Container
    - Container abstraction at the app layer
    - Container can handle more applications and require fewer VMs and Operating Systems
    - Containers take up less space than VM

- Virtual Machine
    - Abstraction of physical hardware
    - Also be slow to boot
    - Each VM includes a full copy of an operating system

### What is Docker ?
- Docker adalah platform open-source yang digunakan untuk mengembangkan, menguji, dan menjalankan aplikasi dalam sebuah wadah (container)
- Docker memungkinkan aplikasi dan semua dependensinya dibungkus dalam wadah yang disebut container. Hal ini membuat aplikasi dapat berjalan secara konsisten di berbagai
- Container Docker dibuat dari image yang berisi sistem file dan dependensi yang diperlukan
- Image dapat digunakan untuk membuat banyak container yang identik
- Dockerfile adalah file yang digunakan untuk membuat image Docker
- Dockerfile berisi instruksi untuk mengkonfigurasi environment aplikasi
- Docker Hub adalah registry online yang menyimpan image Docker yang dapat digunakan oleh siapapun
- Docker Compose adalah alat untuk mendefinisikan dan menjalankan aplikasi yang terdiri dari beberapa container
- Docker Compose mempermudah pengaturan aplikasi yang kompleks
- Docker memungkinkan aplikasi diisolasi dari sistem operasi host yang mana menghindari konflik dependensi dan masalah environment
- Container Docker dapat berjalan di berbagai jenis platform dan environment dari komputer lokal hingga cloud
- Container memiliki overhead yang rendah dan juga memungkinkan untuk penggunaan sumber daya yang efisien
- Mudah menyalin dan mengelola banyak instance container sesuai kebutuhan aplikasi
- Docker menyediakan mekanisme isolasi untuk mengamankan container

### Docker Basic
- Step 1: Install Docker https://docs.docker.com/get-docker/
- Step 2: Verifikasi Instalasi Docker <br>
```bash
        docker --version
        docker-compose --version
```
- Step 3: Buat Dockerfile <br>
    ```Dockerfile
        # Gunakan image base yang sesuai
        FROM golang:1.16

        # Set working directory dalam container
        WORKDIR /app

        # Salin file kode Anda ke dalam container
        COPY . .

        # Build aplikasi Go
        RUN go build -o main .

        # Port yang akan digunakan oleh aplikasi
        EXPOSE 8080

        # Command untuk menjalankan aplikasi
        CMD ["./main"]
    ```
- Step 4: Build Image <br>
```bash
    docker build -t my-go-app:1.0 .
```
- Step 5: Run Container <br>
```bash
    docker run -p 8080:8080 my-go-app:1.0

```
- Step 6: Docker Compose (Opsional) <br>
buat file docker-compose.yml
```bash
    docker-compose up
```
- Step 7: Push Image (Opsional) <br>
```bash
    docker tag image_id fildzaanf/praktikum
    docker push fildzaanf/praktikum
```
- Step 8: Deploy di Lingkungan Produksi (Opsional)