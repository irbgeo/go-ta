package ta_test

import (
	"testing"

	ta "github.com/irbgeo/gota"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

var (
	testPeriodHighest = 3

	testSrcSetHighest = []ta.Value{
		{
			Time:  1,
			Value: decimal.NewFromFloat(10),
		},
		{
			Time:  2,
			Value: decimal.NewFromFloat(15),
		},
		{
			Time:  3,
			Value: decimal.NewFromFloat(12),
		},
		{
			Time:  4,
			Value: decimal.NewFromFloat(8),
		},
		{
			Time:  5,
			Value: decimal.NewFromFloat(20),
		},
		{
			Time:  6,
			Value: decimal.NewFromFloat(5),
		},
		{
			Time:  7,
			Value: decimal.NewFromFloat(18),
		},
		{
			Time:  8,
			Value: decimal.NewFromFloat(35),
		},
		{
			Time:  9,
			Value: decimal.NewFromFloat(9),
		},
		{
			Time:  10,
			Value: decimal.NewFromFloat(3),
		},
		{
			Time:  11,
			Value: decimal.NewFromFloat(36),
		},
	}
)

var expectedHighest = []ta.Value{
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
		Value: decimal.NewFromFloat(15),
	},
	{
		Time:  4,
		Value: decimal.NewFromFloat(15),
	},
	{
		Time:  5,
		Value: decimal.NewFromFloat(20),
	},
	{
		Time:  6,
		Value: decimal.NewFromFloat(20),
	},
	{
		Time:  7,
		Value: decimal.NewFromFloat(20),
	},
	{
		Time:  8,
		Value: decimal.NewFromFloat(35),
	},
	{
		Time:  9,
		Value: decimal.NewFromFloat(35),
	},
	{
		Time:  10,
		Value: decimal.NewFromFloat(35),
	},
	{
		Time:  11,
		Value: decimal.NewFromFloat(36),
	},
}

func TestHighest(t *testing.T) {
	actualHighest := ta.Highest(testSrcSetHighest, testPeriodHighest)

	require.Equal(t, len(expectedHighest), len(actualHighest), "compare len")
	for i := 0; i < len(actualHighest); i++ {
		require.Equalf(t, expectedHighest[i].Time, actualHighest[i].Time, "compare time: %d", i+1)
		require.Equalf(t, expectedHighest[i].Value.Round(4).String(), actualHighest[i].Value.Round(4).String(), "compare : %d", i+1)
	}
}
