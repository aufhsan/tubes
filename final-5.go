package main

import "fmt"

const maxPakaian = 100
const maxOutfit = 100

type pakaian struct {
	id         string
	nama       string
	kategori   string
	warna      string
	formalitas int
	lastUsed   string // Format: YYYY-MM-DD
}

type outfit struct {
	id       string
	nama     string
	atas     pakaian
	bawah    pakaian
	kategori string // Kategori: "casual", "formal", "sporty"
	lastUsed string
}

var (
	pakaianList     [maxPakaian]pakaian
	outfitList      [maxOutfit]outfit
	pakaianCount    int
	outfitCount     int
	idCounter       int
	outfitIDCounter int
)

func main() {
	fmt.Println("\nAplikasi Manajemen Fashion Sederhana")
	fmt.Println("Notes: Gunakan Huruf Kecil Semua")
	menuUtama()
}

func menuUtama() {
	for {
		fmt.Println("\n=== Menu Utama ===")
		fmt.Println("1. Manajemen Pakaian (Tambah/Ubah/Hapus/Lihat)")
		fmt.Println("2. Manajemen Outfit (Tambah/Ubah/Hapus/Lihat)")
		fmt.Println("3. Pencarian (Warna/Kategori)")
		fmt.Println("4. Pengurutan (Formalitas/Terakhir Dipakai)")
		fmt.Println("5. Rekomendasi (Cuaca/Acara)")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			menuManajemenPakaian()
		case 2:
			menuManajemenOutfit()
		case 3:
			pencarian()
		case 4:
			pengurutan()
		case 5:
			rekomendasi()
		case 0:
			fmt.Println("\nTerima Kasih Sudah Menggunakan Aplikasi Ini")
			return
		default:
			fmt.Println("\nPilihan tidak valid!")
		}
	}
}

//==============MANAJEMEN PAKAIAN=============
func menuManajemenPakaian() {
	for {
		fmt.Println("\n=== Manajemen Pakaian ===")
		fmt.Println("1. Tambah Pakaian")
		fmt.Println("2. Ubah Pakaian")
		fmt.Println("3. Hapus Pakaian")
		fmt.Println("4. Lihat Daftar Pakaian")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilihan: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahPakaian()
		case 2:
			ubahPakaian()
		case 3:
			hapusPakaian()
		case 4:
			tampilkanPakaian()
		case 0:
			return
		default:
			fmt.Println("\nPilihan tidak valid!")
		}
	}
}

func tambahPakaian() {
	if pakaianCount >= maxPakaian {
		fmt.Println("\nDatabase pakaian penuh! Tidak bisa menambah lagi.")
		return
	}

	var p pakaian
	idCounter++
	p.id = fmt.Sprintf("P%03d", idCounter)

	fmt.Print("\nNama Pakaian: ")
	fmt.Scan(&p.nama)

	fmt.Print("Kategori (casual/formal/sporty): ")
	fmt.Scan(&p.kategori)

	fmt.Print("Warna: ")
	fmt.Scan(&p.warna)

	fmt.Print("Tingkat Formalitas (1-5): ")
	fmt.Scan(&p.formalitas)

	fmt.Print("Terakhir Dipakai (YYYY-MM-DD): ")
	fmt.Scan(&p.lastUsed)

	pakaianList[pakaianCount] = p
	pakaianCount++
	fmt.Println("\nPakaian berhasil ditambahkan!")
}

