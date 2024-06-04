/*

	============================================
	Progres pengerjaan tubes praktikum alpro
	Kelompok 11
	DS-47-03
	Anggota:
	-	Fauzan Ahsanudin Alfikri (103052300003)
	-	Risma Febriyanti (103052300111)
	============================================
	Jobdesk
	Fauzan: Pembuatan fungsi dan pMenudur lanjutan
	Risma: Pembuatan kategori menu dsb

*/

// Pengembalian, edit buku,edit peminjam

package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	width    = 30
	duration = 3 * time.Second
) // ! const untuk hiasan

type buku struct {
	judul, genre, penulis, penerbit, Status string
	stok, harga, kategori, rankterpinjam    int
	key                                     [3]string
}

type PinjamanBuku struct {
	NomorPeminjaman   int
	Namapeminjam      string
	Buku              [5]string
	TanggalPeminjaman date
	TanggalKembali    date
	Status            string
	tarifharga        int
	totalBukuDipinjam int
}

type date struct {
	tanggal, bulan, tahun int
}

type account struct {
	Username string
	Password string
}

const NMAX int = 10000

type borrow [NMAX]PinjamanBuku

var datapeminjam borrow
var item [NMAX]buku
var akun [4]account

func dummy(nData *int) {

	// ? Pengisian account admin
	akun[0].Username = "Fauzan"
	akun[0].Password = "Stayhalal"

	akun[1].Username = "Imaaachan"
	akun[1].Password = "baikhati"

	akun[2].Username = "1"
	akun[2].Password = "1"

	akun[3].Username = "reikisenpai"
	akun[3].Password = "Bismillah"

	// ? Pengisian data buku

	item[0].judul = "The Hunger Games "
	item[1].judul = "Laskar Pelangi "
	item[2].judul = "Your Name "
	item[3].judul = "Danur "
	item[4].judul = "Hannibal "

	item[0].genre = "Action "
	item[1].genre = "Fiction "
	item[2].genre = "Fantasy "
	item[3].genre = "Horror "
	item[4].genre = "Psychological "

	item[0].penulis = "Suzanne Collins "
	item[1].penulis = "Andrea Hirata "
	item[2].penulis = "Makoto Shinkai "
	item[3].penulis = "Risa Saraswati "
	item[4].penulis = "Thomas Harris "

	item[0].penerbit = "Gramedia Pustaka Utama "
	item[1].penerbit = "Bentang Pustaka "
	item[2].penerbit = "HARU "
	item[3].penerbit = "Bukune "
	item[4].penerbit = "Delacorte Press "

	item[0].stok = 5
	item[1].stok = 4
	item[2].stok = 6
	item[3].stok = 2
	item[4].stok = 3

	item[0].harga = 60000
	item[1].harga = 70000
	item[2].harga = 40000
	item[3].harga = 100000
	item[4].harga = 90000

	item[0].kategori = 1
	item[1].kategori = 2
	item[2].kategori = 3
	item[3].kategori = 4
	item[4].kategori = 2

	item[0].key[0] = "Adventure"
	item[0].key[1] = "Action"
	item[0].key[2] = "Dystopian"

	item[1].key[0] = "Friendship"
	item[1].key[1] = "Adventure"
	item[1].key[2] = "Drama"

	item[2].key[0] = "Japan"
	item[2].key[1] = "TimeTravelling"
	item[2].key[2] = "Romance"

	item[3].key[0] = "TrueStory"
	item[3].key[1] = "Mystery"
	item[3].key[2] = "Supernatural"

	item[4].key[0] = "Crime"
	item[4].key[1] = "Gore"
	item[4].key[2] = "MindBending"

	item[0].Status = "Ready"
	item[1].Status = "Ready"
	item[2].Status = "Ready"
	item[3].Status = "Ready"
	item[4].Status = "Ready"

	*nData = 5
}

func main() {
	var nData, nPeminjam int

	dummy(&nData)
	menu(&nData, &nPeminjam)
}

// ! Fungsi Utama 1:

func menu(nData *int, nPeminjam *int) { // ? menampilkan menu
	var pilih int
	cls()
	logoaplikasi()
	fmt.Println("1. \U0001F4BB Login")
	fmt.Println("2. \U000023F9  End")
	fmt.Print("Pilihan Anda (1/2): ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		loading()
		Login(&*nData, &*nPeminjam)
	case 2:
		logout()
	default:
		fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan ketik ulang")
		jedawaktu(3)
		cls()
		menu(&*nData, &*nPeminjam)
	}
}

func menuadmin(nData *int, nPeminjam *int) {
	cls()
	var pilih int
	headerMenuAdmin()
	fmt.Println(" 1. \U0001F4DA Data buku")
	fmt.Println(" 2. \U0001F4DD Data peminjaman")
	fmt.Println(" 3. \U000025C0 Log Out")
	fmt.Print("Pilihan Anda (1/2/3): ")
	fmt.Scan(&pilih)
	fmt.Println()
	switch pilih {
	case 1:
		menudatabuku(&*nData, &*nPeminjam)
	case 2:
		menupinjaman(&*nData, &*nPeminjam)
	case 3:
		logout()
	default:
		fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan ketik ulang")
		jedawaktu(3)
		cls()
		menuadmin(&*nData, &*nPeminjam)
	}
}

func Login(nData *int, nPeminjam *int) { // ? Procedure untuk login
	cls()
	var i int
	var us, pws bool
	var user, pw string
	us = false
	pws = false
	headerLogin()
	fmt.Print("Masukkan Username Admin Anda : ")
	fmt.Scan(&user)
	fmt.Print("Masukkan Password Admin Anda : ")
	fmt.Scan(&pw)
	for i < 4 {
		if user == akun[i].Username || us {
			us = true
			if pw == akun[i].Password {
				pws = true
			}
		}
		i++
	}
	if us == false {
		fmt.Println("Username incorect")
		jedawaktu(3)
		cls()
		Login(&*nData, &*nPeminjam)
	} else if pws == false {
		fmt.Println("Password incorect")
		jedawaktu(3)
		cls()
		Login(&*nData, &*nPeminjam)
	} else {
		fmt.Println("Username correct")
		fmt.Println("Password correct")
		loading()
		menuadmin(&*nData, &*nPeminjam)
	}
}

func logout() {
	cls()
	fmt.Print("Logout")
	fmt.Print(".")
	jedawaktu(1)
	fmt.Print(".")
	jedawaktu(1)
	fmt.Print(".")
	jedawaktu(1)
	fmt.Print(".")
	cls()
	fmt.Print("Logout berhasil sampai jumpa lain waktu")
	jedawaktu(3)
	cls()
}

// ! ================================================================

// ! FUNGSI DATA BUKU

func menudatabuku(nData *int, nPeminjam *int) { // ? Procedure print output menu data buku
	cls()
	var pilih int
	headerDatabuku()
	fmt.Println("1. \U0001F4D6 Lihat data buku")
	fmt.Println("2. \U0001f58a Edit data buku")
	fmt.Println("3. \U000025C0 Kembali")
	fmt.Print("Pilihan Anda (1/2/3): ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		cls()
		databuku(&*nData, &*nPeminjam)
	case 2:
		cls()
		menuMengedit(&*nData, &*nPeminjam)
	case 3:
		cls()
		menuadmin(&*nData, &*nPeminjam)
	default:
		fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan ketik ulang")
		jedawaktu(3)
		cls()
		menudatabuku(&*nData, &*nPeminjam)
	}
}

func databuku(nData *int, nPeminjam *int) { // ? Procedure melihat data buku apa aja yang ada
	cls()
	var pilih int
	var wait byte
	if *nData == 0 {
		fmt.Println("\x1b[35mMaaf, data buku Anda masih kosong, silahkan untuk menginput dulu data buku Anda\x1b[0m")
		fmt.Println("1. \U000025B6 Lanjut")
		fmt.Println("2. \U000025C0 Kembali ke menu")
		fmt.Print("\x1b[36mPilihan Anda (1/2): \x1b[0m")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			menuMengedit(&*nData, &*nPeminjam)
		case 2:
			menudatabuku(&*nData, &*nPeminjam)
		default:
			fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan untuk mengetik ulang")
			jedawaktu(3)
			cls()
			databuku(&*nData, &*nPeminjam)
		}
	} else {
		headerDatabuku()
		fmt.Println("1. \U0001f50d Cari data buku")
		fmt.Println("2. \U0001f310 Menampilkan semua data buku")
		fmt.Println("3. \U000025C0 Kembali ke menu")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			menucariData(&*nData, &*nPeminjam)
		case 2:
			cetakdatabuku(&*nData, &*nPeminjam)
			fmt.Print("Klik Enter untuk melanjutkan....")
			fmt.Scanf("\n%c", &wait)
			menuadmin(&*nData, &*nPeminjam)
		case 3:
			menudatabuku(&*nData, &*nPeminjam)
		default:
			fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan ketik ulang")
			jedawaktu(3)
			cls()
		}
	}
}

func menucariData(nData *int, nPeminjam *int) {
	cls()
	var pilih int
	headerCariDataBuku()
	fmt.Println("1. \U0001f50d Cari Judul")
	fmt.Println("2. \U0001f50d\U0001f194 Cari Kata Kunci")
	fmt.Println("3. \U000025C0 Kembali")
	fmt.Print("Pilihan Anda (1/2/3): ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		carijudul(&*nData, &*nPeminjam)
	case 2:
		carikatakunci(&*nData, &*nPeminjam)
	case 3:
		databuku(&*nData, &*nPeminjam)
	default:
		fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan ketik ulang")
		jedawaktu(3)
		cls()
		menucariData(&*nData, &*nPeminjam)
	}
}

