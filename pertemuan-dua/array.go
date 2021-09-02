package main

import "fmt"

func main() {
	//deklarasi array
	var x [5]int
	//menambah value ke array
	x[0] = 10
	fmt.Println(x)

	//deklarasi array2
	y := [3]int{1, 2, 3}

	//mengubah isi array
	y[1] = 90
	fmt.Println(y)

	//deklarasi array horizontal style
	fruit := [5]string{"mangga", "apel"}
	//skip isi array
	//deklarasi vertical style
	fruit2 := [5]string{
		"melon",
		"sawo",
		3: "semangka",
		"pepaya",
	}
	fruit2[2] = "anggur"
	fmt.Println(fruit)
	fmt.Println(fruit2)

	//deklarasi array dengan mengikuti value tidak harus deklarasi index terlebih dahulu
	//supaya tidak ada zero value

	animals := [...]string{
		"gajah",
		"semut",
	}
	animals[1] = "kijang"
	//kenapa harus deklarasi index dulu karena di golang array itu fixed
	//panjang array tidak bisa diubah tapi isi value bisa dimanipulasi

	//multidimensi array sebenernya ngga ada di golang tapi bisa disimulasi
	var z = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(z)

}
