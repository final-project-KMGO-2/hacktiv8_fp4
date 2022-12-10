# hacktiv8_fp_4

### Final Project 4 Hacktiv8 Kelompok 2
- Tasya Gracinia
- Muhammad Khoirul Anam
- Alexander

### Deskripsi
Source code pada direktori ini merupakan hasil dari pengerjaan final project 4 studi independen Hacktiv8 - Golang for Backend Engineer kelompok 2. Tugas yang diberikan adalah membuat API sebuah aplikasi belanja yang dapat memproses produk, transaksi, dan juga kategori.

Proyek ini dideploy menggunakan DigitalOcean pada link berikut :

https://hacktiv8-fp-4.gdlx.live/

### Dokumentasi
Untuk melihat dokumentasi postman proyek ini, link berikut dapat diakses :

https://documenter.getpostman.com/view/16732840/2s8YzTSh1j

### Pembagian tugas
Pada pengerjaan final project ini, dilakukan pembagian tugas sebagai berikut :

- Tasya Gracinia
    - Membuat flow category
- Muhammad Khoirul Anam
    - Membuat flow Product
- Alexander
    - Membuat flow Transaction

### Cara Menjalankan
1. Pastikan telah menginstall PostgreSQL pada local machine, lalu buatlah database bernama `toko_belanja`.
2. Buatlah sebuah file bernama `.env` yang isinya mengikuti format pada file `.env.example` lalu isi sesuai dengan kredensial database lokal
3. Pastikan port 8083 sedang tidak digunakan oleh aplikasi lain.
4. Masuk ke direktori project ini dan jalankan perintah `go run main.go` pada terminal.
5. Akses endpoint yang tertera pada dokumentasi.
