package ta_test

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"

	ta "github.com/irbgeo/gota"
)

var (
	testPeriodBB = 3
	testMultiBB  = 5

	testSrcSetBB = []ta.Value{
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
	expectedMiddleBB = []ta.Value{
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

	expectedUpperBB = []ta.Value{
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

	expectedLowerBB = []ta.Value{
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

func TestBB(t *testing.T) {
	actualMiddleBB, actualUpperBB, actualLowerBB := ta.BB(testSrcSetBB, testPeriodBB, testMultiBB)

	require.Equal(t, len(expectedMiddleBB), len(actualMiddleBB), "compare len middle")
	require.Equal(t, len(expectedUpperBB), len(actualUpperBB), "compare len upper")
	require.Equal(t, len(expectedLowerBB), len(actualLowerBB), "compare len lower")

	for i := 0; i < len(actualMiddleBB); i++ {
		require.Equalf(t, expectedMiddleBB[i].Time, actualMiddleBB[i].Time, "compare time: %d", i)
		require.Equalf(t, expectedMiddleBB[i].Value.Round(4).String(), actualMiddleBB[i].Value.Round(4).String(), "compare : %d", actualMiddleBB[i].Time)
	}

	for i := 0; i < len(actualUpperBB); i++ {
		require.Equalf(t, expectedUpperBB[i].Time, actualUpperBB[i].Time, "compare time: %d", i)
		require.Equalf(t, expectedUpperBB[i].Value.Round(4).String(), actualUpperBB[i].Value.Round(4).String(), "compare : %d", actualMiddleBB[i].Time)
	}

	for i := 0; i < len(actualLowerBB); i++ {
		require.Equalf(t, expectedLowerBB[i].Time, actualLowerBB[i].Time, "compare time: %d", i)
		require.Equalf(t, expectedLowerBB[i].Value.Round(4).String(), actualLowerBB[i].Value.Round(4).String(), "compare : %d", actualMiddleBB[i].Time)
	}
}
