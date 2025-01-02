package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
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