func carijudul(nData *int, nPeminjam *int) {
	cls()
	var judul string
	var wait byte
	var pilih int
	var lanjut bool = false
	fmt.Print("Masukkan Judul Yang akan dicari (cth: Danur .): ")
	inputtext(&judul)
	fmt.Println()
	if mencaridata(*nData, judul) != -1 {
		cls()
		cetakbuku(mencaridata(*nData, judul))
		fmt.Print("Klik Enter untuk melanjutkan....")
		fmt.Scanf("\n%c", &wait)
		menuadmin(&*nData, &*nPeminjam)
	} else {
		for !lanjut {
			cls()
			fmt.Println("Maaf, Judul yang anda cari tidak ada")
			fmt.Println("Apakah Anda ingin mencari lagi judul yang lain?")
			fmt.Println("1. \U000025B6 Ya")
			fmt.Println("2. \U00002795\U0001f4dd Tambahkan data buku")
			fmt.Println("3. \U000025C0 Tidak, kembali saja")
			fmt.Println("Jika buku yang anda cari tidak ada")
			fmt.Println("Silahkan untuk menambahkan data buku di menu edit buku / ketik 2")
			fmt.Print("Pilihan anda (1/2/3): ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				lanjut = true
				carijudul(&*nData, &*nPeminjam)
			case 2:
				lanjut = true
				menuMengedit(&*nData, &*nPeminjam)
			case 3:
				lanjut = true
				menucariData(&*nData, &*nPeminjam)
			default:
				fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan ketik ulang")
				jedawaktu(3)
			}
		}
	}
}

func carikatakunci(nData *int, nPeminjam *int) {
	cls()
	var wait byte
	var pilih, batas, i, g, keprint int
	var kunci [5]string
	var tercetak [NMAX]string
	var temp string
	var cek bool = true
	var lanjut bool = false
	keprint = 0
	batas = 0
	fmt.Println("Masukkan kata kunci anda, bila sudah, akhiri dengan . atau Maks 5 kata kunci")
	fmt.Print("Kata Kunci: ")
	for batas < 5 && cek {
		fmt.Scan(&temp)
		if temp != "." {
			kunci[batas] = temp
			batas += 1
		} else {
			cek = false
		}
	}
	if cekkunci(*nData, kunci[0]) || cekkunci(*nData, kunci[1]) || cekkunci(*nData, kunci[2]) || cekkunci(*nData, kunci[3]) || cekkunci(*nData, kunci[4]) {
		cls()
		fmt.Println("||============================||")
		fmt.Println("||     BUKU YANG ANDA CARI    ||")
		fmt.Println("||============================||")
		for i = 0; i < batas; i++ {
			for g = 0; g <= *nData; g++ {
				if mencarikunci(nData, kunci[i], g) == g {
					if sudahtercetak(*nData, tercetak, item[g].judul, keprint) {
						tercetak[keprint] = item[g].judul
						keprint += 1
						cetakbuku(g)
					}
				}
			}
			g = 0
		}
		fmt.Print("Klik Enter untuk melanjutkan....")
		fmt.Scanf("\n%c", &wait)
		menuadmin(&*nData, &*nPeminjam)
	} else {
		for !lanjut {
			cls()
			fmt.Println("Maaf, Key yang anda cari tidak ada")
			fmt.Println("Apakah Anda ingin mencari lagi key yang lain?")
			fmt.Println("1. \U000025B6 Ya")
			fmt.Println("2. \U000025C0 Tidak, kembali saja")
			fmt.Print("Pilihan anda (1/2): ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				lanjut = true
				carikatakunci(&*nData, &*nPeminjam)
			case 2:
				lanjut = true
				menucariData(&*nData, &*nPeminjam)
			default:
				fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan ketik ulang")
				jedawaktu(3)
			}
		}
	}
}

func sudahtercetak(nData int, tercetak [NMAX]string, judul string, batas int) bool {
	var i int = 0
	var cek bool
	if batas == 0 {
		cek = true
	} else {
		for i < batas {
			if judul == tercetak[i] {
				cek = false
				i += 1000
			} else {
				cek = true
			}
			i++
		}
	}
	return cek
}

func cekkunci(nData int, kunci string) bool {
	var cek bool = false
	var i int
	for i = 0; i < nData && !cek; i++ {
		if kunci == item[i].key[0] || kunci == item[i].key[1] || kunci == item[i].key[2] {
			cek = true
		}
	}
	return cek
}

func mencarikunci(nData *int, kunci string, i int) int {
	if kunci == item[i].key[0] || kunci == item[i].key[1] || kunci == item[i].key[2] {
		return i
	} else {
		return -1
	}
}

func cetakdatabuku(nData *int, nPeminjam *int) { // ? Procedure Mencetak data buku
	cls()
	var ascending bool = true
	var pilih int
	var lanjut bool = false
	for !lanjut {
		cls()
		fmt.Println("Pilih metode untuk tampilan data buku")
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Print("Pilihan Anda (1/2): ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			lanjut = true
		case 2:
			lanjut = true
			ascending = false
		default:
			fmt.Println("Pilihan Anda invalid!")
			jedawaktu(3)
		}
	}
	inserctionsortJudul(*nData, ascending)
	selectionsortKategori(*nData)
	cls()
	fmt.Println("||============================||")
	fmt.Println("||         DATA BUKU          ||")
	fmt.Println("||============================||")
	fmt.Println()
	for i := 0; i < *nData; i++ {
		fmt.Printf("\x1b[35mNO: %d\n", i+1)
		cetakbuku(i)
	}
}

func cetakbuku(data int) {
	fmt.Printf("\x1b[32mJUDUL: %s\n", item[data].judul)
	fmt.Printf("\x1b[32mPENULIS: %s\n", item[data].penulis)
	fmt.Printf("\x1b[32mPENERBIT: %s\n", item[data].penerbit)
	fmt.Printf("\x1b[32mGENRE: %s\n", item[data].genre)
	fmt.Printf("\x1b[32mHARGA: %d\n", item[data].harga)
	fmt.Printf("\x1b[32mKATEGORI: %s\n", convertkategori(item[data].kategori))
	fmt.Printf("\x1b[32mSTOK: %d\n\x1b[0m\n", item[data].stok)
	fmt.Print("\x1b[32mKey :")
	for i := 0; i < 3; i++ {
		fmt.Print(" ", item[data].key[i], " ")
	}
	fmt.Print("\n\x1b[0m\n")
	fmt.Println("______________________________")
}

func convertkategori(i int) string { // ? Function mengembalikan var kategori integer menjadi string komik/novel/cerpen/pelajaran
	if i == 1 {
		return "\x1b[31m \U0001F4C3 Cerpen\x1b[0m"
	} else if i == 2 {
		return "\x1b[32m \U0001F4AC Komik\x1b[0m"
	} else if i == 3 {
		return "\033[33m \U0001F3AD Novel\x1b[0m"
	} else if i == 4 {
		return "\x1b[34m \U0001F3EB Pelajaran\x1b[0m"
	} else {
		return "Kategori yang anda masukkan tidak valid"
	}
}

func menuMengedit(nData *int, nPeminjam *int) { // ? Procedure menu edit
	cls()
	var pilih int
	headerEditDataBuku()
	fmt.Println("1. \U0001F539 Tambah data buku")
	fmt.Println("2. \U0001F53B Hapus data buku")
	fmt.Println("3. \U0001F53B Edit Buku")
	fmt.Println("4. \U000025C0 Kembali")
	fmt.Print("Pilihan Anda (1/2/3/4) : ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		menambahkan(&*nData, &*nPeminjam)
	case 2:
		menghapus(&*nData, &*nPeminjam)
	case 3:
		MeditBuku(&*nData, &*nPeminjam)
	case 4:
		menudatabuku(&*nData, &*nPeminjam)
	default:
		fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan ketik ulang")
		jedawaktu(3)
		cls()
		menuMengedit(&*nData, &*nPeminjam)
	}
}

func MeditBuku(nData, nPeminjam *int) {
	var judul string
	var pilih, data int
	var lanjut bool = false
	cetakdatabuku(&*nData, &*nPeminjam)
	fmt.Print("Masukkan judul buku yang ingin di ubah (cth: 'Danur .'):")
	inputtext(&judul)
	if mencaridata(*nData, judul) != -1 {
		data = mencaridata(*nData, judul)
		editBuku(&*nData, &*nPeminjam, data)
	} else {
		for !lanjut {
			cls()
			fmt.Println("Maaf judul buku yang anda masukkan tidak ada")
			fmt.Println("1. \U000025B6 Input Kembali")
			fmt.Println("2. \U000025C0 Kembali")
			fmt.Print("\x1b[36mPilihan Anda (1/2): \x1b[0m")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				lanjut = true
				MeditBuku(&*nData, &*nPeminjam)
			case 2:
				lanjut = true
				menuMengedit(&*nData, &*nPeminjam)
			default:
				fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan ketik ulang")
				jedawaktu(3)
			}
		}
	}
}

