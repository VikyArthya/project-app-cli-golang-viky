package main

import (
	"fmt"
	"project-app-cli-golang-viky/rumus"
)

func main() {
	var pilihan = []int{1, 2, 3, 4, 5}
	var pilih int
	var l rumus.Lingkaran
	var p rumus.Persegi
	var pp rumus.PersegiPanjang
	var s3 rumus.Segitiga
	var k rumus.Ketupat
	fmt.Print("Masukkan Pilihan: ")
	fmt.Scan(&pilih)

	validation := false
	for i := 0; i < len(pilihan); i++ {
		if pilih == pilihan[i] {
			validation = true
			break
		}
	}

	if validation {
		if pilih == 1 {
			fmt.Print("masukkan jari jari: ")
			fmt.Scan(&l.JariJari)
			fmt.Println("Luas Lingkaran: ", l.Luas())
		} else if pilih == 2 {
			fmt.Print("Masukkan Sisi: ")
			fmt.Scan(&p.Sisi)
			fmt.Println("Luas Persegi: ", p.Luas())
		} else if pilih == 3 {
			fmt.Print("Masukkan Panjang: ")
			fmt.Scan(&pp.Panjang)
			fmt.Print("Masukkan Lebar: ")
			fmt.Scan(&pp.Lebar)
			fmt.Println("Luas Persegi Panjang: ", pp.Luas())
		} else if pilih == 4 {
			fmt.Print("Masukkan Alas: ")
			fmt.Scan(&s3.Alas)
			fmt.Print("Masukkan Tinggi: ")
			fmt.Scan(&s3.Tinggi)
			fmt.Println("Luas Segitiga: ", s3.Luas())
		} else if pilih == 5 {
			fmt.Print("Masukkan Diagonal 1: ")
			fmt.Scan(&k.Diagonal1)
			fmt.Print("Masukkan Diagonal 2: ")
			fmt.Scan(&k.Diagonal2)
			fmt.Println("Luas Ketupat: ", k.Luas())
		}
	} else {
		fmt.Print("Output salah")
	}

}
