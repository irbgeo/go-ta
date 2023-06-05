package ta

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

var (
	testHighSet1 = []Value{
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

	testLowSet1 = []Value{
		{
			Time:  1,
			Value: decimal.NewFromFloat(4),
		},
		{
			Time:  2,
			Value: decimal.NewFromFloat(7.1),
		},
		{
			Time:  3,
			Value: decimal.NewFromFloat(0.2993),
		},
		{
			Time:  4,
			Value: decimal.NewFromFloat(1.5712),
		},
		{
			Time:  5,
			Value: decimal.NewFromFloat(-5.0277),
		},
		{
			Time:  6,
			Value: decimal.NewFromFloat(800),
		},
		{
			Time:  7,
			Value: decimal.NewFromFloat(3.5),
		},
		{
			Time:  8,
			Value: decimal.NewFromFloat(0.71),
		},
		{
			Time:  9,
			Value: decimal.NewFromFloat(-1.61),
		},
		{
			Time:  10,
			Value: decimal.NewFromFloat(3.61),
		},
		{
			Time:  11,
			Value: decimal.NewFromFloat(70.12),
		},
	}

	testCloseSet1 = []Value{
		{
			Time:  1,
			Value: decimal.NewFromFloat(6.3),
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

var (
	expectedCCI = []Value{
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
			Value: decimal.NewFromFloat(5.1664333333333333),
		},
		{
			Time:  4,
			Value: decimal.NewFromFloat(4.0346888888888889),
		},
		{
			Time:  5,
			Value: decimal.NewFromFloat(1.3472259259259259),
		},
		{
			Time:  6,
			Value: decimal.NewFromFloat(300.8981506172839506),
		},
		{
			Time:  7,
			Value: decimal.NewFromFloat(202.9321004115226337),
		},
		{
			Time:  8,
			Value: decimal.NewFromFloat(136.1914002743484225),
		},
		{
			Time:  9,
			Value: decimal.NewFromFloat(89.8909335162322817),
		},
		{
			Time:  10,
			Value: decimal.NewFromFloat(61.8306223441548545),
		},
		{
			Time:  11,
			Value: decimal.NewFromFloat(66.9237482294365697),
		},
	}
)

func TestCCI(t *testing.T) {
	actualCCI := CCI(testHighSet1, testLowSet1, testCloseSet1, testPeriod)

	fmt.Println(actualCCI)

	// require.Equal(t, len(expectedCCI), len(actualCCI), "compare len")
	// for i := 0; i < len(actualCCI); i++ {
	// 	require.Equalf(t, expectedCCI[i].Time, actualCCI[i].Time, "compare time: %d", i+1)
	// 	require.Equalf(t, expectedCCI[i].Value.Round(4).String(), actuaCCIA[i].Value.Round(4).String(), "compare : %d", i+1)
	// }
}
