package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func main() {
	name := []string{"kampus", "merdeka"}
	printMessage("hello", name)
	fmt.Println(randomWithRange(1, 10))
	var random int = randomWithRange(1, 10)
	var randomNew = randomWithRange(1, 10)
	randomNewNew := randomWithRange(2, 20)
	fmt.Println(random, randomNew, randomNewNew)
	diameter := 10
	luas, keliling := circleCarculate(float64(diameter))
	fmt.Println(luas, keliling)
	fmt.Println(calculateAvg(1, 2, 3, 4, 5))
}

//basic function
func printMessage(message string, inputName []string) {
	name := strings.Join(inputName, " ")
	fmt.Println("Hello World", name)

}

//function with 1 return value
func randomWithRange(min, max int) int {
	randomNum := rand.Int()%(max-min+1) + min
	return randomNum
}

//function with 2 return value
func circleCarculate(d float64) (float64, float64) {
	luas := math.Pi * math.Pow(d/2, 2)
	keliling := math.Phi * d
	return luas, keliling
}

//predefine function
func circleCarculatePre(d float64) (luasPre, kelilingPre float64) {
	luasPre = math.Pi * math.Pow(d/2, 2)
	kelilingPre = math.Phi * d
	return
}

//variadic function
func calculateAvg(angka ...int) float64 {
	result := 0
	for _, v := range angka {
		result += v
	}
	avgResult := result / len(angka)
	return float64(avgResult)
}
