package ta

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

var expectedATR = []Value{
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
		Value: decimal.NewFromFloat(4.09086667),
	},
	{
		Time:  5,
		Value: decimal.NewFromFloat(4.99354444),
	},
	{
		Time:  6,
		Value: decimal.NewFromFloat(321.33826296),
	},
	{
		Time:  7,
		Value: decimal.NewFromFloat(513.05884198),
	},
	{
		Time:  8,
		Value: decimal.NewFromFloat(344.13589465),
	},
	{
		Time:  9,
		Value: decimal.NewFromFloat(231.723426311),
	},
	{
		Time:  10,
		Value: decimal.NewFromFloat(158.2961754),
	},
	{
		Time:  11,
		Value: decimal.NewFromFloat(135.99745027),
	},
}

func TestATR(t *testing.T) {
	actualATR := ATR(testHighSet, testCloseSet, testLowSet, testPeriod)

	require.Equal(t, len(expectedATR), len(actualATR), "compare len")
	for i := 0; i < len(actualATR); i++ {
		require.Equalf(t, expectedATR[i].Time, actualATR[i].Time, "compare time: %d", i)
		require.Equalf(t, expectedATR[i].Value.Round(4).String(), actualATR[i].Value.Round(4).String(), "compare : %d", actualATR[i].Time)
	}
}
