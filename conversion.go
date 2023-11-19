package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println("\n")

	var name string = "Rahmat Pratama Aji"
	var r uint8 = name[0]
	var rString = string(r)

	fmt.Println(name)
	fmt.Println(r)
	fmt.Println(rString)

}
