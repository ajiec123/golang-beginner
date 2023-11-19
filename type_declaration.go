package main

import (
	"fmt"
)

func main() {

	type NoKTP string

	var ktpAji NoKTP = "3175032323232123"

	var contoh string = "1231231231231231"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(ktpAji)
	fmt.Println(contohKTP)
}
