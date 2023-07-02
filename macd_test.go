package ta_test

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"

	ta "github.com/irbgeo/gota"
)

var (
	testShortPeriodMACD  = 3
	testLongPeriodMACD   = 5
	testSignalPeriodMACD = 10

	testSrcSetMACD = []ta.Value{
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

var (
	expectedMACD = []ta.Value{
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

	expectedSignal = []ta.Value{
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

	expectedHist = []ta.Value{
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
)

func TestMACD(t *testing.T) {
	actualMACD, actualSignal, actualHist := ta.MACD(testSrcSetMACD, testShortPeriodMACD, testLongPeriodMACD, testSignalPeriodMACD)

	require.Equal(t, len(expectedMACD), len(actualMACD), "compare len macd")
	require.Equal(t, len(expectedSignal), len(actualSignal), "compare len signal")
	require.Equal(t, len(expectedHist), len(actualHist), "compare len hist")

	for i := 0; i < len(actualMACD); i++ {
		require.Equalf(t, expectedMACD[i].Time, actualMACD[i].Time, "compare time: %d", i)
		require.Equalf(t, expectedMACD[i].Value.Round(4).String(), actualMACD[i].Value.Round(4).String(), "compare : %d", actualMACD[i].Time)
	}

	for i := 0; i < len(actualSignal); i++ {
		require.Equalf(t, expectedSignal[i].Time, actualSignal[i].Time, "compare time: %d", i)
		require.Equalf(t, expectedSignal[i].Value.Round(4).String(), actualSignal[i].Value.Round(4).String(), "compare : %d", actualMACD[i].Time)
	}

	for i := 0; i < len(actualHist); i++ {
		require.Equalf(t, expectedHist[i].Time, actualHist[i].Time, "compare time: %d", i)
		require.Equalf(t, expectedHist[i].Value.Round(4).String(), actualHist[i].Value.Round(4).String(), "compare : %d", actualMACD[i].Time)
	}
}