func editBuku(nData, nPeminjam *int, data int) {
	cls()
	var pilih int
	fmt.Println("Pilih apa yang akan di ubah:")
	fmt.Println("1. Judul Buku")
	fmt.Println("2. Penulis")
	fmt.Println("3. Penerbit")
	fmt.Println("4. Genre")
	fmt.Println("5. Kategori")
	fmt.Println("6. Harga")
	fmt.Println("7. Stok")
	fmt.Println("8. Key")
	fmt.Println("9. Semua")
	fmt.Println("10. Kembali")
	fmt.Print("Pilihan Anda (1/2/3/4/5/6/7/8/9/10): ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		editjudul(&*nData, &*nPeminjam, data)
		editBuku(&*nData, &*nPeminjam, data)
	case 2:
		editpenulis(&*nData, &*nPeminjam, data)
		editBuku(&*nData, &*nPeminjam, data)
	case 3:
		editpenerbit(&*nData, &*nPeminjam, data)
		editBuku(&*nData, &*nPeminjam, data)
	case 4:
		editgenre(&*nData, &*nPeminjam, data)
		editBuku(&*nData, &*nPeminjam, data)
	case 5:
		editkategori(&*nData, &*nPeminjam, data)
		editBuku(&*nData, &*nPeminjam, data)
	case 6:
		editharga(&*nData, &*nPeminjam, data)
		editBuku(&*nData, &*nPeminjam, data)
	case 7:
		editstok(&*nData, &*nPeminjam, data)
		editBuku(&*nData, &*nPeminjam, data)
	case 8:
		editkey(&*nData, &*nPeminjam, data)
		editBuku(&*nData, &*nPeminjam, data)
	case 9:
		editjudul(&*nData, &*nPeminjam, data)
		editpenulis(&*nData, &*nPeminjam, data)
		editpenerbit(&*nData, &*nPeminjam, data)
		editgenre(&*nData, &*nPeminjam, data)
		editkategori(&*nData, &*nPeminjam, data)
		editharga(&*nData, &*nPeminjam, data)
		editstok(&*nData, &*nPeminjam, data)
		editBuku(&*nData, &*nPeminjam, data)
	case 10:
		menuMengedit(&*nData, &*nPeminjam)
	default:
		fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan ketik ulang")
		jedawaktu(3)
		cls()
		editBuku(&*nData, &*nPeminjam, data)
	}
}

func editjudul(nData, nPeminjam *int, data int) {
	cls()
	var ganti string
	fmt.Print("Masukkan judul yang baru (cth: Villains ./ heroes academia .): ")
	inputtext(&ganti)
	item[data].judul = ganti
	cls()
	fmt.Println("Buku baru:")
	cetakbuku(data)
	jedawaktu(3)
}

func editpenulis(nData, nPeminjam *int, data int) {
	cls()
	var ganti string
	fmt.Print("Masukkan nama penulis yang baru (cth: Nakama ./ Watashi wa .): ")
	inputtext(&ganti)
	item[data].penulis = ganti
	cls()
	fmt.Println("Buku baru:")
	cetakbuku(data)
	jedawaktu(3)
}

func editpenerbit(nData, nPeminjam *int, data int) {
	cls()
	var ganti string
	fmt.Print("Masukkan nama penerbit yang baru (cth: Nvidia ./ Nakama nei .)")
	inputtext(&ganti)
	item[data].penerbit = ganti
	cls()
	fmt.Println("Buku baru:")
	cetakbuku(data)
	jedawaktu(3)
}

func editgenre(nData, nPeminjam *int, data int) {
	cls()
	var ganti string
	fmt.Print("Masukkan nama genre yang baru (cth: Horor ./ Slice Of Life .): ")
	inputtext(&ganti)
	item[data].genre = ganti
	cls()
	fmt.Println("Buku baru:")
	cetakbuku(data)
	jedawaktu(3)
}

func editkategori(nData, nPeminjam *int, data int) {
	cls()
	var kategori string
	fmt.Println("\x1b[36mKategori\x1b[0m")
	fmt.Println("\x1b[31m \U0001F4C3 Cerpen\x1b[0m")
	fmt.Println("\x1b[32m \U0001F4AC Komik\x1b[0m")
	fmt.Println("\x1b[33m \U0001F3AD Novel\x1b[0m")
	fmt.Println("\x1b[34m \U0001F3EB Pelajaran \x1b[0m")
	fmt.Print("Pilihan Anda (Komik/Novel/Cerpen/Pelajaran): ")
	fmt.Scan(&kategori)
	if kategori == "Cerpen" || kategori == "Komik" || kategori == "Novel" || kategori == "Pelajaran" {
		switch kategori {
		case "Cerpen":
			item[data].kategori = 1
			fmt.Println("\x1b[32mBerhasil disimpan, Dalam kategori \U0001F4C3 \x1b[31mCerpen\x1b[0m")
		case "Komik":
			item[data].kategori = 2
			fmt.Println("\x1b[32mBerhasil disimpan, Dalam kategori \U0001F4AC \x1b[32mKomik\x1b[0m")
		case "Novel":
			item[data].kategori = 3
			fmt.Println("\x1b[32mBerhasil disimpan, Dalam kategori \U0001F3AD \x1b[33mNovel\x1b[0m")
		case "Pelajaran":
			item[data].kategori = 4
			fmt.Println("\x1b[32mBerhasil disimpan, Dalam kategori \U0001F3EB \x1b[34mPelajaran\x1b[0m")
		}
		cls()
		fmt.Println("Buku baru:")
		cetakbuku(data)
		jedawaktu(3)
	} else {
		cls()
		fmt.Println("Gagal, tidak ada pilihan ", kategori, " didalam pilihan")
		jedawaktu(3)
		fmt.Println("Silahkan untuk menginput ulang")
		jedawaktu(3)
		editkategori(&*nData, &*nPeminjam, data)
	}
}

func editharga(nData, nPeminjam *int, data int) {
	cls()
	var ganti int
	fmt.Print("Masukkan Harga baru (cth: 2000)")
	fmt.Scan(&ganti)
	if ganti > 0 {
		item[data].harga = ganti
		cls()
		fmt.Println("Buku baru:")
		cetakbuku(data)
		jedawaktu(3)
	} else {
		fmt.Println("Invalid Number")
		jedawaktu(3)
		editharga(&*nData, &*nPeminjam, data)
	}
}

func editstok(nData, nPeminjam *int, data int) {
	cls()
	var ganti int
	fmt.Print("Masukkan stok yang baru (cth: 1)")
	fmt.Scan(&ganti)
	if ganti > 0 {
		item[data].stok = ganti
		cls()
		fmt.Println("Buku baru:")
		cetakbuku(data)
		jedawaktu(3)
	} else {
		fmt.Println("Invalid Number")
		jedawaktu(3)
		editstok(&*nData, &*nPeminjam, data)
	}
}

func editkey(nData, nPeminjam *int, data int) {
	var pilih int
	var baru string
	var lanjut bool = false
	for i := 0; i < 3; i++ {
		for !lanjut {
			cls()
			fmt.Println(i, ". ", item[data].key[i])
			fmt.Println("Apakah anda mau mengubahnya:")
			fmt.Println("1. Iya")
			fmt.Println("2. Tidak")
			fmt.Print("Pilihan anda(1/2):")
			fmt.Scan(&pilih)
			if pilih == 1 {
				cls()
				lanjut = true
				fmt.Print("Masukkan Key anda yang baru (cth: walawe): ")
				fmt.Scan(&baru)
				item[data].key[i] = baru
			} else if pilih == 2 {
				lanjut = true
			} else {
				fmt.Println("Invalid, silahkan untuk menginputkan ulang pilihan anda")
				jedawaktu(3)
			}
		}
		lanjut = false
	}
	fmt.Println("Buku baru:")
	cetakbuku(data)
	jedawaktu(3)
}

func menambahkan(nData *int, nPeminjam *int) { // ? Procedure menambahkan buku kedalam data
	cls()
	var banyakstok, produk int
	fmt.Print("Masukkan jumlah buku yang akan ditambahkan (1/2/3/dsb): ")
	fmt.Scan(&banyakstok)
	if banyakstok > 0 {
		for produk = 0; produk < banyakstok; produk++ {
			menambahkan2(&*nData, &*nPeminjam, &produk, &banyakstok)
		}
		cls()
		menuadmin(&*nData, &*nPeminjam)
	} else {
		fmt.Println("Invalid Number!")
		fmt.Println("Masukkan jumlah baru: ")
		jedawaktu(3)
		menambahkan(&*nData, &*nPeminjam)
	}
}