func ubahPakaian() {
	tampilkanPakaian()
	if pakaianCount == 0 {
		return
	}

	var id string
	var i int
	var found bool
	found = false
	fmt.Print("\nMasukkan ID pakaian yang akan diubah: ")
	fmt.Scan(&id)

	for i = 0; i < pakaianCount && !found; i++ {
		if pakaianList[i].id == id {
			found = true
			fmt.Printf("\nNama Pakaian (%s): ", pakaianList[i].nama)
			fmt.Scan(&pakaianList[i].nama)

			fmt.Printf("Kategori (%s): ", pakaianList[i].kategori)
			fmt.Scan(&pakaianList[i].kategori)

			fmt.Printf("Warna (%s): ", pakaianList[i].warna)
			fmt.Scan(&pakaianList[i].warna)

			fmt.Printf("Tingkat Formalitas (%d): ", pakaianList[i].formalitas)
			fmt.Scan(&pakaianList[i].formalitas)

			fmt.Printf("Terakhir Dipakai (%s): ", pakaianList[i].lastUsed)
			fmt.Scan(&pakaianList[i].lastUsed)

			fmt.Println("\nPakaian berhasil diubah!")
		}
	}

	if !found {
		fmt.Println("\nPakaian dengan ID tersebut tidak ditemukan!")
	}
}

func hapusPakaian() {
	tampilkanPakaian()
	if pakaianCount == 0 {
		fmt.Println("\nTidak ada pakaian untuk dihapus.")
		return
	}

	var id string
	fmt.Print("\nMasukkan ID pakaian yang akan dihapus: ")
	fmt.Scan(&id)

	var i, j int
	var found bool
	found = false
	for i = 0; i < pakaianCount && !found; i++ {
		if pakaianList[i].id == id {
			found = true

			for j = i; j < pakaianCount-1; j++ {
				pakaianList[j] = pakaianList[j+1]
			}
			pakaianCount--

			for j = 0; j < pakaianCount; j++ {
				pakaianList[j].id = fmt.Sprintf("P%03d", j+1)
			}

			idCounter = pakaianCount

			fmt.Println("\nPakaian berhasil dihapus!")
			tampilkanPakaian()
		}
	}

	if !found {
		fmt.Println("\nPakaian dengan ID tersebut tidak ditemukan!")
	}
}

func tampilkanPakaian() {
	var i int
	var p pakaian

	if pakaianCount == 0 {
		fmt.Println("\nBelum ada pakaian yang terdaftar!")
		return
	}

	fmt.Println("\nDaftar Pakaian:")
	for i = 0; i < pakaianCount; i++ {
		p = pakaianList[i]
		fmt.Printf("ID: %s, Nama: %s, Kategori: %s, Warna: %s, Formalitas: %d, Terakhir Dipakai: %s\n",
			p.id, p.nama, p.kategori, p.warna, p.formalitas, p.lastUsed)
	}
}

//==============MANAJEMEN OUTFIT=============
func menuManajemenOutfit() {
	for {
		fmt.Println("\n=== Manajemen Outfit ===")
		fmt.Println("1. Tambah Outfit")
		fmt.Println("2. Ubah Outfit")
		fmt.Println("3. Hapus Outfit")
		fmt.Println("4. Lihat Daftar Outfit")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilihan: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahOutfit()
		case 2:
			ubahOutfit()
		case 3:
			hapusOutfit()
		case 4:
			tampilkanOutfit()
		case 0:
			return
		default:
			fmt.Println("\nPilihan tidak valid!")
		}
	}
}

func tambahOutfit() {
	if outfitCount >= maxOutfit {
		fmt.Println("\nDatabase outfit penuh! Tidak bisa menambah lagi.")
		return
	}

	var o outfit
	outfitIDCounter++
	o.id = fmt.Sprintf("%04d", outfitIDCounter)

	tampilkanPakaian()
	fmt.Print("\nNama Outfit: ")
	fmt.Scan(&o.nama)

	fmt.Print("Masukkan ID Pakaian Atas: ")
	var idAtas string
	fmt.Scan(&idAtas)
	o.atas = getPakaianByID(idAtas)

	fmt.Print("Masukkan ID Pakaian Bawah: ")
	var idBawah string
	fmt.Scan(&idBawah)
	o.bawah = getPakaianByID(idBawah)

	fmt.Print("Kategori Outfit (casual/formal/sporty): ")
	fmt.Scan(&o.kategori)

	fmt.Print("Terakhir Dipakai (YYYY-MM-DD): ")
	fmt.Scan(&o.lastUsed)

	outfitList[outfitCount] = o
	outfitCount++
	fmt.Println("\nOutfit berhasil ditambahkan!")
}

