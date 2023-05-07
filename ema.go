package ta

import "github.com/shopspring/decimal"

// EMA - Exponential Moving Average
func EMA(source Series, period int) Series {
	if len(source) < period {
		return nil
	}

	ema := source[0].Value
	multiplier := decimal.NewFromFloat(2.0 / float64(period+1))
	result := make(Series, len(source))

	for i, val := range source[1:] {
		ema = val.Value.Sub(ema).Mul(multiplier).Add(ema)

		result[i] = Value{Time: val.Time, Value: ema}
	}

	return result
}