func menambahkan2(nData *int, nPeminjam *int, produk, banyakstok *int) { // ? Subs procedure dari procedure menambahkan buku
	var lanjut bool = false
	var kategori string
	var judul string
	fmt.Print("\n\x1b[36mJUDUL (Akhiri dengan spasi .): \x1b[0m")
	inputtext(&judul)
	if mencaridata(*nData, judul) != -1 {
		cls()
		pilihanmenambah(&*nData, &*nPeminjam, &*produk, &*banyakstok, &judul)
	} else {
		item[*nData].judul = judul
		fmt.Print("\x1b[36mPENULIS (cth: Nvidia .): \x1b[0m")
		inputtext(&item[*nData].penulis)
		fmt.Print("\x1b[36mPENERBIT (cth : Imaaachan .): \x1b[0m")
		inputtext(&item[*nData].penerbit)
		fmt.Print("\x1b[36mGENRE (cth: Horor .): \x1b[0m")
		inputtext(&item[*nData].genre)
		fmt.Print("\x1b[36mHARGA (): \x1b[0m")
		fmt.Scan(&item[*nData].harga)
		for !lanjut {
			fmt.Println("\x1b[36mKategori\x1b[0m")
			fmt.Println("\x1b[32m \U0001F4C3 Cerpen\x1b[0m")
			fmt.Println("\x1b[31m \U0001F4AC Komik\x1b[0m")
			fmt.Println("\x1b[33m \U0001F3AD Novel\x1b[0m")
			fmt.Println("\x1b[34m \U0001F3EB Pelajaran \x1b[0m")
			fmt.Print("Pilihan Anda (Cerpen/Komik/Novel/Pelajaran) : ")
			fmt.Scan(&kategori)
			switch kategori {
			case "Cerpen":
				lanjut = true
				item[*nData].kategori = 1
				fmt.Println("\x1b[32mBerhasil disimpan, Dalam kategori \U0001F4C3 \x1b[32mCerpen\x1b[0m")
			case "Komik":
				lanjut = true
				item[*nData].kategori = 2
				fmt.Println("\x1b[32mBerhasil disimpan, Dalam kategori \U0001F4AC \x1b[31mKomik\x1b[0m")
			case "Novel":
				lanjut = true
				item[*nData].kategori = 3
				fmt.Println("\x1b[32mBerhasil disimpan, Dalam kategori \U0001F3AD \x1b[33mNovel\x1b[0m")
			case "Pelajaran":
				lanjut = true
				item[*nData].kategori = 4
				fmt.Println("\x1b[32mBerhasil disimpan, Dalam kategori \U0001F3EB \x1b[34mPelajaran\x1b[0m")
			default:
				fmt.Println("Gagal, tidak ada pilihan ", kategori, " didalam pilihan")
				jedawaktu(3)
				fmt.Println("Silahkan untuk menginput ulang")
				jedawaktu(3)
			}
		}
		fmt.Print("\x1b[36mStok buku ada berapa? \x1b[0m")
		fmt.Scan(&item[*nData].stok)
		fmt.Print("Masukkan 3 Key buku: ")
		for i := 0; i < 3; i++ {
			fmt.Scan(&item[*nData].key[i])
		}
		cls()
		fmt.Println("Berhasil menambahkan:")
		cetakbuku(*nData)
		fmt.Println("_______________________________________")
		jedawaktu(3)
		cls()
		*nData = *nData + 1
	}
}

func inputtext(text *string) {
	*text = ""
	var temp string
	for temp != "." {
		fmt.Scan(&temp)
		if temp != "." {
			*text += temp + " "
		}
	}
}

func pilihanmenambah(nData *int, nPeminjam *int, produk, banyakstok *int, judul *string) { // ? Procedure pilihan jikalau buku yang ditambahkan judul nya sudah ada
	var pilih int
	fmt.Println("||======================================||")
	fmt.Println("||   JUDUL YANG ANDA PILIH SUDAH ADA    ||")
	fmt.Println("||======================================||")
	fmt.Println("Apakah Anda hanya ingin menambahkan stoknya?")
	fmt.Println("1. \U000025B6 Ya")
	fmt.Println("2. \U000025C0 Tidak, kembali ke menu")
	fmt.Print("Pilihan Anda (1/2) : ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		item[mencaridata(*nData, *judul)].stok += 1
		*produk += 1
	case 2:
		*produk = *banyakstok
	default:
		fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan ketik ulang")
		jedawaktu(3)
		cls()
		pilihanmenambah(&*nData, &*nPeminjam, &*produk, &*banyakstok, &*judul)
	}
}

func mencaridata(nData int, search string) int { // ? Function untuk mengembalikan data ke berapa berdasarkan judul buku
	var i int = 0
	for i < nData {
		if search == item[i].judul {
			return i
		}
		i++
	}
	return -1
}

func menghapus(nData *int, nPeminjam *int) { // ? Procedure menghapus data buku ke n
	var judul string
	var pilih, simpandata int
	if *nData == 0 {
		fmt.Println("\x1b[35mMohon maaf, data buku Anda masih kosong, silahkan untuk menginput dulu data buku Anda\x1b[0m")
		fmt.Println("1. \U000025B6 Lanjut")
		fmt.Println("2. \U000025C0 Kembali ke menu")
		fmt.Print("\x1b[36mPilihan Anda: \x1b[0m")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			menambahkan(&*nData, &*nPeminjam)
		case 2:
			menuadmin(&*nData, &*nPeminjam)
		default:
			fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan untuk mengetik ulang")
			jedawaktu(3)
			cls()
			menghapus(&*nData, &*nPeminjam)
		}
	} else {
		cetakdatabuku(&*nData, &*nPeminjam)
		fmt.Print("Masukkan Judul Buku yang akan dihapus (Akhiri dengan spasi .): ")
		inputtext(&judul)
		if mencaridata(*nData, judul) != -1 {
			cls()
			simpandata = mencaridata(*nData, judul)
			fmt.Println("Berhasil menghapus buku: ")
			cetakbuku(simpandata)
			jedawaktu(4)
			for simpandata < *nData-1 {
				item[simpandata].judul = item[simpandata+1].judul
				item[simpandata].kategori = item[simpandata+1].kategori
				item[simpandata].genre = item[simpandata+1].genre
				item[simpandata].penerbit = item[simpandata+1].penerbit
				item[simpandata].penulis = item[simpandata+1].penulis
				item[simpandata].stok = item[simpandata+1].stok
				item[simpandata].harga = item[simpandata+1].harga
				simpandata += 1
			}
			*nData -= 1
			menuMengedit(&*nData, &*nPeminjam)
		} else {
			fmt.Println("Maaf judul buku yang anda masukkan tidak ada")
			fmt.Println("1. \U000025B6 Input Kembali")
			fmt.Println("2. \U000025C0 Kembali")
			fmt.Print("\x1b[36mPilihan Anda (1/2): \x1b[0m")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				menghapus(&*nData, &*nPeminjam)
			case 2:
				menuMengedit(&*nData, &*nPeminjam)
			default:
				fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan untuk mengetik ulang")
				jedawaktu(3)
				cls()
				menghapus(&*nData, &*nPeminjam)
			}
		}
	}
}

func inserctionsortJudul(nData int, ascending bool) {
	var pass, i int
	var temp buku
	pass = 1
	if ascending {
		for pass <= nData-1 {
			i = pass
			temp = item[pass]
			for i > 0 && temp.judul < item[i-1].judul {
				item[i] = item[i-1]
				i -= 1
			}
			item[i] = temp
			pass += 1
		}
	} else {
		for pass <= nData-1 {
			i = pass
			temp = item[pass]
			for i > 0 && temp.judul > item[i-1].judul {
				item[i] = item[i-1]
				i -= 1
			}
			item[i] = temp
			pass += 1
		}
	}
}

func selectionsortKategori(nData int) {
	var pass, idx, i int
	var temp buku
	pass = 1
	for pass <= nData-1 {
		idx = pass - 1
		i = pass
		for i < nData {
			if item[idx].kategori > item[i].kategori {
				idx = i
			}
			i += 1
		}
		temp = item[pass-1]
		item[pass-1] = item[idx]
		item[idx] = temp
		pass += 1
	}
}

// ! =======================================================

// ! FUNGSI DATA PEMINJAMAN

func menupinjaman(nData *int, nPeminjam *int) {
	cls()
	var pilih int
	fmt.Println("||===================||")
	fmt.Println("||   Menu Pinjaman   ||")
	fmt.Println("||===================||")
	fmt.Println("1. \U0001f4da\U0001f504 Peminjaman Buku")
	fmt.Println("2. \U0001f4da Hapus Data Peminjaman Buku")
	fmt.Println("3. \U0001f4cb Daftar Buku Dipinjam")
	fmt.Println("4. \U0001f4cb Rekomendasi 5 Buku Terfavorit")
	fmt.Println("5. \U0001f50d Pencarian Data pinjaman")
	fmt.Println("6. \U00002795 Perpanjangan Peminjaman")
	fmt.Println("7. \U0001f504 Pengembalian Buku")
	fmt.Println("8. \U0001f504 Edit data peminjam")
	fmt.Println("9. \U000025C0 Kembali")
	fmt.Print("Pilihan Anda (1/2/3/4/5/6/7/8/9): ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		pinjaman(&*nData, &*nPeminjam)
	case 2:
		hapusdatapeminjam(&*nData, &*nPeminjam)
	case 3:
		datapinjaman(&*nData, &*nPeminjam)
	case 4:
		Terfavorit(&*nData, &*nPeminjam)
	case 5:
		menucaridatapinjaman(&*nData, &*nPeminjam)
	case 6:
		perpanjangan(&*nData, &*nPeminjam)
	case 7:
		menupengembalian(&*nData, &*nPeminjam)
	case 8:
		MeditPinjaman(&*nData, &*nPeminjam)
	case 9:
		menuadmin(&*nData, &*nPeminjam)
	default:
		fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan ketik ulang")
		jedawaktu(3)
		cls()
		menupinjaman(&*nData, &*nPeminjam)
	}
}

