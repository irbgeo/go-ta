package ta

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

var expectedChange = []Value{
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
	actualChange := Change(testSet, testPeriod)

	require.Equal(t, len(expectedChange), len(actualChange), "compare len")
	for i := 0; i < len(actualChange); i++ {
		require.Equalf(t, expectedChange[i].Time, actualChange[i].Time, "compare time: %d", i)
		require.Equalf(t, expectedChange[i].Value.Round(4).String(), actualChange[i].Value.Round(4).String(), "compare : %d", actualChange[i].Time)
	}
}
