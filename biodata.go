package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func getTemanByAbsen(absen int) Teman {
	var teman Teman
	switch absen {
	case 1:
		teman = Teman{
			Nama:      "Aira",
			Alamat:    "Bukittinggi",
			Pekerjaan: "Developer",
			Alasan:    "Ingin mempelajari bahasa pemrograman",
		}
	case 2:
		teman = Teman{
			Nama:      "Nisa",
			Alamat:    "Padang",
			Pekerjaan: "QA",
			Alasan:    "Ingin memahami konsep programming secara mendalam",
		}
	case 3:
		teman = Teman{
			Nama:      "Hariri",
			Alamat:    "Pariaman",
			Pekerjaan: "Data Analyst",
			Alasan:    "Ingin mengembangkan keterampilan programming",
		}
	default:
		teman = Teman{}
	}
	return teman
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Masukkan nomor absen yang ingin ditampilkan")
		return
	}

	absen, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka")
		return
	}

	teman := getTemanByAbsen(absen)
	if teman.Nama == "" {
		fmt.Println("Tidak ditemukan nomor absen tersebut")
		return
	}

	fmt.Printf("Nama: %s\n", teman.Nama)
	fmt.Printf("Alamat: %s\n", teman.Alamat)
	fmt.Printf("Pekerjaan: %s\n", teman.Pekerjaan)
	fmt.Printf("Alasan: %s\n", teman.Alasan)
}
