
## Daftar Isi
- [Tujuan](#tujuan)
- [Penggunaan](#batasan-masalah)
- [Batasan Masalah](#penggunaan)

## Tujuan
* Tujuan dibuat aplikasi ini adalah untuk belajar Golang.
* Juga untuk membuat label di setiap direktori, sehingga memudahkan untuk pencarian


## Batasan Masalah
* Program ini hanya bisa mencari 1 level direktori dari direktori awal/direktori root

## Penggunaan
* Buka config.env
```
ROOT_DIRECTORY=D:\\Music
LABEL=[
JASS
ROCK
POP
]
```
`ROOT_DIRECTORY` merupakan lokasi awal pencarian
`LABEL` daftar label yang bisa dicari

* Masuk di direktori level 1 setelah `D:\\Music`
* Buat Folder
* Didalam Folder tersebut tambahkan file label.config, yang berisi label yang diinginkan
```
JASS
ROCK
```
![directory](https://image.ibb.co/dtD4Wy/tree.png)

* Jalankan Program label-explorer.exe
* Pilih salah satu tipe pencarian (Ada 2 tipe untuk mencari (SINGLE, MULTIPLE))
* SINGLE, hanya berbentuk selection.
* MULTIPLE, berbentuk string dan dipisahkan dengan satu koma (POP,ROCK).
  untuk keluar dari program jika berada dalam mode pencarian MULTIPLE ketik
    ```
    \q lalu Enter
    ```
  dan untuk membersihkan layar bisa menggunakan
    ```
    \r lalu Enter
    ```
* Ketika sudah menemukan folder yang dicari, pilih satu, tekan Enter. Maka akan muncul di Explorer.

![directory](https://image.ibb.co/cTkJry/single_label.png)
![directory](https://image.ibb.co/nj5wdd/multiple_label.png)
