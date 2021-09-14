package main

import "fmt"

func main() {
	for _, v := range []int{1, 2, 3, 4, 5} {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("terjadi kesalahan", v, "error")
				} else {
					fmt.Println("eksekusi sukses")
				}
			}()
			fmt.Println(v)
			panic("test recover")
		}()
	}
}
