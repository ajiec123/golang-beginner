package main

import "fmt"

func main() {
	var firstName = "Rahmat"
	var lastName = "Aji"

	var result bool = firstName == lastName
	fmt.Println("Perbandingan == firstName dan lastName adalah :", result)

	var nilaiA = 100
	var nilaiB = 80
	var hasilNilai = nilaiA > nilaiB
	fmt.Println("Perbandingan > nilai A dan nilaiB adalah :", hasilNilai)
}
