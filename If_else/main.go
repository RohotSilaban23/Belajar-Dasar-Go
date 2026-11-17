package main

import "fmt"

func main() {
	var point = 8
	//Pengunaan kondisi IF ELSE secara sedderhana
	//dilihat berdasarakan kondisi yang memenuhi syarat
	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	//Variabel Temporary Pada if - else
	var point1 = 8840.0
	//Dimana dibuat sebuah variabel yang hanya
	//berlaku/digunakan dalam satu blok kondisi yang di tempatkan
	if percent := point1 / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	//Pengunaan switch-case dalam bentuk seperti IF-ELSE
	var point3 = 6

	switch {
	case point3 == 8:
		fmt.Println("perfect")
	case (point3 < 8) && (point3 > 3):
		fmt.Println("awesome")
		fallthrough //berfungsi untuk secara pasak meneruskan ke case berikut nya
		//meski case tersbut telah menemukan titik point nya
	case point3 < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	//IF-ELSE dalam bnetuk bersarang
	var point4 = 10

	if point4 > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default: //artinya di jika point3 dibawah 10 tapi di atas 7, makan akan di tampailkan
			fmt.Println("nice!")
		}
	} else {
		if point4 == 5 {
			fmt.Println("not bad")
		} else if point4 == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point4 == 0 {
				fmt.Println("try harder!")
			}
		}
	}

}
