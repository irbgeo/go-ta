package ta

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

var (
	expectedRMA = []Value{
		{
			Time:  1,
			Value: decimal.NewFromFloat(2.3333333333333333),
		},
		{
			Time:  2,
			Value: decimal.NewFromFloat(4.2555555555555555),
		},
		{
			Time:  3,
			Value: decimal.NewFromFloat(2.970137037037037),
		},
		{
			Time:  4,
			Value: decimal.NewFromFloat(2.5704913580246913),
		},
		{
			Time:  5,
			Value: decimal.NewFromFloat(0.3710942386831275),
		},
		{
			Time:  6,
			Value: decimal.NewFromFloat(300.247396159122085),
		},
		{
			Time:  7,
			Value: decimal.NewFromFloat(202.49826410608139),
		},
		{
			Time:  8,
			Value: decimal.NewFromFloat(135.9021760707209267),
		},
		{
			Time:  9,
			Value: decimal.NewFromFloat(89.6981173804806178),
		},
		{
			Time:  10,
			Value: decimal.NewFromFloat(61.7020782536537452),
		},
		{
			Time:  11,
			Value: decimal.NewFromFloat(66.8380521691024968),
		},
	}
)

func TestRMA(t *testing.T) {
	actualRMA := RMA(testSrcSet, 1)

	fmt.Println(actualRMA)

	require.Equal(t, len(expectedRMA), len(actualRMA), "compare len")
	for i := 0; i < len(actualRMA); i++ {
		require.Equalf(t, expectedRMA[i].Time, actualRMA[i].Time, "compare time: %d", i+1)
		require.Equalf(t, expectedRMA[i].Value.Round(4).String(), actualRMA[i].Value.Round(4).String(), "compare : %d", i+1)
	}
}
