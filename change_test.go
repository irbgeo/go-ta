package ta_test

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"

	ta "github.com/irbgeo/gota"
)

var (
	testPeriodChange = 3

	testSrcSetChange = []ta.Value{
		{
			Time:  1,
			Value: decimal.NewFromFloat(7),
		},
		{
			Time:  2,
			Value: decimal.NewFromFloat(8.1),
		},
		{
			Time:  3,
			Value: decimal.NewFromFloat(0.3993),
		},
		{
			Time:  4,
			Value: decimal.NewFromFloat(1.7712),
		},
		{
			Time:  5,
			Value: decimal.NewFromFloat(-4.0277),
		},
		{
			Time:  6,
			Value: decimal.NewFromFloat(900),
		},
		{
			Time:  7,
			Value: decimal.NewFromFloat(7),
		},
		{
			Time:  8,
			Value: decimal.NewFromFloat(2.71),
		},
		{
			Time:  9,
			Value: decimal.NewFromFloat(-2.71),
		},
		{
			Time:  10,
			Value: decimal.NewFromFloat(5.71),
		},
		{
			Time:  11,
			Value: decimal.NewFromFloat(77.11),
		},
	}
)

var expectedChange = []ta.Value{
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
		Value: decimal.NewFromFloat(0),
	},
	{
		Time:  4,
		Value: decimal.NewFromFloat(-5.2288),
	},
	{
		Time:  5,
		Value: decimal.NewFromFloat(-12.1277),
	},
	{
		Time:  6,
		Value: decimal.NewFromFloat(899.6007),
	},
	{
		Time:  7,
		Value: decimal.NewFromFloat(5.2288),
	},
	{
		Time:  8,
		Value: decimal.NewFromFloat(6.7377),
	},
	{
		Time:  9,
		Value: decimal.NewFromFloat(-902.71),
	},
	{
		Time:  10,
		Value: decimal.NewFromFloat(-1.29),
	},
	{
		Time:  11,
		Value: decimal.NewFromFloat(74.4),
	},
}

func TestChange(t *testing.T) {
	actualChange := ta.Change(testSrcSetChange, testPeriodChange)

	require.Equal(t, len(expectedChange), len(actualChange), "compare len")
	for i := 0; i < len(actualChange); i++ {
		require.Equalf(t, expectedChange[i].Time, actualChange[i].Time, "compare time: %d", i)
		require.Equalf(t, expectedChange[i].Value.Round(4).String(), actualChange[i].Value.Round(4).String(), "compare : %d", actualChange[i].Time)
	}
}
