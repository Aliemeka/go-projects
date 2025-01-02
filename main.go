package main

import (
	"fmt"
	"math/big"
	"strings"
	"time"
)

type petrolEngine struct {
	kpl    uint16
	litres uint16
}

func (e petrolEngine) distanceLeft() uint16 {
	return e.kpl * e.litres
}

type electricEngine struct {
	kpkwh uint16
	kwh   uint16
}

func (e electricEngine) distanceLeft() uint16 {
	return e.kpkwh * e.kwh
}

type engine interface {
	distanceLeft() uint16
}

func canIReachDestination(e engine, distance uint16) bool {
	return e.distanceLeft() >= distance
}

func main() {
	// Concatenate strings
	var letters = []string{"s", "u", "b", "c", "r", "i", "b", "e"}

	// Test time to concatenate strings
	duration, word := timeToConcat(letters)
	fmt.Println("It took", duration, "to concatenate the word:", word)

	// Test time using strings.Builder
	buildDuration, buildWord := timeUsingStringBuilder(letters)
	fmt.Println("It took", buildDuration, "to build the word:", buildWord)

	// Test time to join letters
	joinDuration, joinWord := timeUsingJoin(letters)
	fmt.Println("It took", joinDuration, "to join the word:", joinWord)

	var myEngine petrolEngine
	myEngine.kpl = 100
	myEngine.litres = 10
	fmt.Println("My engine capacity is", myEngine.kpl, "kpl and has", myEngine.litres, "litres of fuel")
	fmt.Println("I can travel", myEngine.distanceLeft(), "km with the fuel I have")

	var myElectricEngine electricEngine = electricEngine{kpkwh: 100, kwh: 10}
	fmt.Println("I can travel", myElectricEngine.distanceLeft(), "km with the battery I have")

	distance := int16(1000)
	var canReachWithPetrol bool = canIReachDestination(myEngine, uint16(distance))
	fmt.Println("I", map[bool]string{true: "can", false: "cannot"}[canReachWithPetrol], "reach my destination with the petrol engine")

	number := uint64(5_000_000)

	sum_with_loop, duration_with_loop := sum_nos_from_zero(number)
	fmt.Println("Sum of numbers from 0 to", number, "using loop is", sum_with_loop.String(), "and it took", duration_with_loop)

	sum_with_formula, duration_with_formula := sum_nos_from_zero_with_formula(number)
	fmt.Println("Sum of numbers from 0 to", number, "using formula is", sum_with_formula.String(), "and it took", duration_with_formula)

}

func sum_nos_from_zero(n uint64) (*big.Int, time.Duration) {
	t0 := time.Now()
	sum := big.NewInt(0)
	current := big.NewInt(0)
	for i := uint64(0); i <= n; i++ {
		current.SetUint64(i)
		sum.Add(sum, current)
	}
	return sum, time.Since(t0)
}

func sum_nos_from_zero_with_formula(n uint64) (*big.Int, time.Duration) {
	t0 := time.Now()
	bigN := big.NewInt(int64(n))
	bigNPlus1 := big.NewInt(0).Add(bigN, big.NewInt(1))
	product := big.NewInt(0).Mul(bigN, bigNPlus1)
	div := big.NewInt(0).Div(product, big.NewInt(2))
	return div, time.Since(t0)
}

func timeToConcat(letters []string) (duration time.Duration, outputString string) {
	t0 := time.Now()
	word := ""
	for i := range letters {
		word += letters[i]
	}
	return time.Since(t0), word
}

func timeUsingStringBuilder(letters []string) (duration time.Duration, outputString string) {
	t0 := time.Now()
	var stringBuilder strings.Builder
	for i := range letters {
		stringBuilder.WriteString(letters[i])
	}
	word := stringBuilder.String()
	return time.Since(t0), word
}

func timeUsingJoin(letters []string) (duration time.Duration, outputString string) {
	t0 := time.Now()
	word := strings.Join(letters, "")
	return time.Since(t0), word
}
