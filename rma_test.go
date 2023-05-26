package ta

import "github.com/shopspring/decimal"

var (
	testPeriod = 0.5

	// smoothing factor / n / period = 0.5

	testRmaSet = []Value{
		{
			Time:  1,
			Value: decimal.NewFromFloat(100),
		},
		{
			Time:  2,
			Value: decimal.NewFromFloat(105),
		},
		{
			Time:  3,
			Value: decimal.NewFromFloat(110),
		},
		{
			Time:  4,
			Value: decimal.NewFromFloat(115),
		},
		{
			Time:  5,
			Value: decimal.NewFromFloat(120),
		},
		{
			Time:  6,
			Value: decimal.NewFromFloat(125),
		},
		{
			Time:  7,
			Value: decimal.NewFromFloat(130),
		},
		{
			Time:  8,
			Value: decimal.NewFromFloat(135),
		},
		{
			Time:  9,
			Value: decimal.NewFromFloat(140),
		},
		{
			Time:  10,
			Value: decimal.NewFromFloat(145),
		},
		{
			Time:  11,
			Value: decimal.NewFromFloat(155),
		},
	}

var expectedRMA = []Value{
	{
		Time:  1,
		Value: decimal.NewFromFloat(100.000000),
	},
	{
		Time:  2,
		Value: decimal.NewFromFloat(103.333333),
	},
	{
		Time:  3,
		Value: decimal.NewFromFloat(107.142857),
	},
	{
		Time:  4,
		Value: decimal.NewFromFloat(111.333333),
	},
	{
		Time:  5,
		Value: decimal.NewFromFloat(115.806452),
	},
	{
		Time:  6,
		Value: decimal.NewFromFloat(120.476190),
	},
	{
		Time:  7,
		Value: decimal.NewFromFloat(125.275591),
	},
	{
		Time:  8,
		Value: decimal.NewFromFloat(130.156863),
	},
	{
		Time:  9,
		Value: decimal.NewFromFloat(135.088063),
	},
	{
		Time:  10,
		Value: decimal.NewFromFloat(140.048876),
	},
	{
		Time:  11,
		Value: decimal.NewFromFloat(147.528090),
	},
}