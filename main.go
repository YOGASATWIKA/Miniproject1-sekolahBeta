package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Buku struct {
	kodeBuku string
	judulBuku string
	pengarang string
	penerbit string
	jumlahHalaman int
	tahunTerbit int
	
}

var listBuku []Buku


//Fungsi menambahkan data buku
func TambahBuku() {

	inputanUser := bufio.NewReader(os.Stdin)
	judulBuku := ""
	pengarang := ""
	penerbit := ""
	jumlahHalaman := 0
	tahunTerbit := 0
	fmt.Println("=================================")
	fmt.Println("Tambah buku")
	fmt.Println("=================================")
	fmt.Print("Silahkan Masukan kode buku : ")

	kodeBuku, err := inputanUser.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return
	}

	kodeBuku = strings.Replace(
		kodeBuku,
		"\n",
		"",
		1)

	fmt.Print("Silahkan masukkan judul buku : ")
	_, err = fmt.Scanln(&judulBuku)
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return
	}

	fmt.Print("Silahkan Masukan pengarang : ")
	_, err = fmt.Scanln(&pengarang)
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return
	}
	fmt.Print("Silahkan Masukan penerbit : ")
	_, err = fmt.Scanln(&penerbit)
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return
	}
	fmt.Print("Silahkan Masukan jumlah halaman : ")
	_, err = fmt.Scanln(&jumlahHalaman)
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return
	}
	fmt.Print("Silahkan Masukan tahun terbit : ")
	_, err = fmt.Scanln(&tahunTerbit)
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return
	}

	listBuku = append(listBuku, Buku{
	kodeBuku : kodeBuku,
	judulBuku : judulBuku,
	pengarang : pengarang,
	penerbit :penerbit ,
	jumlahHalaman :jumlahHalaman ,
	tahunTerbit : tahunTerbit,
	})

	fmt.Println("Berhasil Menambah data buku!")
}


//Fungsi melihat data buku
func LihatBuku() {
	fmt.Println("=================================")
	fmt.Println("Lihat Pesanan")
	fmt.Println("=================================")
	for urutan, Tambah := range listBuku {
		fmt.Printf("%d. kode buku : %s\n   judul buku : %s\n   pengarang buku : %s\n   penerbit buku: %s\n   jumlah halaman buku: %d\n   tahun terbit buku : %d\n",
			urutan+1,
			Tambah.kodeBuku,
			Tambah.judulBuku,
			Tambah.pengarang,
			Tambah.penerbit,
			Tambah.jumlahHalaman,
			Tambah.tahunTerbit,
		)
	}

}


//Fungsi mengedit data buku(Belum selesai)
func EditBuku() {
	fmt.Println("=================================")
	fmt.Println("Edit data buku")
	fmt.Println("=================================")
	LihatBuku()
	fmt.Println("=================================")



	fmt.Println("data Buku Berhasil diedit!")
}


//fungsi menghapus data buku
func HapusBuku() {
	fmt.Println("=================================")
	fmt.Println("Hapus data buku")
	fmt.Println("=================================")
	LihatBuku()
	fmt.Println("=================================")
	var UrutanBuku int
	fmt.Print("Masukan Urutan Buku : ")
	_, err := fmt.Scanln(&UrutanBuku)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	if (UrutanBuku-1) < 0 ||
		(UrutanBuku-1) > len(listBuku) {
		fmt.Println("Urutan Buku Tidak ditemukan")
		HapusBuku()
		return
	}

	listBuku = append(
		listBuku[:UrutanBuku-1],
		listBuku[UrutanBuku:]...,
	)

	fmt.Println("data Buku Berhasil Dihapus!")

}


//fungsi utama
func main() {
	var pilihanMenu int
	fmt.Println("Sistem Manajemen Manajemen Daftar Buku Perpustakaan")
	fmt.Println("=================================")
	fmt.Println("Silahkan Inputkan pilihan anda : ")
	fmt.Println("1. Tambah Buku Baru")
	fmt.Println("2. Liat Daftar Buku")
	fmt.Println("3. Hapus Buku")
	fmt.Println("4. Edit Buku")
	fmt.Println("5. Keluar")
	fmt.Println("=================================")
	fmt.Print("Inputkan Pilihan disini : ")
	_, err := fmt.Scanln(&pilihanMenu)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	switch pilihanMenu {
	case 1:
		TambahBuku()
	case 2:
		LihatBuku()
	case 3:
		HapusBuku()
	case 4:
		EditBuku()
	case 5:
		os.Exit(0)
	}

	main()
}