package ta

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

var (
	testHighSet1 = []Value{
		{
			Time: 1,
			Value: decimal.NewFromFloat(25740.00000000),
		},
		{
			Time: 2,
			Value: decimal.NewFromFloat(25740.00000000),
		},
		{
			Time: 3,
			Value: decimal.NewFromFloat(25725.02000000),
		},
		{
			Time: 4,
			Value: decimal.NewFromFloat(25718.88000000),
		},
		{
			Time: 5,
			Value: decimal.NewFromFloat(25720.16000000),
		},
		{
			Time: 6,
			Value: decimal.NewFromFloat(25730.84000000),
		},
		{
			Time: 7,
			Value: decimal.NewFromFloat(25727.51000000),
		},
		{
			Time: 8,
			Value: decimal.NewFromFloat(25728.99000000),
		},
		{
			Time: 9,
			Value: decimal.NewFromFloat(25724.77000000),
		},
		{
			Time: 10,
			Value: decimal.NewFromFloat(25714.84000000),
		},
		{
			Time: 11,
			Value: decimal.NewFromFloat(25738.74000000),
		},
	}

	testLowSet1 = []Value{
		{
			Time: 1,
			Value: decimal.NewFromFloat(25695.00000000),
		},
		{
			Time: 2,
			Value: decimal.NewFromFloat(25703.87000000),
		},
		{
			Time: 3,
			Value: decimal.NewFromFloat(25705.15000000),
		},
		{
			Time: 4,
			Value: decimal.NewFromFloat(25718.87000000),
		},
		{
			Time: 5,
			Value: decimal.NewFromFloat(25716.30000000),
		},
		{
			Time: 6,
			Value: decimal.NewFromFloat(25727.50000000),
		},
		{
			Time: 7,
			Value: decimal.NewFromFloat(25724.76000000),
		},
		{
			Time: 8,
			Value: decimal.NewFromFloat(25710.28000000),
		},
		{
			Time: 9,
			Value: decimal.NewFromFloat(25709.38000000),
		},
		{
			Time: 10,
			Value: decimal.NewFromFloat(25710.21000000),
		},
	}

	testCloseSet1 = []Value{
		{
			Time: 1,
			Value: decimal.NewFromFloat(25719.27000000),
		},
		{
			Time: 2,
			Value: decimal.NewFromFloat(25705.15000000),
		},
		{
			Time: 3,
			Value: decimal.NewFromFloat(25718.87000000),
		},
		{
			Time: 4,
			Value: decimal.NewFromFloat(25719.79000000),
		},
		{
			Time: 5,
			Value: decimal.NewFromFloat(25727.51000000),
		},
		{
			Time: 6,
			Value: decimal.NewFromFloat(25727.51000000),
		},
		{
			Time: 7,
			Value: decimal.NewFromFloat(25724.77000000),
		},
		{
			Time: 8,
			Value: decimal.NewFromFloat(25711.71000000),
		},
		{
			Time: 9,
			Value: decimal.NewFromFloat(25710.21000000),
		},
		{
			Time: 10,
			Value: decimal.NewFromFloat(25710.21000000),
		},
	}
)

var ( // calculated with timeperiod 4
	expectedCCI = []Value{
		{
			Time: 0,
			Value: decimal.NewFromFloat(0),
		},
		{
			Time: 1,
			Value: decimal.NewFromFloat(0),
		},
		{
			Time: 2,
			Value: decimal.NewFromFloat(-70.342562),
		},
		{
			Time: 3,
			Value: decimal.NewFromFloat(-7.943038),
		},
		{
			Time: 4,
			Value: decimal.NewFromFloat(100.000000),
		},
		{
			Time: 5,
			Value: decimal.NewFromFloat(99.811202),
		},
		{
			Time: 6,
			Value: decimal.NewFromFloat(79.863395),
		},
		{
			Time: 7,
			Value: decimal.NewFromFloat(-1.095198),
		},
		{
			Time: 8,
			Value: decimal.NewFromFloat(-100.000000),
		},
		{
			Time: 9,
			Value: decimal.NewFromFloat(-77.908765),
		},
		{
			Time: 10,
			Value: decimal.NewFromFloat(-32.590315),
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
