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
			Value: decimal.NewFromFloat(7.366666),
		},
		{
			Time:  3,
			Value: decimal.NewFromFloat(5.044211111),
		},
		{
			Time:  4,
			Value: decimal.NewFromFloat(3.9532074),
		},
		{
			Time:  5,
			Value: decimal.NewFromFloat(1.292904),
		},
		{
			Time:  6,
			Value: decimal.NewFromFloat(300.8619366),
		},
		{
			Time:  7,
			Value: decimal.NewFromFloat(202.90795775),
		},
		{
			Time:  8,
			Value: decimal.NewFromFloat(136.17530516689),
		},
		{
			Time:  9,
			Value: decimal.NewFromFloat(89.88020344),
		},
		{
			Time:  10,
			Value: decimal.NewFromFloat(61.82346896),
		},
		{
			Time:  11,
			Value: decimal.NewFromFloat(66.9189793),
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