func ubahOutfit() {
	tampilkanOutfit()
	if outfitCount == 0 {
		return
	}

	var id string
	fmt.Print("\nMasukkan ID outfit yang akan diubah: ")
	fmt.Scan(&id)

	var i int
	tampilkanPakaian()
	for i = 0; i < outfitCount; i++ {
		if outfitList[i].id == id {
			fmt.Printf("\nNama Outfit (%s): ", outfitList[i].nama)
			fmt.Scan(&outfitList[i].nama)

			fmt.Printf("Masukkan ID Pakaian Atas Baru (ID: %s, %s): ", outfitList[i].atas.id, outfitList[i].atas.nama)
			var idAtas string
			fmt.Scan(&idAtas)
			outfitList[i].atas = getPakaianByID(idAtas)

			fmt.Printf("Masukkan ID Pakaian Bawah Baru (ID: %s, %s): ", outfitList[i].bawah.id, outfitList[i].bawah.nama)
			var idBawah string
			fmt.Scan(&idBawah)
			outfitList[i].bawah = getPakaianByID(idBawah)

			fmt.Printf("Kategori (%s): ", outfitList[i].kategori)
			fmt.Scan(&outfitList[i].kategori)

			fmt.Printf("Tanggal lastUsed (%s): ", outfitList[i].lastUsed)
			fmt.Scan(&outfitList[i].lastUsed)

			fmt.Println("\nOutfit berhasil diubah!")
			return
		}
	}
	fmt.Println("\nOutfit dengan ID tersebut tidak ditemukan!")
}

func hapusOutfit() {
	var i, j int
	var id string
	var found bool

	tampilkanOutfit()
	if outfitCount == 0 {
		fmt.Println("\nTidak ada outfit untuk dihapus.")
		return
	}

	fmt.Print("\nMasukkan ID outfit yang akan dihapus: ")
	fmt.Scan(&id)

	found = false
	for i = 0; i < outfitCount; i++ {
		if outfitList[i].id == id {
			found = true

			for j = i; j < outfitCount-1; j++ {
				outfitList[j] = outfitList[j+1]
			}
			outfitCount--

			fmt.Println("\nOutfit berhasil dihapus!")

			for j = 0; j < outfitCount; j++ {
				outfitList[j].id = fmt.Sprintf("%04d", j+1)
			}

			outfitIDCounter = outfitCount
			tampilkanOutfit()
			return
		}
	}

	// Jika tidak ditemukan
	if !found {
		fmt.Println("\nOutfit dengan ID tersebut tidak ditemukan!")
	}
}

func tampilkanOutfit() {
	if outfitCount == 0 {
		fmt.Println("\nBelum ada outfit yang terdaftar!")
		return
	}

	fmt.Println("\nDaftar Outfit:")
	var i int
	var o outfit
	for i = 0; i < outfitCount; i++ {
		o = outfitList[i]
		fmt.Printf("ID: %s | %s\nKategori: %s\nAtas: %s (%s)\nBawah: %s (%s)\nTerakhir Dipakai: %s\n\n",
			o.id, o.nama, o.kategori,
			o.atas.nama, o.atas.kategori,
			o.bawah.nama, o.bawah.kategori,
			o.lastUsed)
	}
}

func getPakaianNama(id string) string {
	var i int
	for i = 0; i < pakaianCount; i++ {
		if pakaianList[i].id == id {
			return pakaianList[i].nama
		}
	}
	return "Tidak Ditemukan"
}

func getPakaianByID(id string) pakaian {
	var i int
	for i = 0; i < pakaianCount; i++ {
		if pakaianList[i].id == id {
			return pakaianList[i]
		}
	}
	return pakaian{}

}

