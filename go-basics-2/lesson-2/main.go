package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"unicode/utf8"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {

	rstr := "abcdefghijklmnopqrstuvwxyz"
	pl("Rune Count:", utf8.RuneCountInString(rstr))

	for char, count := range rstr {
		pl(char, count)
		// char is the index of the rune and count is the rune itself
		pf("%d: %#U: %c\n", char, count, count)
	}

	now := time.Now()
	pl(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond())
	seedSecond := now.Unix()
	rand.Seed(seedSecond)
	randomNumber := rand.Intn(100) + 1
	pl(randomNumber)

	pl("5 * 5 =", 55*5)
	pl("5 / 5 =", 55/5)
	pl("5 % 5 =", 55%5)
	pl("5 + 5 =", 55+5)
	pl("5 - 5 =", 55-5)

	mint := 1
	mint++
	pl(mint)

	pl("Float Precision:", 1.0000000000000000000+3.0000000000000000000001)
	pl("Abs(-1):", math.Abs(-1))
	pl("Ceil(1.2):", math.Ceil(1.2))
	pl("Floor(1.2):", math.Floor(1.2))
	pl("Max(1,2):", math.Max(1, 2))
	pl("Min(1,2):", math.Min(1, 2))
	pl("Pow(2,3):", math.Pow(2, 3))
	pl("Sqrt(4):", math.Sqrt(4))
	pl("Round(1.5):", math.Round(1.5))

	r90 := math.Pi / 2
	pl("Sin(90):", math.Sin(r90))

	d90 := 90 * (math.Pi / 180)
	pl("Sin(90):", math.Sin(d90))

	pl("Pi:", math.Pi)

}