func menupengembalian(nData *int, nPeminjam *int) {
	cls()
	var pilih int
	fmt.Println("||===================||")
	fmt.Println("|| Menu pengembalian ||")
	fmt.Println("||===================||")
	fmt.Println("1. \U0001f504 Pengembalian Buku")
	fmt.Println("2. \U0001f6ab Hilang")
	fmt.Println("3. \U000025C0 Kembali")
	fmt.Print("Pilihan Anda (1/2/3): ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		pengembalian(&*nData, &*nPeminjam)
	case 2:
		hilang(&*nData, &*nPeminjam)
	case 3:
		menupinjaman(&*nData, &*nPeminjam)
	default:
		fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan ketik ulang")
		jedawaktu(3)
		cls()
		menupengembalian(&*nData, &*nPeminjam)
	}
}

func hapusdatapeminjam(nData, nPeminjam *int) {
	cls()
	var lanjut bool = false
	var nama string
	var pilih, data int
	if *nPeminjam == 0 {
		fmt.Println("Maaf data peminjam anda masih kosong")
		jedawaktu(5)
		menupinjaman(&*nData, &*nPeminjam)
	} else {
		fmt.Print("Masukkan nama peminjam yang akan di hapus (cth: 'Fauzan .'):")
		inputtext(&nama)
		if caridatapeminjam(*nPeminjam, nama) != -1 {
			data = caridatapeminjam(*nPeminjam, nama)
			cetakdatapeminjam(*nPeminjam, data)
			for g := 0; g < datapeminjam[data].totalBukuDipinjam; g++ {
				item[mencaridata(*nData, datapeminjam[data].Buku[g])].stok += 1
				item[mencaridata(*nData, datapeminjam[data].Buku[g])].rankterpinjam -= 1
			}
			for i := data; i < *nPeminjam; i++ {
				datapeminjam[i] = datapeminjam[i+1]
			}
			*nPeminjam -= 1
			jedawaktu(3)
			fmt.Println("Berhasil Dihapus")
			jedawaktu(3)
			menupinjaman(&*nData, &*nPeminjam)
		} else {
			for !lanjut {
				fmt.Println("Maaf Nama peminjam tidak ada")
				fmt.Println("Apakah anda ingin menginputkannya kembali?")
				fmt.Println("1. \U000025B6 Input kembali")
				fmt.Println("2. \U000025C0 Kembali")
				fmt.Print("Pilihan anda (1/2): ")
				fmt.Scan(&pilih)
				switch pilih {
				case 1:
					lanjut = true
					hapusdatapeminjam(&*nData, &*nPeminjam)
				case 2:
					lanjut = true
					menupinjaman(&*nData, &*nPeminjam)
				default:
					fmt.Println("pilihan anda invalid")
					jedawaktu(3)
					cls()
				}
			}
		}
	}
}

func pengembalian(nData *int, nPeminjam *int) {
	cls()
	var pilih int
	var nama string
	var bulan, tahun, tanggal int
	var denda int
	var i int
	var lanjut bool = false
	fmt.Print("Nama peminjam (cth: 'Fauzan .'): ")
	inputtext(&nama)
	if caridatapeminjam(*nPeminjam, nama) != -1 {
		fmt.Print("\nTanggal Kembali: ")
		fmt.Scan(&tanggal)
		fmt.Print("\nBulan Kembali: ")
		fmt.Scan(&bulan)
		fmt.Print("\nTahun Kembali: ")
		fmt.Scan(&tahun)
		if tanggal > 31 {
			tanggal -= 31
			bulan += 1
		}
		if bulan > 12 {
			bulan -= 12
			tahun += 1
		}
		for i <= datapeminjam[caridatapeminjam(*nPeminjam, nama)].totalBukuDipinjam {
			fmt.Println("Apakah buku ", datapeminjam[caridatapeminjam(*nPeminjam, nama)].Buku[i], " akan dikembalikan?")
			fmt.Println("1. \U00002705 Iya")
			fmt.Println("2. \U0000274c Tidak")
			fmt.Print("Pilihan Anda (1/2): ")
			fmt.Scan(&pilih)
			if pilih == 1 {
				item[mencaridata(*nData, datapeminjam[caridatapeminjam(*nPeminjam, nama)].Buku[i])].stok += 1
				if i < 4 {
					datapeminjam[caridatapeminjam(*nPeminjam, nama)].totalBukuDipinjam -= 1
					datapeminjam[caridatapeminjam(*nPeminjam, nama)].Buku[i] = datapeminjam[caridatapeminjam(*nPeminjam, nama)].Buku[i+1]
					denda += cekdenda(*nPeminjam, *nData, tanggal, bulan, tahun, i, nama)
				} else {
					datapeminjam[caridatapeminjam(*nPeminjam, nama)].Buku[i] = ""
				}
				i++
			} else if pilih == 2 {
				i++
			} else {
				fmt.Println("Pilihan Anda Invalid")
				jedawaktu(3)
			}
		}
		if denda == 0 {

		} else {
			strukpembayaran(&denda)
		}
		if datapeminjam[caridatapeminjam(*nPeminjam, nama)].totalBukuDipinjam == 0 {
			datapeminjam[caridatapeminjam(*nPeminjam, nama)].Status = "Done"
		}
		menupengembalian(&*nData, &*nPeminjam)
	} else {
		for !lanjut {
			cls()
			fmt.Println("Maaf Nama peminjam tidak ada")
			fmt.Println("Silahkan Untuk menginputkan dengan benar")
			fmt.Println("1. \U000025B6 Input kembali")
			fmt.Println("2. \U000025C0 Kembali")
			fmt.Print("Pilihan anda (1/2): ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				lanjut = true
				pengembalian(&*nData, &*nPeminjam)
			case 2:
				lanjut = true
				menupengembalian(&*nData, &*nPeminjam)
			default:
				fmt.Println("\nMaaf pilihan anda invalid, silahkan input kembali")
				jedawaktu(3)
			}
		}
	}
}

func cekdenda(nPeminjam, nData, tanggal, bulan, tahun, i int, nama string) int {
	var denda int
	if tanggal < datapeminjam[caridatapeminjam(nPeminjam, nama)].TanggalKembali.tanggal && tanggal > datapeminjam[caridatapeminjam(nPeminjam, nama)].TanggalPeminjaman.tanggal {
		if bulan <= datapeminjam[caridatapeminjam(nPeminjam, nama)].TanggalKembali.bulan {
			if tahun <= datapeminjam[caridatapeminjam(nPeminjam, nama)].TanggalKembali.tahun {
				return 0
			} else {
				denda = (item[mencaridata(nData, datapeminjam[caridatapeminjam(nPeminjam, nama)].Buku[i])].harga * 20) / 100
			}
		} else {
			denda = (item[mencaridata(nData, datapeminjam[caridatapeminjam(nPeminjam, nama)].Buku[i])].harga * 20) / 100
		}
	} else {
		denda = (item[mencaridata(nData, datapeminjam[caridatapeminjam(nPeminjam, nama)].Buku[i])].harga * 20) / 100
	}
	return denda
}

func hilang(nData *int, nPeminjam *int) {
	cls()
	var lanjut bool = false
	var nama string
	var denda, pilih int
	var i int = 0
	fmt.Print("Nama peminjam: ")
	inputtext(&nama)
	if caridatapeminjam(*nPeminjam, nama) != -1 {
		for i < datapeminjam[caridatapeminjam(*nPeminjam, nama)].totalBukuDipinjam {
			fmt.Println("Apakah buku ", datapeminjam[caridatapeminjam(*nPeminjam, nama)].Buku[i], " hilang?")
			fmt.Println("1. \U00002705 Iya")
			fmt.Println("2. \U0000274c Tidak")
			fmt.Print("Pilihan Anda (1/2): ")
			fmt.Scan(&pilih)
			if pilih == 1 {
				denda += item[mencaridata(*nData, datapeminjam[caridatapeminjam(*nPeminjam, nama)].Buku[i])].harga
				if i < 4 {
					datapeminjam[caridatapeminjam(*nPeminjam, nama)].totalBukuDipinjam -= 1
					datapeminjam[caridatapeminjam(*nPeminjam, nama)].Buku[i] = datapeminjam[caridatapeminjam(*nPeminjam, nama)].Buku[i+1]
				} else {
					datapeminjam[caridatapeminjam(*nPeminjam, nama)].Buku[i] = ""
				}
				i++
			} else if pilih == 2 {
				i++
			} else {
				fmt.Println("Pilihan Anda Invalid")
				jedawaktu(3)
			}
		}
		strukpembayaran(&denda)
		menupengembalian(&*nData, &*nPeminjam)
	} else {
		for !lanjut {
			fmt.Println("Maaf Nama peminjam tidak ada")
			fmt.Println("Silahkan Untuk menginputkan dengan benar")
			fmt.Println("1. \U000025B6 Input kembali")
			fmt.Println("2. \U000025C0 Kembali")
			fmt.Print("Pilihan Anda (1/2): ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				lanjut = true
				hilang(&*nData, &*nPeminjam)
			case 2:
				lanjut = true
				menupengembalian(&*nData, &*nPeminjam)
			default:
				fmt.Println("\nMaaf pilihan anda invalid, silahkan untuk menginputkan kembali")
				jedawaktu(3)
				cls()
			}
		}
	}
}

func strukpembayaran(tagihan *int) {
	cls()
	var bayar int
	fmt.Print("Silahkan untuk membayar sebanyak ", *tagihan, " (cth : 2000): ")
	fmt.Scan(&bayar)
	if bayar == *tagihan {
		fmt.Println("Lunas")
		jedawaktu(3)
	} else if bayar > *tagihan {
		fmt.Println("Ini kembalian anda: ", bayar-*tagihan)
		jedawaktu(3)
	} else {
		fmt.Println("Maaf uang anda kurang: ", *tagihan-bayar)
		jedawaktu(3)
		bayar = *tagihan - bayar
		strukpembayaran(&bayar)
	}
	*tagihan = 0
}

