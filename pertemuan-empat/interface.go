package main

import (
	"fmt"
	"math"
)

type calculate interface {
	hitungKeliling() float64
	hitungLuas() float64
}

type lingkaran struct {
	diameter float64
}

type persegi struct {
	sisi float64
}

func (l lingkaran) hitungLuas() float64 {
	return math.Pi * math.Pow(l.diameter/2, 2)
}
func (l lingkaran) hitungKeliling() float64 {
	return math.Pi * math.Pow(l.diameter/2, 2)
}
func (p persegi) hitungLuas() float64 {
	return p.sisi * p.sisi
}
func (p persegi) hitungKeliling() float64 {
	return p.sisi * 4
}
func main() {
	var bangunDatar calculate = lingkaran{10.0}
	fmt.Println("keliling lingkaran : ", bangunDatar.hitungKeliling())
	fmt.Println("luas lingkaran : ", bangunDatar.hitungKeliling())
	bangunDatar = persegi{10.0}
	fmt.Println("keliling persegi : ", bangunDatar.hitungKeliling())
	fmt.Println("luas persegi : ", bangunDatar.hitungKeliling())

}
