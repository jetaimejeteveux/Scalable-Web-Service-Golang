package main

// Firman Aulia Insani
// GLNG-KS02-001
import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	name       string
	address    string
	job        string
	motivation string
}

var dataStudent []student

func init() {
	dataStudent = []student{
		{
			name:       "Firman Aulia Insani",
			address:    "Cilacap",
			job:        "mahasiswa",
			motivation: "dengan course ini saya berharap bisa mempelajari bidang baru, web khususnya backend karena selama ini hanya berkutat di android",
		},
		{
			name:       "Andri Nur Hidayatulloh",
			address:    "Jogja",
			job:        "mahasiswa",
			motivation: "Saya ikut course ini berharap ada product yg bisa saya buat dan membantu masyarakat",
		},
		{
			name:       "I Komang Widnyana",
			address:    "Bali",
			job:        "mahasiswa",
			motivation: "dengan course ini ku berharap golang bisa memantapkan kemampuan ku dibidang fullstack",
		},
		{
			name:       "Erico",
			address:    "Medan",
			job:        "mahasiswa",
			motivation: "Menambah ilmu baru",
		},
		{
			name:       "Jose Yolanda Purba",
			address:    "Medan",
			job:        "mahasiswa",
			motivation: "Menambah ilmu baru",
		},
		{
			name:       "Andri Kuwito",
			address:    "Bogor",
			job:        "mahasiswa",
			motivation: "Menambah ilmu baru",
		},
		{
			name:       "Sandy Budiman",
			address:    "Medan",
			job:        "mahasiswa",
			motivation: "saya gabung course ini ingin belajar programming, khususnya golang karena sedang hype dan high demand. Semoga bisa membuat product yg dapat bermanfaat bagi semua.",
		},
		{
			name:       "Rafli Andreansyah",
			address:    "Blitar",
			job:        "mahasiswa",
			motivation: "Menambah ilmu baru",
		},
		{
			name:       "Taufiq Hidayah",
			address:    "Yogyakarta",
			job:        "mahasiswa",
			motivation: "Menambah ilmu baru",
		},
		{
			name:       "Melvita Sari",
			address:    "Lhokseumawe",
			job:        "mahasiswa",
			motivation: "melalui course ini saya berharap bisa menambah ilmu dan memperbanyak teman",
		},
	}
}
func main() {
	args := os.Args[1]
	numberInput, _ := strconv.Atoi(args)
	numberInput -= 1
	switch numberInput {
	case 0:
		showData(numberInput)
	case 1:
		showData(numberInput)
	case 2:
		showData(numberInput)
	case 3:
		showData(numberInput)
	case 4:
		showData(numberInput)
	case 5:
		showData(numberInput)
	case 6:
		showData(numberInput)
	case 7:
		showData(numberInput)
	case 8:
		showData(numberInput)
	case 9:
		showData(numberInput)
	}
}

func showData(args int) {
	fmt.Printf("nama : %s \n address : %s \n job : %s \n motivation : %s", dataStudent[args].name, dataStudent[args].address, dataStudent[args].job, dataStudent[args].motivation)
}
