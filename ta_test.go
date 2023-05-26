package ta

import "github.com/shopspring/decimal"

var (
	testPeriod = 3

	testSrcSet = []Value{
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
	// High >= Open,Close >= Low

	testHighSet = []Value{
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

	testLowSet = []Value{
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

	testCloseSet = []Value{
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
