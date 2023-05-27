package ta

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

var (
	expectedRMA = []Value{
		{
			Time:  1,
			Value: decimal.NewFromFloat(7),
		},
		{
			Time:  2,
			Value: decimal.NewFromFloat(7.55),
		},
		{
			Time:  3,
			Value: decimal.NewFromFloat(5.166433333333333),
		},
		{
			Time:  4,
			Value: decimal.NewFromFloat(4.034688888888889),
		},
		{
			Time:  5,
			Value: decimal.NewFromFloat(1.3472259259259258),
		},
		{
			Time:  6,
			Value: decimal.NewFromFloat(300.89815061728393),
		},
		{
			Time:  7,
			Value: decimal.NewFromFloat(202.9321004115226),
		},
		{
			Time:  8,
			Value: decimal.NewFromFloat(136.1914002743484),
		},
		{
			Time:  9,
			Value: decimal.NewFromFloat(89.89093351623228),
		},
		{
			Time:  10,
			Value: decimal.NewFromFloat(61.830622344154854),
		},
		{
			Time:  11,
			Value: decimal.NewFromFloat(66.92374822943657),
		},
	}
)

func TestRMA(t *testing.T) {
	actualRMA := RMA(testSrcSet, testPeriod)

	require.Equal(t, len(expectedRMA), len(actualRMA), "compare len")
	for i := 0; i < len(actualRMA); i++ {
		require.Equalf(t, expectedRMA[i].Time, actualRMA[i].Time, "compare time: %d", i+1)
		require.Equalf(t, expectedRMA[i].Value.Round(4).String(), actualRMA[i].Value.Round(4).String(), "compare : %d", i+1)
	}
}
