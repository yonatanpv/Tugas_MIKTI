package main

import (
	"fmt"
)

func main() {
	// Data 1

	// Data Mark
	beratMark1 := 78.0  // kg
	tinggiMark1 := 1.69 // meter

	// Data Jhon
	beratJohn1 := 92.0  // kg
	tinggiJohn1 := 1.95 // meter

	// Hitung BMI untuk data 1
	bmiMark1 := hitungBMI(beratMark1, tinggiMark1)
	bmiJohn1 := hitungBMI(beratJohn1, tinggiJohn1)

	// Bandingkan BMI
	markHigherBMI1 := bmiMark1 > bmiJohn1

	// Output hasil untuk data 1
	fmt.Printf("Data 1:\n")
	fmt.Printf("BMI Mark: %.2f\n", bmiMark1)
	fmt.Printf("BMI John: %.2f\n", bmiJohn1)
	fmt.Printf("Apakah Mark memiliki BMI lebih tinggi dari John? %t\n\n", markHigherBMI1)

	// Data 2

	// Data Mark
	beratMark2 := 95.0  // kg
	tinggiMark2 := 1.88 // meter

	// DataJhon
	beratJohn2 := 85.0  // kg
	tinggiJohn2 := 1.76 // meter

	// Hitung BMI untuk data 2
	bmiMark2 := hitungBMI(beratMark2, tinggiMark2)
	bmiJohn2 := hitungBMI(beratJohn2, tinggiJohn2)

	// Bandingkan BMI
	markHigherBMI2 := bmiMark2 > bmiJohn2

	// Output hasil untuk data 2
	fmt.Printf("Data 2:\n")
	fmt.Printf("BMI Mark: %.2f\n", bmiMark2)
	fmt.Printf("BMI John: %.2f\n", bmiJohn2)
	fmt.Printf("Apakah Mark memiliki BMI lebih tinggi dari John? %t\n", markHigherBMI2)
}

func hitungBMI(berat, tinggi float64) float64 {
	return berat / (tinggi * tinggi)
}
