package main

import "fmt"

type calculate1 interface {
	calculate2d
	calculate3d
}

type calculate2d interface {
	hitungLuas() float64
}

type calculate3d interface {
	hitungVolume() float64
}

type kubus struct {
	sisi float64
}

func (k kubus) hitungLuas() float64 {
	return k.sisi * k.sisi * 6
}

func (k kubus) hitungVolume() float64 {
	return k.sisi * k.sisi * k.sisi
}

func main() {
	var bangunDatar calculate1 = kubus{10.0}
	fmt.Println("keliling lingkaran : ", bangunDatar.hitungLuas())
	fmt.Println("luas lingkaran : ", bangunDatar.hitungVolume())
}
