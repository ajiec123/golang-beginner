package main

import "fmt"

func main() {
	name := "Rahmat Pratama Aji"
	fmt.Println(name)

	name = "Rahmat Aji"
	fmt.Println("Isi variable name terbaru : ", name)
	fmt.Println("\n====================================\n")

	var (
		firstName  = "Rahmat"
		middleName = "Pratama"
		lastName   = "Aji"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

}
