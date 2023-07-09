package ta_test

import (
	"testing"

	ta "github.com/irbgeo/gota"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

var (
	testPeriodEMA = 3

	testSrcSetEMA = []ta.Value{ // Close
		{
			Time:  1,
			Value: decimal.NewFromFloat(7.5),
		},
		{
			Time:  2,
			Value: decimal.NewFromFloat(8.4),
		},
		{
			Time:  3,
			Value: decimal.NewFromFloat(0.9943),
		},
		{
			Time:  4,
			Value: decimal.NewFromFloat(2.7712),
		},
		{
			Time:  5,
			Value: decimal.NewFromFloat(-1.0277),
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
)

var expectedEMA = []ta.Value{
	{
		Time:  1,
		Value: decimal.NewFromFloat(0),
	},
	{
		Time:  2,
		Value: decimal.NewFromFloat(0),
	},
	{
		Time:  3,
		Value: decimal.NewFromFloat(5.631433),
	},
	{
		Time:  4,
		Value: decimal.NewFromFloat(4.201317),
	},
	{
		Time:  5,
		Value: decimal.NewFromFloat(1.586808),
	},
	{
		Time:  6,
		Value: decimal.NewFromFloat(475.793404),
	},
	{
		Time:  7,
		Value: decimal.NewFromFloat(274.146702),
	},
	{
		Time:  8,
		Value: decimal.NewFromFloat(138.528351),
	},
	{
		Time:  9,
		Value: decimal.NewFromFloat(67.903676),
	},
	{
		Time:  10,
		Value: decimal.NewFromFloat(38.806838),
	},
	{
		Time:  11,
		Value: decimal.NewFromFloat(67.958419),
	},
}

func TestEMA(t *testing.T) {
	actualEMA := ta.EMA(testSrcSetEMA, testPeriodEMA)

	require.Equal(t, len(expectedEMA), len(actualEMA), "compare len")
	for i := 0; i < len(actualEMA); i++ {
		require.Equalf(t, expectedEMA[i].Time, actualEMA[i].Time, "compare time: %d", i+1)
		require.Equalf(t, expectedEMA[i].Value.Round(4).String(), actualEMA[i].Value.Round(4).String(), "compare : %d", i+1)
	}
}
