package main

import (
	"fmt"
)

func main() {
	// Data skor untuk Tim Lumba-lumba
	skorLumbaLumba := [3]int{95, 105, 110}

	// Data skor untuk Tim Koala
	skorKoala := [3]int{100, 98, 120}

	// Hitung skor rata-rata untuk setiap tim
	rataRataLumbaLumba := hitungRataRata(skorLumbaLumba)
	rataRataKoala := hitungRataRata(skorKoala)

	// Bandingkan skor rata-rata untuk menentukan pemenang dari setiap tim
	if rataRataLumbaLumba > rataRataKoala && rataRataLumbaLumba >= 100 {
		fmt.Println("Pemenangnya adalah Tim Lumba-lumba!")
		fmt.Printf("Skor rata-rata Tim Lumba-lumba: %.2f\n", rataRataLumbaLumba)
		fmt.Printf("Skor rata-rata Tim Koala: %.2f\n", rataRataKoala)
	} else if rataRataKoala > rataRataLumbaLumba && rataRataKoala >= 100 {
		fmt.Println("Pemenangnya adalah Tim Koala!")
		fmt.Printf("Skor rata-rata Tim Lumba-lumba: %.2f\n", rataRataLumbaLumba)
		fmt.Printf("Skor rata-rata Tim Koala: %.2f\n", rataRataKoala)
	} else if rataRataLumbaLumba == rataRataKoala && rataRataLumbaLumba >= 100 && rataRataKoala >= 100 {
		fmt.Println("Hasilnya seri!")
		fmt.Printf("Skor rata-rata Tim Lumba-lumba: %.2f\n", rataRataLumbaLumba)
		fmt.Printf("Skor rata-rata Tim Koala: %.2f\n", rataRataKoala)
	} else {
		fmt.Println("Tidak ada pemenang karena skor tidak memenuhi syarat.")
	}
}

func hitungRataRata(skor [3]int) float64 {
	total := 0
	for _, nilai := range skor {
		total += nilai
	}
	return float64(total) / float64(len(skor))
}