//==============PENCARIAN=====================
func pencarian() {
	for {
		fmt.Println("\n=== Pencarian ===")
		fmt.Println("1. Berdasarkan Warna Pakaian")
		fmt.Println("2. Berdasarkan Kategori Pakaian")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilihan: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			pencarianWarna()
		case 2:
			pencarianKategori()
		case 0:
			return
		default:
			fmt.Println("\nPilihan tidak valid!")
		}
	}
}

// 1. SEQUENTIAL SEARCH - pencarianWarna()
func pencarianWarna() {

	var searchPakaian [maxPakaian]pakaian
	var i int
	for i = 0; i < pakaianCount; i++ {
		searchPakaian[i] = pakaianList[i]
	}

	var warna string
	fmt.Print("\nMasukkan warna yang dicari (huruf kecil semua): ")
	fmt.Scan(&warna)

	fmt.Println("\nHasil Pencarian Pakaian (Sequential Search):")
	var found bool
	found = false
	for i = 0; i < pakaianCount; i++ {
		if searchPakaian[i].warna == warna {
			var p pakaian
			p = searchPakaian[i]
			fmt.Printf("ID: %s, Nama: %s, Kategori: %s, Warna: %s\n",
				p.id, p.nama, p.kategori, p.warna)
			found = true
		}
	}

	if !found {
		fmt.Println("\nTidak ditemukan pakaian dengan warna tersebut!")
	}
}

// 2. BINARY SEARCH - pencarianKategori()
func pencarianKategori() {

	var searchPakaian [maxPakaian]pakaian
	var i int
	for i = 0; i < pakaianCount; i++ {
		searchPakaian[i] = pakaianList[i]
	}

	var kategori string
	fmt.Print("\nMasukkan kategori yang dicari (casual/formal/sporty): ")
	fmt.Scan(&kategori)

	var j int
	for i = 0; i < pakaianCount-1; i++ {
		for j = 0; j < pakaianCount-1-i; j++ {

			if searchPakaian[j].kategori > searchPakaian[j+1].kategori {
				var temp pakaian
				temp = searchPakaian[j]
				searchPakaian[j] = searchPakaian[j+1]
				searchPakaian[j+1] = temp
			}
			/*
				MENGGUNAKAN PERBANDINGAN ASCII
				Huruf "c": 99
				Huruf "f": 102
				Huruf "s": 115
			*/

		}
	}

	//AMANNNNNN
	fmt.Println("\nHasil Pencarian Pakaian (Binary Search):")
	var left, right, mid int
	left = 0
	right = pakaianCount - 1
	var found bool
	found = false

	for left <= right && !found {
		mid = (left + right) / 2
		if searchPakaian[mid].kategori == kategori {
			var idx int
			idx = mid
			for idx >= left && searchPakaian[idx].kategori == kategori {
				var p pakaian
				p = searchPakaian[idx]
				fmt.Printf("ID: %s, Nama: %s, Kategori: %s, Warna: %s\n",
					p.id, p.nama, p.kategori, p.warna)
				idx--
			}
			idx = mid + 1
			for idx <= right && searchPakaian[idx].kategori == kategori {
				var p pakaian
				p = searchPakaian[idx]
				fmt.Printf("ID: %s, Nama: %s, Kategori: %s, Warna: %s\n",
					p.id, p.nama, p.kategori, p.warna)
				idx++
			}
			found = true
		} else if searchPakaian[mid].kategori < kategori {
			left = mid + 1
		} else {
			right = mid - 1
		}
		/*
			MENGGUNAKAN PERBANDINGAN ASCII
			Huruf "c": 99
			Huruf "f": 102
			Huruf "s": 115
		*/
	}

	if !found {
		fmt.Println("\nTidak ditemukan pakaian dengan kategori tersebut!")
	}
}

