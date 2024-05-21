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

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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
	Buku              [4]string
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
	akun[0].Username = "Fauzan"
	akun[0].Password = "Stayhalal"

	akun[1].Username = "Imaachan"
	akun[1].Password = "Gkatauapa"

	akun[2].Username = "Seikiaccount "
	akun[2].Password = "Alhamdulillah "

	akun[3].Username = "reikisenpai "
	akun[3].Password = "Bismillah "

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
	fmt.Print("Pilihan Anda : ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		loading()
		Login(&*nData, &*nPeminjam)
	case 2:

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
	fmt.Println(" 3. \U0001F50D Panduan aplikasi")
	fmt.Println(" 4. \U000025C0 Log Out")
	fmt.Print("Pilihan Anda : ")
	fmt.Scan(&pilih)
	fmt.Println()
	switch pilih {
	case 1:
		menudatabuku(&*nData, &*nPeminjam)
	case 2:
		menupinjaman(&*nData, &*nPeminjam)
	case 3:
		panduan(&*nData, &*nPeminjam)
	case 4:

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

// ! ================================================================

// ! FUNGSI DATA BUKU

func menudatabuku(nData *int, nPeminjam *int) { // ? Procedure print output menu data buku
	cls()
	var pilih int
	headerDatabuku()
	fmt.Println("1. \U0001F4D6 Lihat data buku")
	fmt.Println("2. \U0001f58a Edit data buku")
	fmt.Println("3. \U000025C0 Kembali")
	fmt.Print("Pilihan Anda : ")
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
	var wait string
	if *nData == 0 {
		fmt.Println("\x1b[35mMaaf, data buku Anda masih kosong, silahkan untuk menginput dulu data buku Anda\x1b[0m")
		fmt.Println("1. \U000025B6 Lanjut")
		fmt.Println("2. \U000025C0 Kembali ke menu")
		fmt.Print("\x1b[36mPilihan Anda: \x1b[0m")
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
			fmt.Print("masukkan huruf apa saja untuk lanjut")
			fmt.Scan(&wait)
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
	fmt.Print("Pilihan Anda: ")
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
	var judul, wait string
	var pilih int
	fmt.Print("Masukkan Judul (akhiri dengan spasi (.)): ")
	inputtext(&judul)
	fmt.Println()
	if cekjudul(*nData, judul) {
		cls()
		cetakbuku(mencaridata(*nData, judul))
		fmt.Print("\nMasukkan kata apa saja untuk lanjut")
		fmt.Scan(&wait)
		menuadmin(&*nData, &*nPeminjam)
	} else {
		cls()
		fmt.Println("Maaf, Judul yang anda cari tidak ada")
		fmt.Println("Apakah Anda ingin mencari lagi judul yang lain?")
		fmt.Println("1. \U000025B6 Ya")
		fmt.Println("2. \U0001f4dd\U00002795 Tambahkan data buku")
		fmt.Println("3. \U000025C0 Tidak, kembali saja")
		fmt.Println("Jika buku yang anda cari tidak ada")
		fmt.Println("Silahkan untuk menambahkan data buku di menu edit buku")
		fmt.Print("Pilihan anda : ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			carijudul(&*nData, &*nPeminjam)
		case 2:
			menuMengedit(&*nData, &*nPeminjam)
		case 3:
			menucariData(&*nData, &*nPeminjam)
		}
	}
}

func carikatakunci(nData *int, nPeminjam *int) {
	cls()
	var pilih, batas, i, g, keprint int
	var kunci [5]string
	var tercetak [NMAX]string
	var temp string
	var cek bool = true
	keprint = 0
	batas = 0
	fmt.Println("Masukkan kata kunci anda, bila sudah, akhiri dengan . atau Maks 5 kata kunci")
	fmt.Println("! Jika anda ingin memasukkan dua kata, ketik dalam format camel case !")
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
		fmt.Print("Ketik Huruf Apapun lalu enter untuk melanjutkan: ")
		fmt.Scan(&temp)
		menuadmin(&*nData, &*nPeminjam)
	} else {
		cls()
		fmt.Println("Maaf, Key yang anda cari tidak ada")
		fmt.Println("Apakah Anda ingin mencari lagi key yang lain?")
		fmt.Println("1. \U000025B6 Ya")
		fmt.Println("2. \U000025C0 Tidak, kembali saja")
		fmt.Print("Pilihan anda : ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			carikatakunci(&*nData, &*nPeminjam)
		case 2:
			menucariData(&*nData, &*nPeminjam)
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
	var urut int = 0
	fmt.Println("||============================||")
	fmt.Println("||         DATA BUKU          ||")
	fmt.Println("||============================||")
	fmt.Println()
	for i := 0; i < *nData; i++ {
		if item[i].kategori == 1 {
			fmt.Printf("\x1b[35mNO: %d\n", urut+1)
			cetakbuku(i)
			urut += 1
		}
	}
	for i := 0; i < *nData; i++ {
		if item[i].kategori == 2 {
			fmt.Printf("\x1b[35mNO: %d\n", urut+1)
			cetakbuku(i)
			urut += 1
		}
	}
	for i := 0; i < *nData; i++ {
		if item[i].kategori == 3 {
			fmt.Printf("\x1b[35mNO: %d\n", urut+1)
			cetakbuku(i)
			urut += 1
		}
	}
	for i := 0; i < *nData; i++ {
		if item[i].kategori == 4 {
			fmt.Printf("\x1b[35mNO: %d\n", urut+1)
			cetakbuku(i)
			urut += 1
		}
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
		return "\x1b[44m \U0001F4AC Komik\x1b[0m"
	} else if i == 2 {
		return "\x1b[41m \U0001F3AD Novel\x1b[0m"
	} else if i == 3 {
		return "\x1b[42m \U0001F4C3 Cerpen\x1b[0m"
	} else if i == 4 {
		return "\x1b[43m \U0001F3EB Pelajaran : \x1b[0m"
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
	fmt.Println("3. \U000025C0 Kembali")
	fmt.Print("Pilihan Anda : ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		menambahkan(&*nData, &*nPeminjam)
	case 2:
		menghapus(&*nData, &*nPeminjam)
	case 3:
		menudatabuku(&*nData, &*nPeminjam)
	default:
		fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan ketik ulang")
		jedawaktu(3)
		cls()
		menuMengedit(&*nData, &*nPeminjam)
	}
}

func menambahkan(nData *int, nPeminjam *int) { // ? Procedure menambahkan buku kedalam data
	cls()
	var banyakstok, produk int
	fmt.Print("Masukkan jumlah buku yang akan ditambahkan: ")
	fmt.Scan(&banyakstok)
	for produk = 0; produk < banyakstok; produk++ {
		menambahkan2(&*nData, &*nPeminjam, &produk, &banyakstok)
	}
	cls()
	menuadmin(&*nData, &*nPeminjam)
}

func menambahkan2(nData *int, nPeminjam *int, produk, banyakstok *int) { // ? Subs procedure dari procedure menambahkan buku
	var kategori string
	var judul string
	fmt.Print("\n\x1b[36mJUDUL (Akhiri dengan spasi .): \x1b[0m")
	inputtext(&judul)
	if cekjudul(*nData, judul) {
		cls()
		pilihanmenambah(&*nData, &*nPeminjam, &*produk, &*banyakstok, &judul)
	} else {
		item[*nData].judul = judul
		fmt.Print("\x1b[36mPENULIS (Akhiri dengan spasi .): \x1b[0m")
		inputtext(&item[*nData].penulis)
		fmt.Print("\x1b[36mPENERBIT (Akhiri dengan spasi .): \x1b[0m")
		inputtext(&item[*nData].penerbit)
		fmt.Print("\x1b[36mGENRE (Akhiri dengan spasi .): \x1b[0m")
		inputtext(&item[*nData].genre)
		fmt.Print("\x1b[36mHARGA : \x1b[0m")
		fmt.Scan(&item[*nData].harga)
		fmt.Println("\x1b[36mKategori\x1b[0m")
		fmt.Println("\x1b[44m \U0001F4AC Komik\x1b[0m")
		fmt.Println("\x1b[41m \U0001F3AD Novel\x1b[0m")
		fmt.Println("\x1b[42m \U0001F4C3 Cerpen\x1b[0m")
		fmt.Println("\x1b[43m \U0001F3EB Pelajaran \x1b[0m")
		fmt.Print("Pilihan Anda : ")
		fmt.Scan(&kategori)
		switch kategori {
		case "Komik":
			item[*nData].kategori = 1
			fmt.Println("\x1b[32mBerhasil disimpan, Dalam kategori \U0001F4AC Komik\x1b[0m")
		case "Novel":
			item[*nData].kategori = 2
			fmt.Println("\x1b[32mBerhasil disimpan, Dalam kategori \U0001F3AD Novel\x1b[0m")
		case "Cerpen":
			item[*nData].kategori = 3
			fmt.Println("\x1b[32mBerhasil disimpan, Dalam kategori \U0001F4C3 Cerpen\x1b[0m")
		case "Pelajaran":
			item[*nData].kategori = 4
			fmt.Println("\x1b[32mBerhasil disimpan, Dalam kategori \U0001F3EB Pelajaran\x1b[0m")
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
	var temp string
	for temp != "." {
		fmt.Scan(&temp)
		if temp != "." {
			*text += temp + " "
		}
	}
}

func cekjudul(nData int, judul string) bool { // ? Procedure untuk mengecek apakah buku ada atau tidak berdasarkan judul buku
	var cek bool = false
	for i := 0; i < nData && cek != true; i++ {
		if judul == item[i].judul {
			cek = true
		}
	}
	return cek
}

func pilihanmenambah(nData *int, nPeminjam *int, produk, banyakstok *int, judul *string) { // ? Procedure pilihan jikalau buku yang ditambahkan judul nya sudah ada
	var pilih int
	fmt.Println("||======================================||")
	fmt.Println("||   JUDUL YANG ANDA PILIH SUDAH ADA    ||")
	fmt.Println("||======================================||")
	fmt.Println("Apakah Anda hanya ingin menambahkan stoknya?")
	fmt.Println("1. \U000025B6 Ya")
	fmt.Println("2. \U000025C0 Tidak, kembali ke menu")
	fmt.Print("Pilihan Anda : ")
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
	var data int
	for i < nData {
		if search == item[i].judul {
			data = i
		}
		i++
	}
	return data
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
		if cekjudul(*nData, judul) {
			simpandata = mencaridata(*nData, judul)
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
			jedawaktu(6)
			menuMengedit(&*nData, &*nPeminjam)
		}
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
	fmt.Println("2. \U0001f4cb Daftar Buku Dipinjam")
	fmt.Println("3. \U0001f4cb Rekomendasi 5 Buku Terfavorit")
	fmt.Println("4. \U0001f50d Pencarian Data pinjaman")
	fmt.Println("5. \U00002795 Perpanjangan Peminjaman")
	fmt.Println("6. \U0001f504 Pengembalian Buku")
	fmt.Println("7. \U0000fe0f Kembali")
	fmt.Print("Pilihan Anda : ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		pinjaman(&*nData, &*nPeminjam)
	case 2:
		datapinjaman(&*nData, &*nPeminjam)
	case 3:
		Terfavorit(&*nData, &*nPeminjam)
	case 4:
		menucaridatapinjaman(&*nData, &*nPeminjam)
	case 5:
		perpanjangan(&*nData, &*nPeminjam)
	case 6:
		menupengembalian(&*nData, &*nPeminjam)
	case 7:
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
	fmt.Println("3. \U0000fe0f Kembali")
	fmt.Print("Pilihan Anda: ")
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

func pengembalian(nData *int, nPeminjam *int) {
	cls()
	var pilih int
	var nama string
	var bulan, tahun, tanggal int
	var denda int
	fmt.Print("Nama peminjam: ")
	fmt.Scan(&nama)
	if ceknamapeminjam(nPeminjam, nama) {
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
		if tanggal < datapeminjam[caridatapeminjam(*nPeminjam, nama)].TanggalKembali.tanggal && tanggal > datapeminjam[caridatapeminjam(*nPeminjam, nama)].TanggalPeminjaman.tanggal {
			if bulan <= datapeminjam[caridatapeminjam(*nPeminjam, nama)].TanggalKembali.bulan {
				if tahun <= datapeminjam[caridatapeminjam(*nPeminjam, nama)].TanggalKembali.tahun {
					datapeminjam[caridatapeminjam(*nPeminjam, nama)].Status = "Done"
				} else {
					for i := 0; i < datapeminjam[caridatapeminjam(*nPeminjam, nama)].totalBukuDipinjam; i++ {
						denda += (item[mencaridata(*nData, datapeminjam[caridatapeminjam(*nPeminjam, nama)].Buku[i])].harga * 20) / 100
					}
				}
			} else {
				for i := 0; i < datapeminjam[caridatapeminjam(*nPeminjam, nama)].totalBukuDipinjam; i++ {
					denda += (item[mencaridata(*nData, datapeminjam[caridatapeminjam(*nPeminjam, nama)].Buku[i])].harga * 20) / 100
				}
			}
		} else {
			for i := 0; i < datapeminjam[caridatapeminjam(*nPeminjam, nama)].totalBukuDipinjam; i++ {
				denda += (item[mencaridata(*nData, datapeminjam[caridatapeminjam(*nPeminjam, nama)].Buku[i])].harga * 20) / 100
			}
		}
		if denda == 0 {

		} else {
			strukpembayaran(&denda)
		}
	} else {
		fmt.Println("Maaf Nama peminjam tidak ada")
		fmt.Println("Silahkan Untuk menginputkan dengan benar")
		fmt.Println("1. Input kembali")
		fmt.Println("2. Kembali")
		switch pilih {
		case 1:
			pengembalian(&*nData, &*nPeminjam)
		case 2:
			menupengembalian(&*nData, &*nPeminjam)
		}
	}
}

func hilang(nData *int, nPeminjam *int) {
	cls()
	var nama string
	var denda, pilih int
	fmt.Print("Nama peminjam: ")
	fmt.Scan(&nama)
	if ceknamapeminjam(nPeminjam, nama) {
		for i := 0; i < datapeminjam[caridatapeminjam(*nPeminjam, nama)].totalBukuDipinjam; i++ {
			fmt.Print("Apakah buku ", datapeminjam[caridatapeminjam(*nPeminjam, nama)].Buku[i], " hilang?")
			fmt.Print("1. \U00002705 Iya")
			fmt.Print("2. \U0000274c Tidak")
			fmt.Print("Pilihan Anda: ")
			fmt.Scan(&pilih)
			if pilih == 1 {
				denda += item[mencaridata(*nData, datapeminjam[caridatapeminjam(*nPeminjam, nama)].Buku[i])].harga
			}
		}
		strukpembayaran(&denda)
	}else {
		fmt.Println("Maaf Nama peminjam tidak ada")
		fmt.Println("Silahkan Untuk menginputkan dengan benar")
		fmt.Println("1. Input kembali")
		fmt.Println("2. Kembali")
		switch pilih {
		case 1:
			hilang(&*nData, &*nPeminjam)
		case 2:
			menupengembalian(&*nData, &*nPeminjam)
		}
	}
}

func strukpembayaran(tagihan *int) {
	var bayar int
	fmt.Print("Silahkan untuk membayar sebanyak ", *tagihan, ": ")
	fmt.Scan(&bayar)
	if bayar == *tagihan {
		fmt.Println("Lunas")
	} else if bayar > *tagihan {
		fmt.Println("Ini kembalian anda: ", bayar-*tagihan)
	} else {
		fmt.Println("Maaf uang anda kurang: ", *tagihan-bayar)
		bayar = *tagihan - bayar
		strukpembayaran(&bayar)
	}
}

func perpanjangan(nData *int, nPeminjam *int) {
	var nama string
	var hari int
	fmt.Print("Masukkan Nama saat anda meminjam buku: ")
	fmt.Scan(&nama)
	if ceknamapeminjam(nPeminjam, nama) {
		fmt.Print("Masukkan jumlah perpanjangan hari: ")
		fmt.Scan(&hari)
		datapeminjam[caridatapeminjam(*nPeminjam, nama)].tarifharga += 5000 * hari
		if (datapeminjam[caridatapeminjam(*nPeminjam, nama)].TanggalKembali.tanggal + hari) > 31 {
			datapeminjam[caridatapeminjam(*nPeminjam, nama)].TanggalKembali.tanggal -= 31
			datapeminjam[caridatapeminjam(*nPeminjam, nama)].TanggalKembali.tanggal += 1
		} else {
			datapeminjam[caridatapeminjam(*nPeminjam, nama)].TanggalKembali.tanggal += hari
		}
	}
}

func pinjaman(nData *int, nPeminjam *int) {
	cls()
	var judul, nama string
	var data, i, pilihan int
	cetakdatabuku(&*nData, &*nPeminjam)
	fmt.Print("\nNama (Akhiri dengan spasi .): ")
	inputtext(&nama)
	if ceknamapeminjam(nPeminjam, nama) {
		fmt.Println("Maaf, nama sudah dimiliki orang lain silahkan untuk menginput nama lain")
		jedawaktu(5)
		pinjaman(&*nData, &*nPeminjam)
	} else {
		fmt.Print("\nMau pinjam berapa buku (MAKS 5 Buku): ")
		fmt.Scan(&data)
		if data <= 5 {
			datapeminjam[*nPeminjam].Namapeminjam = nama
			datapeminjam[*nPeminjam].NomorPeminjaman = *nPeminjam + 1
			for i < data {
				inputjudul(&*nData, &*nPeminjam, &judul, &i, data)
				datapeminjam[*nPeminjam].totalBukuDipinjam = datapeminjam[*nPeminjam].totalBukuDipinjam + 1
				i++
			}
			if i > data {
				menupinjaman(&*nData, &*nPeminjam)
			} else {
				inputdate(&*nPeminjam)
				datapeminjam[*nPeminjam].Status = "Dipinjam"
				tarifharga(*nPeminjam)
				strukPeminjaman(datapeminjam[*nPeminjam])
				*nPeminjam += 1
				pengurutan(*nData)
				menuadmin(&*nData, &*nPeminjam)
			}
		} else {
			cls()
			fmt.Println("Maaf anda meminjam lebih dari batas")
			fmt.Println("Apakah anda ingin menginput kembali?")
			fmt.Println("1. Iya")
			fmt.Println("2. Tidak,Kembali saja")
			switch pilihan {
			case 1:
				pinjaman(&*nData, &*nPeminjam)
			case 2:
				menupinjaman(&*nData, &*nPeminjam)
			}
		}
	}
}

func tarifharga(data int) {
	var hari int
	fmt.Print("Masukkan batas waktu peminjaman: ")
	fmt.Scan(&hari)
	if hari <= 0 {
		fmt.Println("Maaf, hari yang anda masukkan tidak valid!")
		fmt.Println("Silahkan untuk menginput ulang kembali")
		tarifharga(data)
	} else {
		datapeminjam[data].tarifharga = hari * 5000
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
		strukpembayaran(&datapeminjam[data].tarifharga)
		datapeminjam[data].tarifharga = 0
		cls()
	}
}

func inputjudul(nData *int, nPeminjam *int, judul *string, i *int, data int) {
	var pilih int
	fmt.Print("Judul buku yang akan dipinjam (Akhiri dengan spasi .): ")
	*judul = ""
	inputtext(&*judul)
	fmt.Print(*judul)
	if cekjudul(*nData, *judul) && item[mencaridata(*nData, *judul)].stok > 0 {
		datapeminjam[*nPeminjam].Buku[*i] = *judul
		item[mencaridata(*nData, *judul)].rankterpinjam += 1
		item[mencaridata(*nData, *judul)].stok -= 1
	} else {
		fmt.Println(cekjudul(*nData, *judul)," ",item[mencaridata(*nData, *judul)].stok)
		fmt.Println("Maaf, judul yang anda masukkan tidak ada/sudah habis stock")
		fmt.Println("Apakah anda ingin menginputkan kembali?")
		fmt.Println("1. \U00002705 Iya")
		fmt.Println("2. \U0000274c Tidak")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			inputjudul(&*nData, &*nPeminjam, &*judul, &*i, data)
		case 2:
			item[mencaridata(*nData, *judul)].rankterpinjam -= *i
			item[mencaridata(*nData, *judul)].stok -= *i
			*i = data + 1
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
	fmt.Print("Tanggal : ")
	fmt.Scan(&temp)
	if temp <= 31 && temp > 0 {
		datapeminjam[*nPeminjam].TanggalPeminjaman.tanggal = temp
	} else {
		fmt.Println("Maaf tanggal yang anda inputkan invalid, silahkan inputkan ulang")
		jedawaktu(5)
		inputtanggal(&*nPeminjam)
	}
}

func inputbulan(nPeminjam *int) {
	var temp int
	fmt.Print("\nBulan : ")
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
	fmt.Print("\nTahun : ")
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
		fmt.Println("Maaf nomor yang anda pilih tidak ada, silahkan untuk mengetik ulang")
		jedawaktu(3)
		cls()
		menucaridatapinjaman(&*nData, &*nPeminjam)
	}
}

func carinamapeminjam(nData *int, nPeminjam *int) {
	var nama string
	fmt.Print("Nama Peminjam: ")
	fmt.Scan(&nama)
	fmt.Println()
	if ceknamapeminjam(nPeminjam, nama) {
		cetakdatapeminjam(*nPeminjam, caridatapeminjam(*nPeminjam, nama))
	}
}

func carinomorpeminjam(nData *int, nPeminjam *int) {
	var nomor int
	fmt.Print("Nomor Peminjam: ")
	fmt.Scan(&nomor)
	fmt.Println()
	if nomor >= *nPeminjam {
		fmt.Println("Maaf Nomor yang anda berikan tidak ada")
		jedawaktu(5)
		carinomorpeminjam(&*nData, &*nPeminjam)
	} else {
		cetakdatapeminjam(*nPeminjam, nomor-1)
	}
}

func ceknamapeminjam(nPeminjam *int, nama string) bool {
	var cek bool = false
	for i := 0; i < *nPeminjam; i++ {
		if nama == datapeminjam[i].Namapeminjam {
			cek = true
		}
	}
	return cek
}

func caridatapeminjam(nPeminjam int, nama string) int {
	var data int
	for i := 0; i < nPeminjam; i++ {
		if nama == datapeminjam[i].Namapeminjam {
			data = i
			i = nPeminjam + 1
		}
	}
	return data
}

func datapinjaman(nData *int, nPeminjam *int) {
	var wait int
	if *nPeminjam == 0 {
		fmt.Println("\x1b[35mMaaf, data peminjam Anda masih kosong, harap tunggu hingga ada yang meminjam\x1b[0m")
		jedawaktu(5)
		menuadmin(&*nData, &*nPeminjam)
	} else {
		fmt.Println("||=========================||")
		fmt.Println("||      DATA PEMINJAM      ||")
		fmt.Println("||=========================||")
		for i := 0; i < *nPeminjam; i++ {
			cetakdatapeminjam(*nPeminjam, i)
		}
		fmt.Print("Ketikkan huruf apa saja untuk melanjutkan")
		fmt.Scan(&wait)
		menuadmin(&*nData, &*nPeminjam)
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

func pengurutan(nData int) {
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
	var wait string
	fmt.Println("||==============================||")
	fmt.Println("||   REKOMENDASI 5 TERFAVORIT   ||")
	fmt.Println("||==============================||")
	for i := 0; i < 5; i++ {
		cetakbuku(i)
	}
	fmt.Print("Masukkan apa saja untuk lanjut...")
	fmt.Scan(&wait)
	menupinjaman(&*nData, &*nPeminjam)
}

// ! ===============================================

// ! FUNGSI PANDUAN

func panduan(nData *int, nPeminjam *int) {
	fmt.Println("Tata Cara Pengguna")
	fmt.Println("Tata Cara Pengguna")
	fmt.Println("Tata Cara Pengguna")
	fmt.Println("Tata Cara Pengguna")
	fmt.Println("Tata Cara Pengguna")
	fmt.Println("Tata Cara Pengguna")
	fmt.Println("Tata Cara Pengguna")
	menuadmin(&*nData, &*nPeminjam)
}

//! =============================================

// ! Fungsi Pembagus/Cantik

// func cls() { // ? Procedure untuk menghapus atau clear terminal dengan kode
// 	fmt.Print("\033[2J")
// 	fmt.Print("\033[H")
// }

func cls() { // ? Procedure menghapus atau clear terminal dengan package os
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

	}
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
	_, err = fmt.Fprintf(file, "Harga Pinjaman     : Rp %d\n", p.tarifharga)
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
	// 	░░░░┼░░░░░░░
	// ░░░░░░░░░░░░
	// ░░░░░░░░░░░┼
	// ░░░░░░░░░░░░
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
