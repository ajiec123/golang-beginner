package main

import "fmt"

func main() {
	var nilaiA = 100
	var nilaiB = 80
	var nilaiC = 70
	var nilaiD = 50
	var nilaiE = 30

	var penambahan_variable = nilaiA + nilaiB + nilaiC + nilaiD + nilaiE
	fmt.Println("Nilai dari operasi matematika adalah :", penambahan_variable, "\n")

	// augmented assignments
	// operasi untuk menambahkan nilai kepada variable atau dirinya sendiri

	var augNilaiA = 10
	augNilaiA += 5 // augNilaiA = 10 + 5

	fmt.Println("Nilai dari augmented assignment augNilaiA adalah :", augNilaiA)

	// unary operator
	// ++ , -- , - , + , !

	var j = 10
	j++ // j = 10 + 1
	fmt.Println("Hasil dari variable J adalah :", j)

}
