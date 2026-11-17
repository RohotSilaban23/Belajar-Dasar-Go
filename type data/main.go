package main

import "fmt"

func main() {
	//Type data Numerik positif dan negatif
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	//Type data decimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	//type data B0olean
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)
}