func perpanjangan(nData *int, nPeminjam *int) {
	cls()
	var lanjut bool = false
	var nama string
	var hari, pilih int
	fmt.Print("Masukkan Nama saat anda meminjam buku (cth : Imaachan .): ")
	inputtext(&nama)
	if caridatapeminjam(*nPeminjam, nama) != -1 {
		fmt.Print("Masukkan jumlah perpanjangan hari (1/2/3/dsb): ")
		fmt.Scan(&hari)
		datapeminjam[caridatapeminjam(*nPeminjam, nama)].tarifharga += 5000
		if (datapeminjam[caridatapeminjam(*nPeminjam, nama)].TanggalKembali.tanggal + hari) > 31 {
			datapeminjam[caridatapeminjam(*nPeminjam, nama)].TanggalKembali.tanggal -= 31
			datapeminjam[caridatapeminjam(*nPeminjam, nama)].TanggalKembali.tanggal += 1
		} else {
			datapeminjam[caridatapeminjam(*nPeminjam, nama)].TanggalKembali.tanggal += hari
		}
		strukpembayaran(&datapeminjam[caridatapeminjam(*nPeminjam, nama)].tarifharga)
		menupinjaman(&*nData, &*nPeminjam)
	} else {
		for !lanjut {
			fmt.Println("Maaf data peminjam tidak ada")
			fmt.Println("Apakah Anda ingin menginputkan kembali?")
			fmt.Println("1. Iya")
			fmt.Println("2. Tidak")
			fmt.Print("Pilihan Anda (1/2):")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				lanjut = true
				perpanjangan(&*nData, &*nPeminjam)
			case 2:
				lanjut = true
				menupinjaman(&*nData, &*nPeminjam)
			default:
				fmt.Println("\nMaaf pilihan anda invalid, silahkan input kembali")
				jedawaktu(3)
				cls()
			}
		}
	}
}

func pinjaman(nData *int, nPeminjam *int) {
	cls()
	var lanjut bool = false
	var judul, nama string
	var data, i, pilihan int
	fmt.Print("\nNama Peminjam (cth: 'Fauzan .'): ")
	inputtext(&nama)
	if caridatapeminjam(*nPeminjam, nama) != -1 {
		fmt.Println("Maaf, nama sudah dimiliki orang lain silahkan untuk menginput nama lain")
		jedawaktu(5)
		pinjaman(&*nData, &*nPeminjam)
	} else {
		for !lanjut {
			fmt.Print("\nMau pinjam berapa buku (MAKS 5 Buku): ")
			fmt.Scan(&data)
			if data <= 5 {
				datapeminjam[*nPeminjam].Namapeminjam = nama
				datapeminjam[*nPeminjam].NomorPeminjaman = *nPeminjam + 1
				for i < data {
					cetakdatabuku(&*nData, &*nPeminjam)
					inputjudul(&*nData, &*nPeminjam, &judul, &i, data)
					datapeminjam[*nPeminjam].totalBukuDipinjam = datapeminjam[*nPeminjam].totalBukuDipinjam + 1
					i++
				}
				if i > data {
					menupinjaman(&*nData, &*nPeminjam)
				} else {
					inputdate(&*nPeminjam)
					datapeminjam[*nPeminjam].Status = "Dipinjam"
					pengisiandatatanggalkembali(*nPeminjam)
					strukPeminjaman(datapeminjam[*nPeminjam])
					strukpembayaran(&datapeminjam[*nPeminjam].tarifharga)
					*nPeminjam += 1
					menuadmin(&*nData, &*nPeminjam)
				}
				lanjut = true
			} else {
				for !lanjut {
					cls()
					fmt.Println("Maaf anda meminjam lebih dari batas")
					fmt.Println("Apakah anda ingin menginput kembali?")
					fmt.Println("1. \U00002705 Iya")
					fmt.Println("2. \U000025C0 Tidak, Kembali saja")
					fmt.Print("Pilihan Anda (1/2): ")
					switch pilihan {
					case 1:
						lanjut = true
					case 2:
						lanjut = true
						menupinjaman(&*nData, &*nPeminjam)
					default:
						fmt.Println("\nMaaf pilihan anda salah, silahkan untuk menginput kembali")
						jedawaktu(3)
					}
				}
				lanjut = false
			}
		}
	}
}

func pengisiandatatanggalkembali(data int) {
	var hari int
	fmt.Print("Masukkan berapa lama buku akan dipinjam (MAKS 7 HARI, CTH : (1/2/3/4/5/6/7)) : ")
	fmt.Scan(&hari)
	if hari <= 7 && hari > 0 {
		datapeminjam[data].tarifharga += 5000
		datapeminjam[data].TanggalKembali.tanggal += datapeminjam[data].TanggalPeminjaman.tanggal
		datapeminjam[data].TanggalKembali.bulan += datapeminjam[data].TanggalPeminjaman.bulan
		datapeminjam[data].TanggalKembali.tahun += datapeminjam[data].TanggalPeminjaman.tahun
		datapeminjam[data].TanggalKembali.tanggal += hari
		if datapeminjam[data].TanggalKembali.tanggal > 31 {
			datapeminjam[data].TanggalKembali.bulan += 1
			datapeminjam[data].TanggalKembali.tanggal -= 31
			if datapeminjam[data].TanggalKembali.bulan > 12 {
				datapeminjam[data].TanggalKembali.bulan -= 12
				datapeminjam[data].TanggalKembali.tahun += 1
			}
		}
		cls()
	} else {
		fmt.Println("Maaf, hari yang anda masukkan tidak valid!")
		fmt.Println("Silahkan untuk menginput ulang")
		pengisiandatatanggalkembali(data)
	}
}

func inputjudul(nData *int, nPeminjam *int, judul *string, i *int, data int) {
	var lanjut bool = false
	var pilih int
	fmt.Print("Judul buku yang akan dipinjam (Akhiri dengan spasi .): ")
	inputtext(&*judul)
	if mencaridata(*nData, *judul) != -1 && item[mencaridata(*nData, *judul)].stok > 0 {
		datapeminjam[*nPeminjam].Buku[*i] = *judul
		item[mencaridata(*nData, *judul)].rankterpinjam += 1
		item[mencaridata(*nData, *judul)].stok -= 1
	} else {
		for !lanjut {
			fmt.Println("Maaf, judul yang anda masukkan tidak ada/sudah habis stock")
			fmt.Println("Apakah anda ingin menginputkan kembali?")
			fmt.Println("1. \U00002705 Iya")
			fmt.Println("2. \U0000274c Tidak")
			fmt.Print("Pilihan Anda (1/2): ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				lanjut = true
				inputjudul(&*nData, &*nPeminjam, &*judul, &*i, data)
			case 2:
				lanjut = true
				item[mencaridata(*nData, *judul)].rankterpinjam -= *i
				item[mencaridata(*nData, *judul)].stok -= *i
				*i = data + 1
			default:
				fmt.Println("\nMaaf pilihan anda invalid, silahkan untuk menginputkan kembali")
				jedawaktu(3)
			}
		}
	}
}

func inputdate(nPeminjam *int) {
	inputtanggal(&*nPeminjam)
	inputbulan(&*nPeminjam)
	inputtahun(&*nPeminjam)
}

func inputtanggal(nPeminjam *int) {
	var temp int
	fmt.Print("Tanggal Peminjaman (1/2/3/4..../31): ")
	fmt.Scan(&temp)
	if temp <= 31 && temp > 0 {
		datapeminjam[*nPeminjam].TanggalPeminjaman.tanggal = temp
	} else {
		fmt.Println("Maaf tanggal yang anda inputkan invalid, silahkan inputkan ulang")
		jedawaktu(5)
		cls()
		inputtanggal(&*nPeminjam)
	}
}

func inputbulan(nPeminjam *int) {
	var temp int
	fmt.Print("\nBulan (cth: bulan januari dimasukkan '1'): ")
	fmt.Scan(&temp)
	if temp <= 12 && temp > 0 {
		datapeminjam[*nPeminjam].TanggalPeminjaman.bulan = temp
	} else {
		fmt.Println("Maaf bulan yang anda inputkan invalid, silahkan inputkan ulang")
		jedawaktu(5)
		inputbulan(&*nPeminjam)
	}
}

func inputtahun(nPeminjam *int) {
	var temp int
	fmt.Print("\nTahun (cth : '2005'): ")
	fmt.Scan(&temp)
	if temp > 0 {
		datapeminjam[*nPeminjam].TanggalPeminjaman.tahun = temp
	} else {
		fmt.Println("Maaf tahun yang anda inputkan invalid, silahkan inputkan ulang")
		jedawaktu(5)
		inputtahun(&*nPeminjam)
	}
}

func convertbulan(bulan int) string {
	if bulan == 1 {
		return "Januari"
	} else if bulan == 2 {
		return "Februari"
	} else if bulan == 3 {
		return "Maret"
	} else if bulan == 4 {
		return "April"
	} else if bulan == 5 {
		return "Mei"
	} else if bulan == 6 {
		return "Juni"
	} else if bulan == 7 {
		return "Juli"
	} else if bulan == 8 {
		return "Agustus"
	} else if bulan == 9 {
		return "September"
	} else if bulan == 10 {
		return "Oktober"
	} else if bulan == 11 {
		return "November"
	} else if bulan == 12 {
		return "Desember"
	} else {
		return "invalid"
	}
}