//==============PENGURUTAN=====================
func pengurutan() {
	for {
		fmt.Println("\n=== Pengurutan ===")
		fmt.Println("1. Pakaian Berdasarkan Formalitas (Selection Sort)")
		fmt.Println("2. Pakaian Berdasarkan Terakhir Dipakai (Insertion Sort)")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilihan: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			pengurutanFormalitas()
		case 2:
			pengurutanLastUsed()
		case 0:
			return
		default:
			fmt.Println("\nPilihan tidak valid!")
		}
	}
}

// 3. SELECTION SORT - pengurutanFormalitas()
func pengurutanFormalitas() {
	if pakaianCount == 0 {
		fmt.Println("\nBelum ada pakaian yang terdaftar!")
		return
	}

	var sortedPakaian [maxPakaian]pakaian
	var i int
	for i = 0; i < pakaianCount; i++ {
		sortedPakaian[i] = pakaianList[i]
	}

	var n, j int
	n = pakaianCount
	for i = 0; i < n-1; i++ {
		var minIdx int
		minIdx = i
		for j = i + 1; j < n; j++ {
			if sortedPakaian[j].formalitas < sortedPakaian[minIdx].formalitas {
				minIdx = j
			}
		}
		var temp pakaian
		temp = sortedPakaian[i]
		sortedPakaian[i] = sortedPakaian[minIdx]
		sortedPakaian[minIdx] = temp
	}

	fmt.Println("\nHasil Pengurutan Pakaian (Selection Sort - Formalitas):")
	for i = 0; i < pakaianCount; i++ {
		var p pakaian
		p = sortedPakaian[i]
		fmt.Printf("ID: %s, Nama: %s, Formalitas: %d\n", p.id, p.nama, p.formalitas)
	}
}

// 4. INSERTION SORT - pengurutanLastUsed()
func pengurutanLastUsed() {
	if pakaianCount == 0 {
		fmt.Println("\nBelum ada pakaian yang terdaftar!")
		return
	}

	var sortedPakaian [maxPakaian]pakaian
	var i int
	for i = 0; i < pakaianCount; i++ {
		sortedPakaian[i] = pakaianList[i]
	}

	var n, j int
	n = pakaianCount
	for i = 1; i < n; i++ {
		var key pakaian
		key = sortedPakaian[i]
		j = i - 1

		for j >= 0 && sortedPakaian[j].lastUsed < key.lastUsed {
			sortedPakaian[j+1] = sortedPakaian[j]
			j--
		}
		sortedPakaian[j+1] = key
	}

	fmt.Println("\nHasil Pengurutan Pakaian (Insertion Sort - Terakhir Dipakai):")
	for i = 0; i < pakaianCount; i++ {
		var p pakaian
		p = sortedPakaian[i]
		fmt.Printf("ID: %s, Nama: %s, Terakhir Dipakai: %s\n", p.id, p.nama, p.lastUsed)
	}

}

//==============REKOMENDASI=====================
func rekomendasi() {
	for {
		fmt.Println("\n=== Rekomendasi ===")
		fmt.Println("1. Berdasarkan Cuaca")
		fmt.Println("2. Berdasarkan Acara")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilihan: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			rekomendasiCuaca()
		case 2:
			rekomendasiAcara()
		case 0:
			return
		default:
			fmt.Println("\nPilihan tidak valid!")
		}
	}
}

