package ta_test

import (
	"testing"

	ta "github.com/irbgeo/gota"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

var (
	testPeriodSMA = 2

	testSrcSetSMA = []ta.Value{ // Close
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

var expectedSMA = []ta.Value{
	{
		Time:  1,
		Value: decimal.NewFromFloat(0),
	},
	{
		Time:  2,
		Value: decimal.NewFromFloat(7.95000),
	},
	{
		Time:  3,
		Value: decimal.NewFromFloat(4.69715),
	},
	{
		Time:  4,
		Value: decimal.NewFromFloat(1.88275),
	},
	{
		Time:  5,
		Value: decimal.NewFromFloat(0.87175),
	},
	{
		Time:  6,
		Value: decimal.NewFromFloat(474.48615),
	},
	{
		Time:  7,
		Value: decimal.NewFromFloat(511.25000),
	},
	{
		Time:  8,
		Value: decimal.NewFromFloat(37.70500),
	},
	{
		Time:  9,
		Value: decimal.NewFromFloat(0.09450),
	},
	{
		Time:  10,
		Value: decimal.NewFromFloat(3.49450),
	},
	{
		Time:  11,
		Value: decimal.NewFromFloat(53.41000),
	},
}

func TestSMA(t *testing.T) {
	actualSMA := ta.SMA(testSrcSetSMA, testPeriodSMA)

	require.Equal(t, len(expectedSMA), len(actualSMA), "compare len")
	for i := 0; i < len(actualSMA); i++ {
		require.Equalf(t, expectedSMA[i].Time, actualSMA[i].Time, "compare time: %d", i+1)
		require.Equalf(t, expectedSMA[i].Value.Round(4).String(), actualSMA[i].Value.Round(4).String(), "compare : %d", i+1)
	}
}