func menucaridatapinjaman(nData *int, nPeminjam *int) {
	cls()
	var pilih int
	fmt.Println("||=========================||")
	fmt.Println("||    CARI DATA PEMINJAM   ||")
	fmt.Println("||=========================||")
	fmt.Println("1. \U0001f50d Cari Nama Peminjam")
	fmt.Println("2. \U0001f50d\U0001f194 Cari nomor peminjam")
	fmt.Println("3. \U000025C0 Kembali")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		carinamapeminjam(&*nData, &*nPeminjam)
	case 2:
		carinomorpeminjam(&*nData, &*nPeminjam)
	case 3:
		menupinjaman(&*nData, &*nPeminjam)
	default:
		fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan ketik ulang")
		jedawaktu(3)
		cls()
		menucaridatapinjaman(&*nData, &*nPeminjam)
	}
}

func carinamapeminjam(nData *int, nPeminjam *int) {
	var lanjut bool = false
	var nama string
	var pilih int
	var wait byte
	fmt.Print("Nama Peminjam (Akhiri dengan spasi .): ")
	inputtext(&nama)
	fmt.Println()
	if caridatapeminjam(*nPeminjam, nama) != -1 {
		cetakdatapeminjam(*nPeminjam, caridatapeminjam(*nPeminjam, nama))
		fmt.Print("Klik Enter untuk melanjutkan....")
		fmt.Scanf("\n%c", &wait)
		menucaridatapinjaman(&*nData, &*nPeminjam)
	} else {
		for !lanjut {
			fmt.Println("Maaf nama peminjam yang anda cari tidak ada")
			fmt.Println("1. \U0001f50d Cari nama lain")
			fmt.Println("2. \U000025C0 Kembali")
			fmt.Print("Pilihan anda (1/2): ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				lanjut = true
				carinamapeminjam(&*nData, &*nPeminjam)
			case 2:
				lanjut = true
				menucaridatapinjaman(&*nData, &*nPeminjam)
			default:
				fmt.Println("\nMaaf pilihan anda invalid, silahkan input kembali")
				jedawaktu(3)
				cls()
			}
		}
	}
}

func carinomorpeminjam(nData *int, nPeminjam *int) {
	cls()
	var lanjut bool = false
	var nomor, pilih int
	var wait byte
	fmt.Print("Nomor Peminjam: ")
	fmt.Scan(&nomor)
	fmt.Println()
	if nomor-1 >= *nPeminjam {
		for !lanjut {
			fmt.Println("Maaf Nomor yang anda berikan tidak ada")
			fmt.Println("Apakah anda ingin menginputkan nya kembali?")
			fmt.Println("1. Iya")
			fmt.Println("2. Kembali")
			fmt.Print("Pilihan Anda (1/2):")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				lanjut = true
				carinomorpeminjam(&*nData, &*nPeminjam)
			case 2:
				lanjut = true
				menucaridatapinjaman(&*nData, &*nPeminjam)
			default:
				fmt.Println("\nMaaf pilihan anda invalid,silahkan untuk menginputkan kembali")
				jedawaktu(3)
				cls()
			}
		}
	} else {
		cetakdatapeminjam(*nPeminjam, nomor-1)
		fmt.Print("Klik Enter untuk melanjutkan....")
		fmt.Scanf("\n%c", &wait)
		menucaridatapinjaman(&*nData, &*nPeminjam)
	}
}

func caridatapeminjam(nPeminjam int, nama string) int {
	for i := 0; i < nPeminjam; i++ {
		if nama == datapeminjam[i].Namapeminjam {
			return i
		}
	}
	return -1
}

func datapinjaman(nData *int, nPeminjam *int) {
	cls()
	var wait byte
	if *nPeminjam == 0 {
		fmt.Println("\x1b[35mMaaf, data peminjam Anda masih kosong\x1b[0m")
		jedawaktu(5)
		menupinjaman(&*nData, &*nPeminjam)
	} else {
		fmt.Println("||=========================||")
		fmt.Println("||      DATA PEMINJAM      ||")
		fmt.Println("||=========================||")
		for i := 0; i < *nPeminjam; i++ {
			cetakdatapeminjam(*nPeminjam, i)
		}
		fmt.Print("Klik Enter untuk melanjutkan....")
		fmt.Scanf("\n%c", &wait)
		menupinjaman(&*nData, &*nPeminjam)
	}
}

func cetakdatapeminjam(nPeminjam int, data int) {
	fmt.Println("\x1b[32mNomor Peminjaman:\x1b[0m", datapeminjam[data].NomorPeminjaman)
	fmt.Println("\x1b[32mNama Peminjam:\x1b[0m", datapeminjam[data].Namapeminjam)
	fmt.Print("\x1b[32mBuku:\x1b[0m")
	for i := 0; i < datapeminjam[data].totalBukuDipinjam; i++ {
		fmt.Print(datapeminjam[data].Buku[i], ", ")
	}
	fmt.Println()
	fmt.Println("\x1b[32mTanggal Peminjaman:\x1b[0m", datapeminjam[data].TanggalPeminjaman.tanggal, "-", convertbulan(datapeminjam[data].TanggalPeminjaman.bulan), "-", datapeminjam[data].TanggalPeminjaman.tahun)
	fmt.Println("\x1b[32mTanggal Kembali:\x1b[0m", datapeminjam[data].TanggalKembali.tanggal, "-", convertbulan(datapeminjam[data].TanggalKembali.bulan), "-", datapeminjam[data].TanggalKembali.tahun)
	fmt.Println("\x1b[32mStatus:\x1b[0m", datapeminjam[data].Status)
	fmt.Println("\x1b[32mTarif Harga:\x1b[0m", datapeminjam[data].tarifharga)
	fmt.Println("\x1b[32mTotal Buku Dipinjam:\x1b[0m", datapeminjam[data].totalBukuDipinjam)
	fmt.Println("_______________________________________________")
}

func pengurutanterfavorit(nData int) {
	for i := 0; i < nData-1; i++ {
		for j := 0; j < nData-i-1; j++ {
			if item[j].rankterpinjam < item[j+1].rankterpinjam {
				item[j], item[j+1] = item[j+1], item[j]
			}
		}
	}
}

func Terfavorit(nData, nPeminjam *int) {
	cls()
	var wait byte
	pengurutanterfavorit(*nData)
	fmt.Println("||==============================||")
	fmt.Println("||   REKOMENDASI 5 TERFAVORIT   ||")
	fmt.Println("||==============================||")
	for i := 0; i < 5; i++ {
		cetakbuku(i)
	}
	fmt.Print("Klik Enter untuk melanjutkan....")
	fmt.Scanf("\n%c", &wait)
	menupinjaman(&*nData, &*nPeminjam)
}

func MeditPinjaman(nData, nPeminjam *int) {
	cls()
	var lanjut bool = false
	var judul string
	var pilih, data int
	fmt.Print("Masukkan Nama peminjam yang ingin di ubah (cth: 'Seiki .'):")
	inputtext(&judul)
	if caridatapeminjam(*nPeminjam, judul) != -1 {
		data = caridatapeminjam(*nPeminjam, judul)
		editPinjaman(&*nData, &*nPeminjam, data)
	} else {
		for !lanjut {
			fmt.Println("Maaf judul buku yang anda masukkan tidak ada")
			fmt.Println("1. \U000025B6 Input Kembali")
			fmt.Println("2. \U000025C0 Kembali")
			fmt.Print("\x1b[36mPilihan Anda (1/2): \x1b[0m")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				lanjut = true
				MeditPinjaman(&*nData, &*nPeminjam)
			case 2:
				lanjut = true
				menupinjaman(&*nData, &*nPeminjam)
			default:
				fmt.Println("\nMaaf pilihan anda invalid, silahkan input kembali")
				jedawaktu(3)
				cls()
			}
		}
	}
}

func editPinjaman(nData, nPeminjam *int, data int) {
	cls()
	var pilih int
	fmt.Println("Pilih apa yang akan di ubah:")
	fmt.Println("1. Nama Peminjam")
	fmt.Println("2. Buku Yang dipinjam")
	fmt.Println("3. Tanggal Pinjaman")
	fmt.Println("4. Semua")
	fmt.Println("5. Kembali")
	fmt.Print("Pilihan Anda (1/2/3/4/5/6): ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		editnamapeminjam(&*nData, &*nPeminjam, data)
		editPinjaman(&*nData, &*nPeminjam, data)
	case 2:
		editbukudipinjam(&*nData, &*nPeminjam, data)
		editPinjaman(&*nData, &*nPeminjam, data)
	case 3:
		edittanggalpeminjaman(&*nData, &*nPeminjam, data)
		editPinjaman(&*nData, &*nPeminjam, data)
	case 4:
		editnamapeminjam(&*nData, &*nPeminjam, data)
		editbukudipinjam(&*nData, &*nPeminjam, data)
		edittanggalpeminjaman(&*nData, &*nPeminjam, data)
		editPinjaman(&*nData, &*nPeminjam, data)
	case 5:
		menupinjaman(&*nData, &*nPeminjam)
	default:
		fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan ketik ulang")
		jedawaktu(3)
		cls()
		editPinjaman(&*nData, &*nPeminjam, data)
	}
}

