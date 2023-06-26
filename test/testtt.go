package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Value struct {
	Time  int
	Value decimal.Decimal
}

func main() {
	testHighSet1 := []Value{
		{
			Time:  1,
			Value: decimal.NewFromFloat(25740.00000000),
		},
		{
			Time:  2,
			Value: decimal.NewFromFloat(25740.00000000),
		},
		{
			Time:  3,
			Value: decimal.NewFromFloat(25725.02000000),
		},
		{
			Time:  4,
			Value: decimal.NewFromFloat(25718.88000000),
		},
		{
			Time:  5,
			Value: decimal.NewFromFloat(25720.16000000),
		},
		{
			Time:  6,
			Value: decimal.NewFromFloat(950),
		},
		{
			Time:  7,
			Value: decimal.NewFromFloat(72.5),
		},
		{
			Time:  8,
			Value: decimal.NewFromFloat(2.91),
		},
		{
			Time:  9,
			Value: decimal.NewFromFloat(-2.721),
		},
		{
			Time:  10,
			Value: decimal.NewFromFloat(9.71),
		},
		{
			Time:  11,
			Value: decimal.NewFromFloat(97.11),
		},
	}

	high := []float64{
		25740.00000000,
		25740.00000000,
		25725.02000000,
		25718.88000000,
		25720.16000000,
		25730.84000000,
		25727.51000000,
		25728.99000000,
		25724.77000000,
		25712.65000000,
		25714.84000000,
	}

	for i := range high {
		f, _ := testHighSet1[i].Value.Float64()
		if high[i] > f {
			testHighSet1[i].Value = decimal.NewFromFloat(high[i])
		}
	}

	fmt.Println(testHighSet1)
}