func rekomendasiCuaca() {
	if outfitCount == 0 {
		fmt.Println("\nBelum ada outfit yang terdaftar untuk direkomendasikan.")
		return
	}

	var cuaca string
	fmt.Print("\nMasukkan jenis cuaca (hujan/panas/dingin): ")
	fmt.Scan(&cuaca)

	fmt.Println("\nRekomendasi Outfit untuk Cuaca", cuaca)

	var i int
	var foundRecommendation bool
	foundRecommendation = false

	for i = 0; i < outfitCount; i++ {
		var o outfit
		o = outfitList[i]
		var isRecommended bool
		isRecommended = false

		if cuaca == "hujan" {
			if o.nama == "hujan" ||
				o.atas.nama == "jaket" ||
				o.atas.nama == "jaket_hujan" ||
				o.atas.nama == "jaket_waterproof" ||
				o.atas.nama == "raincoat" ||
				o.atas.nama == "hoodie" ||
				o.atas.nama == "sweater" ||
				o.atas.nama == "jas_hujan" ||
				o.atas.nama == "hoodie_tahan_air" ||
				o.atas.nama == "ponco" ||
				o.bawah.nama == "celana_jas_hujan" ||
				o.bawah.nama == "celana_panjang_waterproof" ||
				o.bawah.nama == "celana_hujan" ||
				o.bawah.nama == "jogger" ||
				o.bawah.nama == "celana_panjang" {
				isRecommended = true
			}
		} else if cuaca == "panas" {
			if o.nama == "panas" ||
				o.atas.nama == "kaos" ||
				o.atas.nama == "kemeja_lengan_pendek" ||
				o.atas.nama == "tank_top" ||
				o.atas.nama == "jersey" ||
				o.atas.nama == "blouse_ringan" ||
				o.atas.nama == "kemeja_katun" ||
				o.atas.nama == "crop_top" ||
				o.bawah.nama == "rok_pendek" ||
				o.bawah.nama == "celana_kain_tipis" ||
				o.bawah.nama == "celana_palazzo" ||
				o.bawah.nama == "celana_capri" ||
				o.bawah.nama == "celana_chino" ||
				o.bawah.nama == "celana_pendek" ||
				o.bawah.nama == "rok" {
				isRecommended = true
			}
		} else if cuaca == "dingin" {
			if o.nama == "dingin" ||
				o.atas.nama == "sweater" ||
				o.atas.nama == "jaket" ||
				o.atas.nama == "jaket_wol" ||
				o.atas.nama == "jaket_tebal" ||
				o.atas.nama == "hoodie" ||
				o.atas.nama == "cardigan" ||
				o.atas.nama == "thermal_inner" ||
				o.atas.nama == "jaket_windbreaker" ||
				o.atas.nama == "coat" ||
				o.atas.nama == "vest_berbulu" ||
				o.bawah.nama == "celana_panjang" ||
				o.bawah.nama == "celana_jeans" ||
				o.bawah.nama == "celana_fleece" ||
				o.bawah.nama == "celana_corduroy" ||
				o.bawah.nama == "celana_jeans_berlapis" ||
				o.bawah.nama == "legging_thermal" {
				isRecommended = true
			}
		}

		if isRecommended {
			fmt.Printf("ID: %s, Nama: %s (Atas: %s, Bawah: %s)\n", o.id, o.nama, o.atas.nama, o.bawah.nama)
			foundRecommendation = true
		}
	}

	if !foundRecommendation {
		fmt.Println("\nTidak ada rekomendasi outfit untuk cuaca tersebut.")
	}
}

func rekomendasiAcara() {
	if outfitCount == 0 {
		fmt.Println("\nBelum ada outfit yang terdaftar!")
		return
	}

	var acara string
	fmt.Print("\nMasukkan jenis acara (casual/formal/sporty): ")
	fmt.Scan(&acara)

	fmt.Println("\nRekomendasi Outfit untuk Acara", acara+":")

	var found bool
	var i int
	var o outfit
	found = false

	for i = 0; i < outfitCount; i++ {
		if outfitList[i].kategori == acara {
			o = outfitList[i]
			fmt.Printf("ID: %s, Nama: %s\nAtas: %s (%s)\nBawah: %s (%s)\nLast Used: %s\n\n",
				o.id, o.nama,
				o.atas.nama, o.atas.kategori,
				o.bawah.nama, o.bawah.kategori,
				o.lastUsed)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada outfit dengan kategori", acara)
	}
}

/*
NOTES:
formalitas untuk FORMAL = 4-5
formalitas untuk CASUAL = 2-3
formalitas untuk OLAHRAGA = - (jika ada atasan "sporty" atau bawahan "sporty")
*/