func editnamapeminjam(nData, nPeminjam *int, data int) {
	cls()
	var nama string
	fmt.Print("Masukkan nama peminjam yang baru (cth: 'Imaachan .'):")
	inputtext(&nama)
	datapeminjam[data].Namapeminjam = nama
	cls()
	fmt.Println("Data peminjam baru: ")
	cetakdatapeminjam(*nPeminjam, data)
	jedawaktu(3)
}

func editbukudipinjam(nData, nPeminjam *int, data int) {
	cls()
	var lanjut bool = false
	var buku string
	var pilih int
	var status bool = true
	for i := 0; i < datapeminjam[data].totalBukuDipinjam; i++ {
		for !lanjut {
			fmt.Println("apakah buku ", datapeminjam[data].Buku[i], " akan diubah?")
			fmt.Println("1. Iya")
			fmt.Println("2. Tidak")
			fmt.Print("Pilihan Anda (1/2):")
			fmt.Scan(&pilih)
			if pilih == 1 {
				lanjut = true
				for status {
					cetakdatabuku(&*nData, &*nPeminjam)
					fmt.Print("Masukkan judul buku baru yang akan dipinjam (cth: 'Danur .'): ")
					inputtext(&buku)
					if mencaridata(*nData, buku) != -1 {
						item[mencaridata(*nData, datapeminjam[data].Buku[i])].stok += 1
						datapeminjam[data].Buku[i] = buku
						item[mencaridata(*nData, datapeminjam[data].Buku[i])].stok -= 1
						cls()
						fmt.Println("Data peminjam baru:")
						cetakdatapeminjam(*nPeminjam, data)
						status = false
					} else {
						fmt.Print("Maaf judul yang anda masukkan tidak ada, silahkan inputkan kembali")
					}
					jedawaktu(3)
					cls()
				}
				status = true
			} else if pilih == 2 {
				lanjut = true
			} else {
				fmt.Println("\nMaaf pilihan anda invalid, silahkan input kembali")
				jedawaktu(3)
				cls()
			}
		}
	}
}

func edittanggalpeminjaman(nData, nPeminjam *int, data int) {
	cls()
	inputdate(&data)
	pengisiandatatanggalkembali(data)
	datapeminjam[data].tarifharga = 0
	cls()
	fmt.Println("Data peminjam baru:")
	cetakdatapeminjam(*nPeminjam, data)
	jedawaktu(3)
}

// ! ===============================================

// ! Fungsi Pembagus/Cantik

// func cls() { // ? Procedure untuk menghapus atau clear terminal dengan kode
// 	fmt.Print("\033[2J")
// 	fmt.Print("\033[H")
// }

func cls() { // ? Procedure menghapus atau clear terminal dengan package os
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func drawProgressBar(progress float64) {
	bar := ""
	pos := int(progress * float64(width))
	for i := 0; i < width; i++ {
		if i < pos {
			bar += "█"
		} else {
			bar += " "
		}
	}
	bar += ""
	fmt.Printf("\r%s %3.0f%%", bar, progress*100)
}

func loading() { // ? Procedure untuk loadig screen
	cls()
	duration := 3 * time.Second

	fmt.Println("||==================================||")
	fmt.Println("||    ~ Mohon Tunggu Sebentar ~     ||")
	fmt.Println("||==================================||")

	startTime := time.Now()
	for {
		elapsed := time.Since(startTime)
		progress := elapsed.Seconds() / duration.Seconds()
		if progress > 1 {
			progress = 1
		}
		drawProgressBar(progress)
		time.Sleep(100 * time.Millisecond)
		if progress == 1 {
			break
		}
	}
	fmt.Println()
}

func jedawaktu(detik int) { // ? Procedure untuk jeda waktu
	for i := 0; i < (detik)*1000000000; i++ {

	}
}

func strukPeminjaman(data PinjamanBuku) {
	Filename := "strukPembayaran.txt"
	err := cetakStruk(data, Filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Membuka struk peminjaman")
	if err := exec.Command("notepad", Filename).Run(); err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func cetakStruk(p PinjamanBuku, namaFile string) error {
	file, err := os.Create(namaFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "===============================\n")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(file, "======= Struk Peminjaman ======\n")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(file, "===============================\n")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(file, "No Peminjam        : %d\n", p.NomorPeminjaman)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(file, "Nama Peminjam      : %s\n", p.Namapeminjam)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(file, "Tanggal Peminjaman     : %d%s%s%s%d\n", p.TanggalPeminjaman.tanggal, "-", convertbulan(p.TanggalPeminjaman.bulan), "-", p.TanggalPeminjaman.tahun)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(file, "Tanggal Pengembalian     : %d%s%s%s%d\n", p.TanggalKembali.tanggal, "-", convertbulan(p.TanggalKembali.bulan), "-", p.TanggalKembali.tahun)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(file, "Buku yang dipinjam : ")
	if err != nil {
		return err
	}
	for i := 0; i < p.totalBukuDipinjam; i++ {
		_, err = fmt.Fprintf(file, "%s ", p.Buku[i])
		if err != nil {
			return err
		}
	}
	fmt.Fprintln(file)
	_, err = fmt.Fprintf(file, "Jumlah             : %d\n", p.totalBukuDipinjam)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(file, "===============================\n")
	return err
}

func logoaplikasi() {

	var logo = []string{
		"     +   .   *          + .     .    ",
		"   .         .    ██    .    █     * ",
		"         +      ██  █       █ █      ",
		"  *  .        ███    ██   ██  █    . ",
		"       .    ████   ██████████ █      ",
		"     +     ████████████████████      ",
		"  .       ██████████████████████     ",
		"         ████████████████████████    ",
		"         ████████    ████    █████   ",
		"  ▄██████████████████████████████    ",
		"██████▀▀▀▀██████████████████████     ",
		"████        ██████████████████       ",
		"                                     ",
		"    ┌─┐       ┌─┬┐   ┌┐              ",
		"    │┼├─┬┬┬─┬┬┤─┤└┬─┐│├┬─┐┌─┐┌─┬┐    ",
		"    │┌┤┴┤┌┤┼││├─│┌┤┼└┤─┤┼└┤┼└┤│││    ",
		"    └┘└─┴┘│┌┴─┴─┴─┴──┴┴┴──┴──┴┴─┘    ",
		"          └┘                         ",
		"=====================================",
	}

	maxLen := 0
	for _, line := range logo {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	for i := 0; i <= maxLen; i++ {
		cls()
		for _, line := range logo {
			if len(line) > i {
				fmt.Println(line[:i])
			} else {
				fmt.Println(line)
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func headerMenuAdmin() {
	fmt.Println("  __  __ ___ _  _ _   _     _   ___  __  __ ___ _  _ ")
	fmt.Println(" |  \\/  | __| \\| | | | |   /_\\ |   \\|  \\/  |_ _| \\| |")
	fmt.Println(" | |\\/| | _|| .` | |_| |  / _ \\| |) | |\\/| || || .` |")
	fmt.Println(" |_|  |_|___|_| \\_\\___/  /_/ \\_\\___/|_|  |_|___|_|\\_|")
	fmt.Println("                                                     ")
	fmt.Println("======================================================")
}

func headerDatabuku() {
	fmt.Println("  ___   _ _____ _     ___ _   _ _  ___   _ ")
	fmt.Println(" |   \\ /_\\_   _/_\\   | _ ) | | | |/ / | | |")
	fmt.Println(" | |) / _ \\| |/ _ \\  | _ \\ |_| | ' <| |_| |")
	fmt.Println(" |___/_/ \\_\\_/_/ \\_\\ |___/\\___/|_|\\_\\\\___/ ")
	fmt.Println("============================================")
}

func headerCariDataBuku() {
	fmt.Println(`
    ___   _   ___ ___   ___   _ _____ _     ___ _   _ _  ___   _ 
   / __| /_\ | _ \_ _| |   \ /_\_   _/_\   | _ ) | | | |/ / | | |
  | (__ / _ \|   /| |  | |) / _ \| |/ _ \  | _ \ |_| | ' <| |_| |
   \___/_/ \_\_|_\___| |___/_/ \_\_/_/ \_\ |___/\___/|_|\_\\___/ 
=====================================================================`)
}

func headerEditDataBuku() {
	fmt.Println(`
====================================================================	
   ___ ___ ___ _____   ___   _ _____ _     ___ _   _ _  ___   _ 
   | __|   \_ _|_   _| |   \ /_\_   _/_\   | _ ) | | | |/ / | | |
   | _|| |) | |  | |   | |) / _ \| |/ _ \  | _ \ |_| | ' <| |_| |
   |___|___/___| |_|   |___/_/ \_\_/_/ \_\ |___/\___/|_|\_\\___/ 
====================================================================`)
}

func headerLogin() {
	fmt.Println("  _    ___   ___ ___ _  _     _   ___  __  __ ___ _  _ ")
	fmt.Println(" | |  / _ \\ / __|_ _| \\| |   /_\\ |   \\|  \\/  |_ _| \\| |")
	fmt.Println(" | |_| (_) | (_ || || .` |  / _ \\| |) | |\\/| || || .` |")
	fmt.Println(" |____\\___/ \\___|___|_|\\_| /_/ \\_\\___/|_|  |_|___|_|\\_|")
	fmt.Println("=======================================================")
}
