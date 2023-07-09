package ta_test

import (
	"testing"

	ta "github.com/irbgeo/gota"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

var (
	testPeriodStDev = 2

	testSrcSetStDev = []ta.Value{
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

var expectedStDev = []ta.Value{
	{
		Time:  1,
		Value: decimal.NewFromFloat(0),
	},
	{
		Time:  2,
		Value: decimal.NewFromFloat(0.45000),
	},
	{
		Time:  3,
		Value: decimal.NewFromFloat(3.70285),
	},
	{
		Time:  4,
		Value: decimal.NewFromFloat(0.88845),
	},
	{
		Time:  5,
		Value: decimal.NewFromFloat(1.89945),
	},
	{
		Time:  6,
		Value: decimal.NewFromFloat(475.51385),
	},
	{
		Time:  7,
		Value: decimal.NewFromFloat(438.75000),
	},
	{
		Time:  8,
		Value: decimal.NewFromFloat(34.79500),
	},
	{
		Time:  9,
		Value: decimal.NewFromFloat(2.81550),
	},
	{
		Time:  10,
		Value: decimal.NewFromFloat(6.21550),
	},
	{
		Time:  11,
		Value: decimal.NewFromFloat(43.7),
	},
}

func TestStDev(t *testing.T) {
	actualStDev := ta.StDev(testSrcSetStDev, testPeriodStDev)

	require.Equal(t, len(expectedStDev), len(actualStDev), "compare len")
	for i := 0; i < len(actualStDev); i++ {
		require.Equalf(t, expectedStDev[i].Time, actualStDev[i].Time, "compare time: %d", i+1)
		require.Equalf(t, expectedStDev[i].Value.Round(4).String(), actualStDev[i].Value.Round(4).String(), "compare : %d", i+1)
	}
}
