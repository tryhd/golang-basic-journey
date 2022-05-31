package main

import "fmt"

func main() {
	var nilaiUAS float32 = 80
	var nilaiTugas float32 = 75
	var absensi float32 = 70

	var nilaiAkhir float32 = (nilaiUAS * 0.5) + (nilaiTugas * 0.2) + (absensi * 0.3)

	if nilaiUAS < 50 || absensi < 80 {
		fmt.Println("Nilai UAS dan absensi kurang dari 50, maka tidak lulus")
	} else if nilaiAkhir >= 60 {
		fmt.Println("Nilai akhir anda adalah", nilaiAkhir, "Maka anda lulus")
	} else {
		fmt.Println("Nilai akhir anda adalah", nilaiAkhir, "Maka anda tidak lulus")
	}
}
