package main

import "fmt"

func main() {
	//deklarasi slice literal
	var x []int
	fmt.Println(x)

	fruits := []string{"mangga", "apel", "anggur"}
	fmt.Printf("element: %v, len: %d, cap:%d,", fruits, len(fruits), cap(fruits))
	//fmt.Println(fruits)

	fruits = append(fruits, "sawo")
	fmt.Printf("element: %v, len: %d, cap:%d\n", fruits, len(fruits), cap(fruits))

	fruits = append(fruits, "naga")
	fmt.Printf("element: %v, len: %d, cap:%d\n", fruits, len(fruits), cap(fruits))
	//kesimpulan len adalah jumlah item, sedangkan cap adalah kapasitas slice
	//kapasitas slice akan bertambah dua kali dari kapasitas awal apabila tidak cukup
	//manipulasi value

	fruits[1] = "salak"
	//fmt.Println(fruits)

	//deklarasi slice menggunakan make
	fruitsB := make([]string, 2)
	copy(fruitsB, fruits)
	fmt.Println(fruits)
	//fmt.Printf("fruits b nih %v", fruitsB)

	alphabet := []string{"a", "b", "c", "d", "e"}
	//slicing slice
	newAlphabet := alphabet[1:4]
	fmt.Println(newAlphabet)

	// [b c d]
	newAlphabet[1] = "z"
	fmt.Println(newAlphabet)
	fmt.Println(alphabet)
	//hati hati dengan memotong slice, karena dapat memotong storage dan dapat mempengaruhi variabel sebelumnya

	newAlphabet = append(newAlphabet, "z", "x", "t", "l")
	fmt.Printf("new alphabet %v", newAlphabet)
	fmt.Printf("alphabet  %v", alphabet)

	//copy mengcopy value
	//sedangkan slice mengcopy alamat memory
}
