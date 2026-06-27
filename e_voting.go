package main

import "fmt"

type kandidat struct {
	No	int
	Nama string
	Visi string
	Suara int
}

func sequential(data []kandidat, no int) int {
	for i := 0; i < len(data); i++ {
		if data[i].No == no {
			return i
		}
	}
	return -1
}

func binary(data []kandidat, no int) int {
	kiri := 0
	kanan := len(data) - 1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if data[tengah].No == no {
			return tengah
		} else if data[tengah].No < no {
			kiri = tengah + 1
		} else {
			kiri = tengah - 1
		}
	}
	return -1
}

func selection(data []kandidat, kriteria int) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			// suara
			if kriteria == 1 {
				if data[j].Suara > data[min].Suara {
					min = j
				}
			// nomor urut
			} else {
				if data[j].No < data[min].No {
					min = j
				}
			}
		}
		data[i], data[min] = data[min], data[i]
	}
}

func insertion(data []kandidat, kriteria int) {
	n := len(data)
	for i := 1; i < n; i++ {
		x := data[i]
		j := i - 1
		// suara
		if kriteria == 1 {
			for j >= 0 && data[j].Suara < x.Suara {
				data[j+1] = data[j]
				j--
			}
		// nomor urut	
		} else {
			for j >= 0 && data[j].No > x.No {
				data[j+1] = data[j]
				j--
			}
		}
		data[j+1] = x
	}
}


func main() {
	var n int
	var no int

	dataKandidat := []kandidat{
		{1, "Anies - Muhaimin", "BBM gratis", 11111},
		{2, "Prabowo - Gibran", "Makan gratis", 22222},
		{3, "Ganjar - Mahfud", "Internet gratis", 33333},
	// "nomor" || "nama"  	||	"visi"    ||	"suara"
	}

	for {
		fmt.Println()
		fmt.Println("1. tambah kandidat")
		fmt.Println("2. edit kandidat")
		fmt.Println("3. hapus kandidat")
		fmt.Println("4. cari kandidat")
		fmt.Println("5. urutkann kandidat")
		fmt.Println("6. statistik suara")
		fmt.Println("7. voting")
		fmt.Println("8. selesai")
		fmt.Print("Pilih menu : ")
		fmt.Scanln(&n)

		if n == 8 {
			// fmt.Println("Program selesai.")
			break
		}

		switch n {
		// menambahkan	
		case 1: 
			var nama string
			var visi string
			var suara int

			fmt.Print("Masukkan nomor urut : ")
			fmt.Scanln(&no)

			if sequential(dataKandidat, no) != -1 {
				fmt.Println("Nomor urut sudah dipakai.")
				continue
			}

			fmt.Print("Masukkan nama kandidat : ")
			fmt.Scanln(&nama)

			fmt.Print("Masukkan visi : ")
			fmt.Scanln(&visi)

			fmt.Print("Masukkan suara awal : ")
			fmt.Scanln(&suara)

			dataKandidat = append(dataKandidat, kandidat {
				No: no,
				Nama: nama,
				Visi: visi,
				Suara: suara,
			})

			fmt.Println("kandidat telah ditambah")
			
		// mengubah
		case 2:
			fmt.Print("masukkan nomor urut kandidat :")
			fmt.Scanln(&no)
			
			posisi := sequential(dataKandidat, no)
			//?
			if posisi == -1 {
				fmt.Println("tidak valid")
			} else {
				fmt.Print("nama baru : ")
				fmt.Scanln(&dataKandidat[posisi].Nama)

				fmt.Print("Visi : ")
				fmt.Scanln(&dataKandidat[posisi].Visi)

				fmt.Println("kandidat telah ditambah")
			}

		// menghapus
		case 3:
			fmt.Print("masukkan nomor urut kandidat : ")
			fmt.Scanln(&no)

			posisi := sequential(dataKandidat, no)

			if posisi == -1 {
				fmt.Println("tidak valid")
			} else {
				dataKandidat = append(dataKandidat[:posisi], dataKandidat[posisi+1:]...) //?
				fmt.Println("kandidat telah dihapus")
			}

		// mencari
		case 4:
			var cari int

			if len(dataKandidat) == 0 {
				fmt.Println("data kosong")
				break
			}

			fmt.Print("nomor ynag dicari : ")
			fmt.Scanln(&no)

			fmt.Println("1. sequential ??")
			fmt.Println("2. binary ??")
			fmt.Scanln(&cari)

			posisi := -1

			if cari == 1 {
				posisi = sequential(dataKandidat, no)
			} else if cari == 2 {
				insertion(dataKandidat, 2)
				posisi = binary(dataKandidat, no)
			} else {
				fmt.Println("tidak valid.")
				continue
			}

			if posisi != -1 {
				fmt.Println()
				fmt.Println("no :", dataKandidat[posisi].No)
				fmt.Println("nama :", dataKandidat[posisi].Nama)
				fmt.Println("visi :", dataKandidat[posisi].Visi)
				fmt.Println("total suara :", dataKandidat[posisi].Suara)
			} else {
				fmt.Println("tiadk ditemukan")
			}

		case 5:
			var urutan int
			var sort int

			if len(dataKandidat) == 0 {
				fmt.Println("data kosong.")
				break
			}

			fmt.Println("1. Suara Terbanyak")
			fmt.Println("2. Nomor Urut")
			fmt.Scanln(&urutan) //?

			fmt.Println("1. selection")
			fmt.Println("2. insertion")
			fmt.Scanln(&sort)

			if sort == 1 {
				selection(dataKandidat, urutan)
			} else if sort == 2 {
				insertion(dataKandidat, urutan)
			} else {
				fmt.Println("tidak valid")
				break
			}

			fmt.Println("hasil")
			for i := 0; i < len(dataKandidat); i++ {
				fmt.Println("no :", dataKandidat[i].No)
				fmt.Println("nama  :", dataKandidat[i].Nama)
				fmt.Println("visi :", dataKandidat[i].Visi)
				fmt.Println("total suara  :", dataKandidat[i].Suara)
			}

		case 6:
			total := 0

			if len(dataKandidat) == 0 {
				fmt.Println("kosong")
				break
			}

			for i := 0; i < len(dataKandidat); i++ {
				total += dataKandidat[i].Suara
			}

			// fmt.Println("\nSTATISTIK SUARA")
			fmt.Println("total suara :", total)

			
			for i := 0; i < len(dataKandidat); i++ {
				persen := 0.0
			
				if total > 0 {
					persen = float64(dataKandidat[i].Suara) / float64(total) * 100
				}
				fmt.Printf("No %d - %s = %d suara (%.2f%%)\n",
					dataKandidat[i].No,
					dataKandidat[i].Nama,
					dataKandidat[i].Suara,
					persen)
			}

		case 7:
			fmt.Print("masukkan nomor urut : ")
			fmt.Scanln(&no)

			posisi := sequential(dataKandidat, no)

			if posisi == -1 {
				fmt.Println("tidak ditemukan")
			} else {
				dataKandidat[posisi].Suara++
				fmt.Println("voting selesai")
			}
			
		default:
			fmt.Println("pilihan tidak valid")
		}
	}
}
